package consts

const (
	ProjectName              = "goframe-shop"
	ProjectUsage             = "ilotus"
	ProjectBrief             = "start http server"
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	GtokenAdminPrefix        = "Admin:"             // gtoken登录 管理后台 前缀区分
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
	CacheModeRedis           = 2
	BackendServerName        = "goframe-shop"
	MultiLogin               = true
	GTokenExpireIn           = 10 * 24 * 60 * 60
	// 统一管理错误信息
	CodeMissingParameterMsg = "请检查是否缺少参数"
	ErrLoginFaulMsg         = "登录失败，账号或密码错误"
)
