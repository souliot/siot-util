package sutil

type ControllerError struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Type     string `json:"type"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
}

type ControllerSuccess struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// define Col operations
const (
	TypeErrSystem   = "服务器错误"
	TypeErrDatabase = "数据库错误"
	TypeErrMq       = "消息服务错误"
	TypeErrCache    = "缓存服务错误"
	TypeErrUser     = "用户操作错误"
)

const (
	StatusSuccess         = 200
	StatusTypeErrSystem   = 500
	StatusTypeErrDatabase = 600
	StatusTypeErrMq       = 700
	StatusTypeErrCache    = 800
	StatusTypeErrUser     = 400
)

var (
	Actionsuccess = &ControllerError{200, 90000, "操作成功", "操作成功", ""}
)

var (
	ErrDb                  = &ControllerError{600, 10001, "数据库错误", "数据库操作错误", ""}
	ErrDupRecord           = &ControllerError{600, 10002, "数据库错误", "数据库记录重复", ""}
	ErrNoRecord            = &ControllerError{600, 10003, "数据库错误", "数据库记录不存在", ""}
	ErrUserPass            = &ControllerError{600, 10004, "数据库错误", "用户名或密码不正确", ""}
	ErrDbInsert            = &ControllerError{600, 10005, "数据库错误", "数据添加错误", ""}
	ErrDbUpdate            = &ControllerError{600, 10006, "数据库错误", "数据更新错误", ""}
	ErrDbDelete            = &ControllerError{600, 10007, "数据库错误", "数据删除错误", ""}
	ErrChangeAccountStatus = &ControllerError{600, 10008, "数据库错误", "更新账户状态错误", ""}
)

var (
	ErrInvalidUser = &ControllerError{400, 10008, "用户信息不正确", "Session信息不正确", ""}
	ErrOpenFile    = &ControllerError{500, 10009, "服务器错误", "打开文件出错", ""}
	ErrWriteFile   = &ControllerError{500, 10010, "服务器错误", "写文件出错", ""}
	ErrSystem      = &ControllerError{500, 10011, "服务器错误", "操作系统错误", ""}
	ErrExpired     = &ControllerError{400, 10012, "登录已过期", "验证token过期", ""}
	ErrPermission  = &ControllerError{400, 10013, "没有权限", "没有操作权限", ""}
)

var (
	Err404          = &ControllerError{StatusTypeErrUser, 404, "page not found", "page not found", ""}
	ErrInputData    = &ControllerError{400, 10001, "数据输入错误", "客户端参数错误", ""}
	ErrTransferData = &ControllerError{500, 20001, "数据转换错误", "Json 字符串转 Map 错误", ""}
)
