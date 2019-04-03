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

var (
	Err404       = &ControllerError{404, 404, "page not found", "page not found", ""}
	ErrInputData = &ControllerError{400, 10001, "数据输入错误", "客户端参数错误", ""}

	ErrTransferData = &ControllerError{500, 20001, "数据转换错误", "Json 字符串转 Map 错误", ""}

	ErrDatabase     = &ControllerError{400, 10002, "数据库错误", "数据库操作错误", ""}
	ErrDupRecord    = &ControllerError{400, 10003, "数据库错误", "数据库记录重复", ""}
	ErrNoRecord     = &ControllerError{400, 10004, "数据库错误", "数据库记录不存在", ""}
	ErrPass         = &ControllerError{400, 10005, "数据库错误", "密码不正确", ""}
	ErrNoUserPass   = &ControllerError{400, 10006, "数据库错误", "数据库记录不存在或密码不正确", ""}
	ErrNoUserChange = &ControllerError{400, 10007, "数据库错误", "数据库记录不存在或数据未改变", ""}

	ErrInvalidUser = &ControllerError{400, 10008, "用户信息不正确", "Session信息不正确", ""}

	ErrOpenFile   = &ControllerError{500, 10009, "服务器错误", "打开文件出错", ""}
	ErrWriteFile  = &ControllerError{500, 10010, "服务器错误", "写文件出错", ""}
	ErrSystem     = &ControllerError{500, 10011, "服务器错误", "操作系统错误", ""}
	ErrExpired    = &ControllerError{400, 10012, "登录已过期", "验证token过期", ""}
	ErrPermission = &ControllerError{400, 10013, "没有权限", "没有操作权限", ""}
	Actionsuccess = &ControllerError{200, 90000, "操作成功", "操作成功", ""}

	ErrDeviceQrAdd = &ControllerError{400, 20000, "添加失败", "设备未在线或已被别的用户管理！", ""}
)
