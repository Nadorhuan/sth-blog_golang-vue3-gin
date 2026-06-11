package middleware

import (
	"blog/internal/common"
	"blog/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// JwtAuth Gin JWT登录校验中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取请求头 Authorization
		authHeader := c.GetHeader("Authorization") // 从请求头获取Authorization字段
		if authHeader == "" {
			common.Error(c, common.ERROR_UNAUTHORIZED, common.MsgUnauthorized)
			c.Abort() // 终止后续处理
			return
		}

		// 2. 校验格式：Bearer token
		parts := strings.SplitN(authHeader, " ", 2) // 按空格分割成两部分，最多分割一次
		// 正确格式是： Bearer token，所以分割后应该是两部分，且第一部分是"Bearer"
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			common.Error(c, common.ERROR_UNAUTHORIZED, common.MsgTokenInvalid)
			c.Abort()
			return
		}

		// 3. 解析Token（调用已写好的utils/jwt.go）
		claims, err := utils.ParseToken(parts[1]) // parts[1]是token字符串
		if err != nil {
			common.Error(c, common.ERROR_UNAUTHORIZED, common.MsgTokenInvalid)
			c.Abort()
			return
		}

		// 4. 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		// 5. 放行
		c.Next()
	}
}
