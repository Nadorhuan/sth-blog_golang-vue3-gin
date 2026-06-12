package request

// InitSysConfigReq 首次初始化提交表单入参
type InitSysConfigReq struct {
	SiteName      string `json:"site_name" binding:"required,min=2,max=100"`
	AdminUsername string `json:"admin_username" binding:"required,min=4,max=30"`
	AdminPassword string `json:"admin_password" binding:"required,min=6,max=30"`
	SiteDesc      string `json:"site_desc" binding:"max=1000"`
	Email         string `json:"email" binding:"email,max=100"`
	SiteIcon      string `json:"site_icon" binding:"max=255,url"`
	ServerPort     string `json:"server_port" binding:"required,numeric"`
}

// ResetInitReq 管理员重置初始化状态入参
type ResetInitReq struct{}
