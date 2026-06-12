package models

import (
	"gorm.io/gorm"
)

type SysConfig struct {
	gorm.Model
	SiteName      string `gorm:"type:varchar(100);not null;comment:站点名称"`
	AdminUsername string `gorm:"type:varchar(30);not null;comment:管理员用户名"`
	AdminPassword string `gorm:"type:varchar(100);not null;comment:管理员密码"`
	SiteDesc      string `gorm:"type:text;comment:站点描述"`
	Email         string `gorm:"type:varchar(100);comment:站点邮箱"`
	SiteIcon      string `gorm:"type:varchar(255);comment:站点图标URL"`
	ServerPort    string `gorm:"type:varchar(20);comment:服务器端口"`
	IsInitialized bool   `gorm:"type:boolean;default:false;comment:false=未初始化,true=已初始化"`
}

func (SysConfig) TableName() string {
	return "sys_config"
}
