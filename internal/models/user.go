package models

import (
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model        //gorm.Model是一个包含ID、CreatedAt、UpdatedAt和DeletedAt字段的结构体，嵌入后会自动拥有这些字段
	Username   string `gorm:"size:50;not null;unique;comment:用户名"`
	Email      string `gorm:"size:100;unique;not null;comment:邮箱地址"`
	Password   string `gorm:"size:100;not null;comment:加密后的密码"` //密码是 bcrypt 哈希串，长度约 60 位
	Role       string `gorm:"size:20;default:user;not null;comment:角色 user/admin"`
	Status     int    `gorm:"default:1;not null;comment:状态 1:正常 0:禁用"`
}
