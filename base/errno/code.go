package errno

/*
服务级别错误码：1 位数进行表示，比如 1 为系统级错误；2 为普通错误，通常是由用户非法操作引起。
模块级错误码：2 位数进行表示，比如 01 为用户模块；02 为订单模块。
具体错误码：2 位数进行表示，比如 01 为手机号不合法；02 为验证码输入错误。
*/

var (
	// OK
	OK = NewError(0, "OK")

	// 服务级错误码
	ErrServer        = NewError(10001, "服务异常，请联系管理员")
	ErrParam         = NewError(10002, "参数有误")
	ErrSignParam     = NewError(10003, "签名参数有误")
	ErrUnknownServer = NewError(10004, "服务未知错误")

	// 模块级错误码 - 用户模块
	ErrLink         = NewError(20101, "链接格式错误")
	ErrLinkNotExist = NewError(20102, "连接不存在")
)
