package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 加密密码（注册时用）
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //bcrypt.DefaultCost是一个常量，表示默认的加密强度，值为10。这个值越大，加密过程就越慢，但安全性也更高。通常10是一个合理的选择，既能提供足够的安全性，又不会导致过长的加密时间。
	return string(bytes), err
}

// CheckPasswordHash 验证密码（登录时用）
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) //CompareHashAndPassword函数接受两个参数：第一个是存储在数据库中的哈希值，第二个是用户输入的密码。函数会将用户输入的密码进行相同的哈希处理，然后与存储的哈希值进行比较。如果两者匹配，函数返回nil；如果不匹配，返回一个错误。
	return err == nil
}
