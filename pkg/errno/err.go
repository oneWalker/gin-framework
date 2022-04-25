package errno

// Errno - 义错误码
type Errno struct {
	Code    int
	Message string
}

// NewError .
func NewError(code int, msg string) *Errno {
	return &Errno{
		Code:    code,
		Message: msg,
	}
}

func (e *Errno) Error() string {
	return e.Message
}
