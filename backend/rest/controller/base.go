package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/logger"
	"net/http"
	"strconv"
	"time"
)

func GetString(request *http.Request, key string) (result string) {
	apiName := request.RequestURI

	request.ParseForm()
	if len(request.Form[key]) > 0 {
		result = request.Form[key][0]
		logger.Info("Api Param", logger.String("Api", apiName), logger.String(key, result))
	}
	return
}

func GetInt(request *http.Request, key string) (result int) {
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

func Var(request *http.Request, key string) (result string) {
	args := mux.Vars(request)
	result = args[key]
	return
}

func GetPage(r *http.Request) (int, int) {
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

func writeResponse(writer http.ResponseWriter, data interface{}) {
	var bz []byte
	//var resp = model.Response{
	//	Code: types.CodeSuccess.Code,
	//	Msg:  types.CodeSuccess.Msg,
	//}
	//
	//switch data.(type) {
	//case types.BizCode:
	//	err := data.(types.BizCode)
	//	resp.Code = err.Code
	//	resp.Msg = err.Msg
	//default:
	//	resp.Data = data
	//}
	//
	//bz, _ = json.Marshal(resp)

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

// 用户处理逻辑函数
type Action func(request *http.Request) interface{}

//封装用户处理逻辑函数，捕获panic异常，统一处理
func wrap(action Action) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiName := request.RequestURI
		api := logger.String("Api", apiName)

		defer func() {
			if r := recover(); r != nil {
				switch r.(type) {
				case types.BizCode:
					writeResponse(writer, r)
					break
				case error:
					err := r.(error)
					e := types.BizCode{
						Code: types.CodeUnKnown.Code,
						Msg:  err.Error(),
					}
					writeResponse(writer, e)
					break
				}
				logger.Error("API execute failed", api, logger.Any("errMsg", r))

			}
		}()

		logger.Info(fmt.Sprintf("========================[api]========================"))
		logger.Info("Value From Form", api, logger.Any("Params", request.Form))
		logger.Info("Value From Url", api, logger.Any("Params", mux.Vars(request)))

		start := time.Now()
		//执行用户的api逻辑
		result := action(request)
		end := time.Now()

		bizTime := end.Unix() - start.Unix()

		logger.Info("Api Result", api, logger.Any("result", result))
		logger.Info("Api Coast Time", api, logger.Int64("coast(s)", bizTime))

		if bizTime >= 3 {
			logger.Warn("api coast most time", api)
		}
		logger.Info(fmt.Sprintf("========================[api]========================"))

		writeResponse(writer, result)
	}
}

//处理api接口
// url : api路径
// method : api请求方法
// action : 用户处理逻辑
func doApi(r *mux.Router, url, method string, action Action) {
	wrapperAction := wrap(action)
	r.HandleFunc(url, wrapperAction).Methods(method)
}
