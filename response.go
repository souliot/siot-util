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
	StatusTypeErrUser     = 400
	StatusTypeErrSystem   = 500
	StatusTypeErrDatabase = 600
	StatusTypeErrMq       = 700
	StatusTypeErrCache    = 800
	StatusTypeErrEtcd     = 900
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
	Err404          = &ControllerError{400, 404, "page not found", "page not found", ""}
	ErrUserExist    = &ControllerError{400, 10001, "用户操作错误", "用户账户已存在", ""}
	ErrUserInput    = &ControllerError{400, 10002, "用户操作错误", "用户输入参数错误", ""}
	ErrUserModify   = &ControllerError{400, 10003, "用户操作错误", "修改用户错误", ""}
	ErrUserDelete   = &ControllerError{400, 10004, "用户操作错误", "删除用户错误", ""}
	ErrNoUser       = &ControllerError{400, 10005, "用户操作错误", "用户账户不存在", ""}
	ErrUserGet      = &ControllerError{400, 10006, "用户操作错误", "获取所有用户错误", ""}
	ErrUserGetLogs  = &ControllerError{400, 10007, "用户操作错误", "获取用户登录日志错误", ""}
	ErrAppid        = &ControllerError{400, 10011, "用户认证错误", "无效的Appid或Secret", ""}
	ErrExpired      = &ControllerError{400, 10012, "登录已过期", "验证token过期", ""}
	ErrPermission   = &ControllerError{400, 10013, "没有权限", "没有操作权限", ""}
	ErrInputData    = &ControllerError{400, 20001, "数据输入错误", "客户端参数错误", ""}
	ErrVersionCheck = &ControllerError{400, 20002, "版本检查", "当前版本过低", ""}
)

var (
	ErrOpenFile     = &ControllerError{500, 10009, "服务器错误", "打开文件出错", ""}
	Err500          = &ControllerError{500, 10001, "服务器错误", "接口访问出错，请确认参数正确！", ""}
	ErrWriteFile    = &ControllerError{500, 10010, "服务器错误", "写文件出错", ""}
	ErrSystem       = &ControllerError{500, 10011, "服务器错误", "操作系统错误", ""}
	ErrTransferData = &ControllerError{500, 20001, "数据转换错误", "Json 字符串转 Map 错误", ""}
)
var (
	ErrEtcdClient = &ControllerError{900, 10001, "Etcd错误", "获取etcd连接错误！", ""}
	ErrEtcdPut    = &ControllerError{900, 10002, "Etcd错误", "etcd put 错误！", ""}
	ErrEtcdGet    = &ControllerError{900, 10003, "Etcd错误", "etcd get 错误！", ""}
	ErrEtcdDelete = &ControllerError{900, 10004, "Etcd错误", "etcd delete 错误！", ""}
	ErrEtcdWatch  = &ControllerError{900, 10005, "Etcd错误", "etcd watch 错误！", ""}
)
