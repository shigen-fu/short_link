package errno

type WrapError struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}

func (err WrapError) Error() string {
	return err.msg
}

func (err WrapError) New(errno Error) WrapError {
	return WrapError{
		code: errno.GetCode(),
		msg:  errno.GetMsg(),
	}
}
