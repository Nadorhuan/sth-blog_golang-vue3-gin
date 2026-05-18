package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"sfh-blog/internal/common"
	"sfh-blog/internal/database"
	"sfh-blog/internal/models"
	"sfh-blog/internal/request"
	"sfh-blog/internal/utils"
)

// Register 用户注册接口
func Register(c *gin.Context) {
	// 1. 定义变量接收前端json参数
	var req request.RegisterRequest
	// 2. 绑定并自动校验参数
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, common.ERROR_BAD_REQUEST, common.MsgParamError)
		return
	}
	// 3. 校验邮箱是否已注册
	var count int64
	database.DB.
		Model(&models.User{}).
		Where("email = ?", req.Email).
		Count(&count)
	if count > 0 {
		common.Error(c, common.ERROR_BAD_REQUEST, common.MsgEmailExists)
		return
	}
	//4. 校验用户是否已注册
	var count2 int64
	database.DB.
		Model(&models.User{}).
		Where("username = ?", req.Username).
		Count(&count2)
	if count2 > 0 {
		common.Error(c, common.ERROR_BAD_REQUEST, common.MsgUsernameExists)
		return
	}
	// 5. 对铭文密码进行bcrypt加密
	hashPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		common.Error(c, common.ERROR_INTERNAL_SERVER, common.MsgServerError)
		return
	}
	// 6. 组装用户模型
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPwd,
		// Role和Status有默认值，不需要手动设置
	}
	// 7. 写入数据库
	if err := database.DB.Create(&user).Error; err != nil {
		common.Error(c, common.ERROR_INTERNAL_SERVER, common.MsgServerError)
		return
	}
	// 8. 返回成功响应
	common.SuccessMsg(c, common.MsgSuccess, nil)
}

// Login用户登录接口
func Login(c *gin.Context) { // 首字母必须大写，才能被main.go中的handler.Login访问到。
	// 1. 接收并自动校验登录参数
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, common.ERROR_BAD_REQUEST, common.MsgParamError)
		return
	}
	// 2. 根据用户名 or 邮箱查询用户
	var user models.User
	err := database.DB.
		Where("username = ? OR email = ?", req.Username, req.Username). // 这里的req.Username既可以是用户名，也可以是邮箱，因为前端登录时可能输入用户名或邮箱
		First(&user).Error                                              // First方法会返回第一个匹配的记录，如果没有找到会返回错误
	if err != nil {
		common.Error(c, common.ERROR_UNAUTHORIZED, common.MsgUserNotFound)
		return
	}

	// 3. 比对密码
	ok := utils.CheckPasswordHash(req.Password, user.Password)
	if !ok {
		common.Error(c, common.ERROR_UNAUTHORIZED, common.MsgPasswordError)
		return
	}
	// 4. 登录成功
	// 生成JWT token
	token, _ := utils.GenerateToken(user.ID, user.Username)
	// 返回用户信息和token
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"userId":   user.ID,
			"username": user.Username,
			"email":    user.Email,
			"token":    token, // 返回jwt
		},
	})
}
