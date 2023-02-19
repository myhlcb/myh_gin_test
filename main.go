package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	_ "github.com/thinkerou/favicon"
)

// 中间件
func myHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware")
		c.Set("cookie", "1111111")
		c.Next() // 方形
		// c.Abort() // 阻止
	}
}
func main() {
	//创建服务
	ginServer := gin.Default()
	// 添加图标
	ginServer.Use(favicon.New("./icon.jpg"))
	// 加载index文件
	ginServer.LoadHTMLGlob("templates/*")
	// 加载静态文件
	ginServer.Static("static", "./static")
	// url访问
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})
	// index.html
	ginServer.GET("/index", func(context *gin.Context) {
		// http.StatusOk = 200
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是后台res数据",
		})
	})
	// 404.html
	ginServer.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	// user/info?xx=xx&xx=xx
	ginServer.GET("/user/info", func(context *gin.Context) {
		// http.StatusOk = 200
		uid := context.Query("userId")
		userName := context.Query("userName")
		context.JSON(http.StatusOK, gin.H{
			"userId":   uid,
			"userName": userName,
		})
	})
	// 重定向
	ginServer.GET("/test", func(context *gin.Context) {
		context.Redirect(301, "http://www.baidu.com")
	})
	// user/info/1/xx
	ginServer.GET("/user/info/:userId/:userName", func(context *gin.Context) {
		// http.StatusOk = 200
		uid := context.Param("userId")
		userName := context.Param("userName")
		context.JSON(http.StatusOK, gin.H{
			"userId":   uid,
			"userName": userName,
		})
	})
	// raw json
	ginServer.POST("/user", func(c *gin.Context) {
		// request body
		data, _ := c.GetRawData()
		var m map[string]interface{}
		// 包装json
		json.Unmarshal(data, &m)
		c.JSON(200, m)
	})
	// x-www-form-urlencoded
	ginServer.POST("/user2", func(c *gin.Context) {
		// request body
		userName := c.PostForm("userName")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"username": userName,
			"password": password,
		})
	})

	//路由组/user/add
	userGroup := ginServer.Group("/user")
	{
		userGroup.POST("/add", myHandler(), func(c *gin.Context) {
			// request body
			cookie := c.MustGet("cookie").(string)
			userName := c.PostForm("userName")
			password := c.PostForm("password")
			c.JSON(200, gin.H{
				"username": userName,
				"password": password,
				"cookie":   cookie,
			})
		})
		// userGroup.GET("/add")
		// userGroup.GET("/add")
	}

	// 服务器端口
	ginServer.Run(":8082")
}
