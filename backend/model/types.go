package model

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type CountVo struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Count float64
}

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
	TraceId int64
	Start   time.Time
}
