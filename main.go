package main

import (
	"blog/internal/config"
	"blog/internal/database"
	"blog/internal/handler"
	"blog/internal/middleware"
	"log"

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
	// 公开路由，无需鉴权
	public := r.Group("/api")
	{
		// 用户注册接口
		public.POST("/register", handler.Register)
		// 用户登录接口
		public.POST("/login", handler.Login)
	}
	// 私有路由，需要鉴权
	authApi := r.Group("/api")
	authApi.Use(middleware.JwtAuth())
	{
		// 这里可以添加需要鉴权的接口，例如用户信息、文章管理等

	}

	// 从配置启动服务
	log.Printf("server running on :%s\n", config.AppConfig.Port)
	err := r.Run(":" + config.AppConfig.Port)
	if err != nil {
		log.Fatal("服务启动失败：", err)
	}
}
