package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建服务
	ginServer := gin.Default()
	//访问地址
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})
	// 服务器端口
	ginServer.Run(":8082")
}
