package main

import (
	"log"
	"sfh-blog/internal/config"
	"sfh-blog/internal/database"
	"sfh-blog/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载config
	config.InitConfig()
	// 连接数据库
	database.InitDB()
	// 设置Gin运行模式
	gin.SetMode(gin.DebugMode)
	// 创建路由引擎
	r := gin.Default()
	// 测试接口
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "wellcome to SFHaven",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "ok",
		})
	})
	// 用户注册接口
	r.POST("/api/auth/register", handler.Register)
	// 用户登录接口
	r.POST("/api/auth/login", handler.Login)

	// 从配置启动服务
	log.Printf("server running on :%s\n", config.AppConfig.Port)
	err := r.Run(":" + config.AppConfig.Port)
	if err != nil {
		log.Fatal("服务启动失败：", err)
	}
}
