package utils

import (
	"blog/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义载荷（存用户信息）
type CustomClaims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID uint, username string) (string, error) {
	// 过期时间：7天
	expireTime := time.Now().Add(time.Hour * time.Duration(config.AppConfig.JWTexpireHours))
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析并验证token
	// 指令1: tokenString是要解析token文本
	// 指令2: &CustomClaims{}解析后把用户信息存到这个自定义结构体中
	// 指令3: func(token *jwt.Token) (interface{}, error)签名验证的规则函数
	//jwt.ParseWithClaims函数接受三个参数：第一个是要解析的token字符串，第二个是一个空的CustomClaims结构体实例，用于存储解析后的用户信息，第三个是一个回调函数，用于验证token的签名是否合法。在回调函数中，我们首先检查token的签名方法是否为HS256，如果不是，则返回一个错误；如果是，则返回JWT密钥供库进行签名验证。
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法是否为HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.AppConfig.JWTSecret), nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
