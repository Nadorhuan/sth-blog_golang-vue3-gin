package handler

import (
	"blog/internal/common"
	"blog/internal/database"
	"blog/internal/models"
	"blog/internal/request"
	"blog/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CheckInitStatus 查询系统初始化状态 【公开接口，无需登录鉴权】
func CheckInitStatus(c *gin.Context) {
	// 1. 声明sys_config模型变量，用于接收数据库查询结果
	var cfg models.SysConfig
	// 2. GORM First查询第一条配置记录
	err := database.DB.First(&cfg).Error
	// 3. 判断查询是否出错
	if err != nil {
		// 场景1：查询不到记录 = 数据库无配置 = 系统未初始化
		if err == gorm.ErrRecordNotFound {
			// 组装返回给前端的数据map
			data := gin.H{"is_initialized": false}
			// 调用公共成功返回方法，统一格式输出JSON
			common.SuccessMsg(c, common.MsgSuccess, data)
			// 终止当前函数，不再向下执行
			return
		}
		// 场景2：数据库连接/查询异常，属于服务端500错误
		common.Error(c, common.ERROR_INTERNAL_SERVER, "查询配置失败:"+err.Error())
		return
	}
	// 4. 正常查到配置数据，返回初始化布尔状态
	data := gin.H{"is_initialized": cfg.IsInitialized}
	common.SuccessMsg(c, common.MsgSuccess, data)
}

// SubmitInit 提交初始化表单配置 【公开接口，首次部署可用】
func SubmitInit(c *gin.Context) {
	// 1. 声明接收前端JSON的请求结构体变量
	var req request.InitSysConfigReq
	// 2. ShouldBindJSON自动绑定+校验JSON参数，失败进入错误分支
	if err := c.ShouldBindJSON(&req); err != nil {
		// 参数格式/校验失败属于客户端400错误
		common.Error(c, common.ERROR_BAD_REQUEST, common.MsgParamError)
		return
	}

	// 3. 先查询是否已经存在初始化配置
	var existCfg models.SysConfig
	err := database.DB.First(&existCfg).Error
	// 无报错 + 标记已初始化 = 禁止重复提交
	if err == nil && existCfg.IsInitialized {
		// 重复操作属于客户端非法请求400
		common.Error(c, common.ERROR_BAD_REQUEST, "系统已初始化，禁止重复提交")
		return
	}

	// 4. 使用项目工具类加密管理员密码
	encryptPwd, err := utils.HashPassword(req.AdminPassword)
	if err != nil {
		// 加密失败是服务内部异常500
		common.Error(c, common.ERROR_INTERNAL_SERVER, "密码加密失败")
		return
	}

	// 5. 组装完整数据库模型数据
	newCfg := models.SysConfig{
		SiteName:      req.SiteName,      // 站点名称
		AdminUsername: req.AdminUsername, // 管理员账号
		AdminPassword: encryptPwd,        // 加密后的密码
		SiteDesc:      req.SiteDesc,      // 网站简介
		Email:         req.Email,         // 联系邮箱
		SiteIcon:      req.SiteIcon,      // 站点图标地址
		ServerPort:    req.ServerPort,    // 后端服务端口
		IsInitialized: true,              // 标记：已完成初始化
	}

	// 6. Save方法：无记录则新增，有记录则覆盖更新
	if err := database.DB.Save(&newCfg).Error; err != nil {
		// 数据库写入失败属于服务端500
		common.Error(c, common.ERROR_INTERNAL_SERVER, "初始化保存失败:"+err.Error())
		return
	}

	// 7. 全部流程无误，返回成功提示，无额外data数据填nil
	common.SuccessMsg(c, "系统初始化完成", nil)
}

// ResetInit 重置初始化状态 【管理员专属接口，必须携带JWT鉴权】
func ResetInit(c *gin.Context) {
	// 1. 空请求结构体占位，统一代码规范
	var req request.ResetInitReq
	// 2. 校验请求体格式（哪怕无字段，也统一绑定逻辑）
	if err := c.ShouldBindJSON(&req); err != nil {
		// 请求体格式错误属于客户端400
		common.Error(c, common.ERROR_BAD_REQUEST, common.MsgParamError)
		return
	}

	// 3. 批量更新sys_config表is_initialized字段为false（重置未初始化）
	err := database.DB.Model(&models.SysConfig{}).Update("is_initialized", false).Error
	if err != nil {
		// 数据库更新失败属于服务端500
		common.Error(c, common.ERROR_INTERNAL_SERVER, "重置初始化状态失败:"+err.Error())
		return
	}

	// 4. 重置操作成功返回提示
	common.SuccessMsg(c, "已重置系统初始化状态", nil)
}
