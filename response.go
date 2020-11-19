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
	ErrOpenFile     = &ControllerError{500, 10001, "服务器错误", "打开文件出错", ""}
	Err500          = &ControllerError{500, 10002, "服务器错误", "接口访问出错", ""}
	ErrWriteFile    = &ControllerError{500, 10003, "服务器错误", "写文件出错", ""}
	ErrSystem       = &ControllerError{500, 10004, "服务器错误", "操作系统错误", ""}
	ErrTransferData = &ControllerError{500, 10005, "数据转换错误", "Json 字符串转 Map 错误", ""}
)

var (
	ErrDb                  = &ControllerError{600, 10001, "数据库错误", "数据库操作错误", ""}
	ErrDupRecord           = &ControllerError{600, 10002, "数据库错误", "数据库记录重复", ""}
	ErrNoRecord            = &ControllerError{600, 10003, "数据库错误", "数据库记录不存在", ""}
	ErrUserPass            = &ControllerError{600, 10004, "数据库错误", "用户名或密码不正确", ""}
	ErrDbInsert            = &ControllerError{600, 10005, "数据库错误", "数据添加错误", ""}
	ErrDbRead              = &ControllerError{600, 10006, "数据库错误", "数据读取错误", ""}
	ErrDbUpdate            = &ControllerError{600, 10007, "数据库错误", "数据更新错误", ""}
	ErrDbDelete            = &ControllerError{600, 10008, "数据库错误", "数据删除错误", ""}
	ErrChangeAccountStatus = &ControllerError{600, 10009, "数据库错误", "更新账户状态错误", ""}
)

var (
	ErrFileSystem       = &ControllerError{700, 10000, "文件系统错误", "文件系统连接错误！", ""}
	ErrFileExist        = &ControllerError{700, 10001, "文件系统错误", "文件已存在！", ""}
	ErrFileNonExist     = &ControllerError{700, 10002, "文件系统错误", "文件不存在！", ""}
	ErrFileUpload       = &ControllerError{700, 10003, "文件系统错误", "文件上传错误！", ""}
	ErrFileDownload     = &ControllerError{700, 10004, "文件系统错误", "文件下载错误！", ""}
	ErrFileDelete       = &ControllerError{700, 10005, "文件系统错误", "文件删除错误！", ""}
	ErrFileList         = &ControllerError{700, 10006, "文件系统错误", "获取文件列表错误！", ""}
	ErrFileCopy         = &ControllerError{700, 10007, "文件系统错误", "文件转存错误！", ""}
	ErrFileSetLifeCycle = &ControllerError{700, 10007, "文件系统错误", "设置存储桶生命周期错误！", ""}
	ErrFileGetLifeCycle = &ControllerError{700, 10007, "文件系统错误", "获取存储桶生命周期错误！", ""}
)

var (
	ErrEtcdClient = &ControllerError{900, 10001, "Etcd错误", "获取etcd连接错误！", ""}
	ErrEtcdPut    = &ControllerError{900, 10002, "Etcd错误", "etcd put 错误！", ""}
	ErrEtcdGet    = &ControllerError{900, 10003, "Etcd错误", "etcd get 错误！", ""}
	ErrEtcdDelete = &ControllerError{900, 10004, "Etcd错误", "etcd delete 错误！", ""}
	ErrEtcdWatch  = &ControllerError{900, 10005, "Etcd错误", "etcd watch 错误！", ""}
)

var (
	ErrProxyInput = &ControllerError{1000, 10001, "网关代理错误", "无效的代理地址！", ""}
)
