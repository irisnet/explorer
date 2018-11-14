package types

var (
	ErrorCodeNotFound     = NewError("EX-100001", "Data Not Found")
	ErrorCodeSysFailed    = NewError("EX-100002", "System Exception")
	ErrorCodeInValidParam = NewError("EX-100003", "invalid param")
	ErrorCodeExtSysFailed = NewError("EX-100004", "External system exception")
)

type Error struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code string, msg string) Error {
	return Error{
		Code: code,
		Msg:  msg,
	}
}

func (e Error) Success() bool {
	return len(e.Code) == 0
}
