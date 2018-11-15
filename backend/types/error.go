package types

var (
	ErrorCodeSuccess      = NewError("0", "success")
	ErrorCodeNotFound     = NewError("EX-100001", "data Not found")
	ErrorCodeSysFailed    = NewError("EX-100002", "system exception")
	ErrorCodeInValidParam = NewError("EX-100003", "invalid input param")
	ErrorCodeExtSysFailed = NewError("EX-100004", "external system exception")
	ErrorCodeUnKnown      = NewError("EX-100005", "unknown error ")
	ErrorCodeUnSupportTx  = NewError("EX-100006", "unsupported tx type")
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
	return e.Code == "0"
}
