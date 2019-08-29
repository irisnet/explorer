package vo

import (
	"net/http"
	"time"
)

type PageVo struct {
	Count int
	Data  interface{}
}

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code, msg string, data interface{}) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

type IrisReq struct {
	*http.Request
	TraceId string
	Start   time.Time
}

type Coin struct {
	Denom  string  `json:"denom"`
	Amount float64 `json:"amount"`
}
