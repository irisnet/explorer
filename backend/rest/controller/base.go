package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/logger"
	"net/http"
	"strconv"
	"time"
)

type IrisReq struct {
	*http.Request
	TraceId int64
	Start   time.Time
}

// user business action
type Action func(request IrisReq) interface{}

func GetString(request IrisReq, key string) (result string) {
	request.ParseForm()
	if len(request.Form[key]) > 0 {
		result = request.Form[key][0]
	}
	return
}

func GetInt(request IrisReq, key string) (result int) {
	value := GetString(request, key)
	if len(value) == 0 {
		return
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		logger.Error("param is not int type", logger.String("param", key))
	}
	return
}

func Var(request IrisReq, key string) (result string) {
	args := mux.Vars(request.Request)
	result = args[key]
	return
}

func GetPage(r IrisReq) (int, int) {
	page := Var(r, "page")
	size := Var(r, "size")
	iPage := 0
	iSize := 20
	if p, ok := utils.ParseInt(page); ok {
		iPage = int(p)
	}
	if s, ok := utils.ParseInt(size); ok {
		iSize = int(s)
	}
	return iPage, iSize
}

// doBefore display user's request information,optional
func doBefore(request IrisReq) {
	traceId := logger.Int64("traceId", request.TraceId)
	apiPath := logger.String("path", request.RequestURI)
	formValue := logger.Any("form", request.Form)
	urlValue := logger.Any("url", mux.Vars(request.Request))
	agent := logger.String("agent", request.UserAgent())
	defer func() {
		if r := recover(); r != nil {
			logger.Error("doBefore failed", traceId)
		}
	}()

	logger.Info("doBefore", traceId, apiPath, agent, formValue, urlValue)
}

// execute user's business code
func doAction(request IrisReq, action Action) interface{} {
	//do business action
	logger.Info("doAction", logger.Int64("traceId", request.TraceId))
	result := action(request)
	return result
}

// doAfter display user's request information,optional
func doAfter(request IrisReq, result interface{}) {
	traceId := logger.Int64("traceId", request.TraceId)
	res := logger.Any("result", result)
	coastSecond := time.Now().Unix() - request.Start.Unix()
	coast := logger.Int64("coast", coastSecond)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("doAfter failed", traceId)
		}
	}()
	logger.Info("doAfter", traceId, res, coast)
	if coastSecond >= 3 {
		logger.Warn("doAfter api coast most time", traceId)
	}
}

// deal with exception for business action
func doException(request IrisReq, writer http.ResponseWriter) {
	if r := recover(); r != nil {
		trace := logger.Int64("traceId", request.TraceId)
		errMsg := logger.Any("errMsg", r)
		switch r.(type) {
		case types.BizCode:
			doResponse(writer, r)
			break
		case error:
			err := r.(error)
			e := types.BizCode{
				Code: types.CodeUnKnown.Code,
				Msg:  err.Error(),
			}
			doResponse(writer, e)
			break
		}
		logger.Error("doException", trace, errMsg)
	}
}

// response action's result to user
func doResponse(writer http.ResponseWriter, data interface{}) {
	var bz []byte

	switch data.(type) {
	case []byte:
		bz = data.([]byte)
	case int64:
		var resp = model.Response{
			Code: types.CodeSuccess.Code,
			Msg:  types.CodeSuccess.Msg,
			Data: data.(int64),
		}
		bz, _ = json.Marshal(resp)
	default:
		bz, _ = json.Marshal(data)
	}

	writer.Write(bz)
}

// doApi
// url : api path
// method : api method type
// action : business code
func doApi(r *mux.Router, url, method string, action Action) {
	//wrap business code
	wrapperAction := func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		req := IrisReq{
			Request: request,
			TraceId: start.UnixNano(),
			Start:   start,
		}
		defer doException(req, writer)
		doBefore(req)
		result := doAction(req, action)
		doAfter(req, result)
		doResponse(writer, result)
	}
	r.HandleFunc(url, wrapperAction).Methods(method)
}
