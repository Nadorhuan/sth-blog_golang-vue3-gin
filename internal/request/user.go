package request

// RegisterRequest 用户注册参数
// binding 标签是 Gin 自带的参数校验，自动校验格式、长度、非空
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"` //用户名必须，长度3-50
	Email    string `json:"email" binding:"required,email,max=100"`   //邮箱必须，格式email，长度不超过100
	Password string `json:"password" binding:"required,min=8,max=20"` //密码必须，长度8-20
}

// LoginRequest 用户登录参数
type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=100"` //用户名必须，长度3-50
	Password string `json:"password" binding:"required,min=8,max=20"`  //密码必须，长度8-20
}
