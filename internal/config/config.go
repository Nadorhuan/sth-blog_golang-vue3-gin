package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// 工具函数：读取环境变量（已修正括号问题）
func getEnv(key string, defaultValue string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	}
	return val
}

// config 全局配置结构体
type Config struct {
	Port           string
	GIN_MODE       string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMode      string
	JWTSecret      string //JWT密钥
	JWTexpireHours int    //JWT过期小时数
	AllowedOrigins string //允许的跨域来源，逗号分隔
}

// 全局变量，整个项目用这个实例
var AppConfig Config

// 初始化配置
func InitConfig() {
	//加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Print("警告：未找到.env文件，将使用系统环境变量默认值")
	}
	//读取数字类型配置：JWT过期小时数
	//第一步：用 getEnv 读 .env 里的 JWT_EXPIRE_HOURS，找不到就用默认值 "24"
	expireHoursStr := getEnv("JWT_EXPIRE_HOURS", "24") // 默认24小时
	//第二步：把字符串转成数字，同时捕获错误
	expireHours, err := strconv.Atoi(expireHoursStr)
	//第三步：处理转数字失败的情况（比如用户在 .env 里写了 "abc"）
	if err != nil {
		log.Printf("警告：JWT_EXPIRE_HOURS配置无效，使用默认值24小时: %v", err) // %v是万能占位符，这里的作用是显示任何类型的err输出信息
		expireHours = 24
	}
	//赋值全局配置变量
	AppConfig = Config{
		Port:           getEnv("PORT", "8080"),
		GIN_MODE:       getEnv("GIN_MODE", "debug"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", "sfh_blog"),
		DBSSLMode:      getEnv("DB_SSLMODE", "disable"),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTexpireHours: expireHours,
		AllowedOrigins: getEnv("ALLOW_ORIGINS", "http://localhost:5173"), //这才是0
	}

}
