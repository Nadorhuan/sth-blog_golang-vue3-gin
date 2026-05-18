package database

import (
	"log"
	"sfh-blog/internal/config"
	"sfh-blog/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 声明全局数据库连接变量，大写 = 公开 全局访问，小写 = 私有 只有包内访问
var DB *gorm.DB

func InitDB() {
	// 拼接postgres连接字符串
	host := config.AppConfig.DBHost
	port := config.AppConfig.DBPort
	user := config.AppConfig.DBUser
	password := config.AppConfig.DBPassword
	dbname := config.AppConfig.DBName
	sslmode := config.AppConfig.DBSSLMode
	// 拼接DSN
	dsn := "host=" + host +
		" port=" + port +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" sslmode=" + sslmode
	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // 多返回值；这里的postgres.Open是gorm的postgres驱动提供的函数，用来创建一个gorm.Dialector接口的实例，参数是我们拼接好的DSN字符串。gorm.Open函数接受这个Dialector实例和一些可选的配置参数，返回一个gorm.DB实例和一个错误对象。
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	DB = db
	log.Println("数据库连接成功")

	// 自动迁移：根据模型结构自动创建或更新数据库表
	err = DB.AutoMigrate(&models.User{}) // 这里的&models.User{}是一个指向User结构体类型的指针，gorm会根据这个结构体的定义来创建或更新数据库中的users表。
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
	log.Println("✅ 数据库初始化完成，用户表已自动创建")
}
