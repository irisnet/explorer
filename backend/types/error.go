package types

import (
	"fmt"
)

var (
	CodeSuccess        = NewCode("0", "success")
	CodeSysmaintenance = NewCode("EX-100000", "the system is under maintenance")
	CodeNotFound       = NewCode("EX-100001", "data Not found")
	CodeSysFailed      = NewCode("EX-100002", "system exception")
	CodeInValidParam   = NewCode("EX-100003", "invalid input param")
	CodeExtSysFailed   = NewCode("EX-100004", "external system exception")
	CodeUnKnown        = NewCode("EX-100005", "unknown error ")
	CodeUnSupportTx    = NewCode("EX-100006", "unsupported tx type")
	CodeJsonMarshal    = NewCode("EX-100007", "json marshal error")
	CodeRateLimit      = NewCode("EX-100008", "exceed the upper limit")
)

type BizCode struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func NewCode(code string, msg string) BizCode {
	return BizCode{
		Code: code,
		Msg:  msg,
	}
}

func (e BizCode) Success() bool {
	return e.Code == "0"
}

func (e BizCode) Error() string {
	return fmt.Sprintf("Code:%s,Msg:%s", e.Code, e.Msg)
}

func ErrForEmpty(args ...string) BizCode {
	if len(args) == 0 {
		return CodeNotFound
	}

	msg := fmt.Sprintf("data not find %s", args)
	return BizCode{
		Code: CodeNotFound.Code,
		Msg:  msg,
	}
}
