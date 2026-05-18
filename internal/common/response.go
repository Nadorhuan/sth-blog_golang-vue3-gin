package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 全局统一返回结构体
type Result struct {
	Code int         `json:"code"` //状态码，200成功，其它失败
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

// Success 成功返回（带数据）
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  MsgSuccess,
		Data: data,
	})
}

// SuccessMsg 成功返回（自定义提示）
func SuccessMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  msg,
		Data: data,
	})
}

// Error 失败返回(默认提示)
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
