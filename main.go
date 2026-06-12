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
	// ========== 第一层：按鉴权等级拆分两大顶级组 ==========
	// 1. 公开无鉴权大组
	publicGroup := r.Group("/api/public")
	{
		// 在公开组内，再拆分user业务子组
		publicUser := publicGroup.Group("/user")
		{
			publicUser.POST("/register", handler.Register)
			publicUser.POST("/login", handler.Login)
		}

		// 在公开组内，拆分初始化公开子组
		publicInit := publicGroup.Group("/init")
		{
			publicInit.GET("/status", handler.CheckInitStatus)
			publicInit.POST("/submit", handler.SubmitInit)
		}
	}

	// 2. 管理员鉴权大组：全局挂载JWT中间件，组内所有接口强制鉴权
	adminGroup := r.Group("/api/admin")
	adminGroup.Use(middleware.JwtAuth())
	{
		// 鉴权组内，拆分初始化管理子组
		adminInit := adminGroup.Group("/init")
		{
			adminInit.POST("/reset", handler.ResetInit)
		}

		// 未来拓展：登录后用户操作、后台管理接口都在这里新建子组
		// adminUser := adminGroup.Group("/user")
		// adminUser.POST("/modify-pwd", handler.ModifyPwd)
	}

	// 从配置启动服务
	log.Printf("server running on :%s\n", config.AppConfig.Port)
	err := r.Run(":" + config.AppConfig.Port)
	if err != nil {
		log.Fatal("服务启动失败：", err)
	}
}
