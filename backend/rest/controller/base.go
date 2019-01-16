package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/rest/filter"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"net/http"
	"strconv"
)

// user business action
type Action func(request model.IrisReq) interface{}

func GetString(request model.IrisReq, key string) (result string) {
	request.ParseForm()
	if len(request.Form[key]) > 0 {
		result = request.Form[key][0]
	}
	return
}

func GetInt(request model.IrisReq, key string) (result int) {
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

func Var(request model.IrisReq, key string) (result string) {
	args := mux.Vars(request.Request)
	result = args[key]
	return
}

func GetPage(r model.IrisReq) (int, int) {
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

// execute user's business code
func doAction(request model.IrisReq, action Action) interface{} {
	//do business action
	logger.Debug("doAction", logger.Int64("traceId", request.TraceId))
	result := action(request)
	logger.Debug("doAction result", logger.Int64("traceId", request.TraceId), logger.Any("result", result))
	return result
}

// deal with exception for business action
func doException(request model.IrisReq, writer http.ResponseWriter) {
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
		req := model.IrisReq{
			Request: request,
		}
		defer doException(req, writer)
		rs, _, err := filter.DoFilters(&req, filter.Pre)
		if !rs {
			panic(err)
		}
		result := doAction(req, action)
		rs, _, err = filter.DoFilters(&req, filter.Post)
		if !rs {
			panic(err)
		}
		doResponse(writer, result)
	}
	r.HandleFunc(url, wrapperAction).Methods(method)
}
