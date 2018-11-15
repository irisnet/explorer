package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"log"
	"net/http"
	"strconv"
)

func GetString(request *http.Request, key string) (result string) {
	apiName := request.RequestURI

	request.ParseForm()
	if len(request.Form[key]) > 0 {
		result = request.Form[key][0]
		log.Println(fmt.Sprintf("Api[%s] Param[%s=%s]", apiName, key, result))
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
		log.Println(fmt.Sprintf("%s is not int type", key))
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

func WriteResponse(writer http.ResponseWriter, data ...interface{}) {
	var resp = model.Response{
		Code: types.ErrorCodeSuccess.Code,
	}

	if len(data) == 2 {
		if succ, ok := data[1].(bool); ok && !succ {
			err := data[0].(types.Error)
			resp.Code = err.Code
			resp.Msg = err.Msg
		}
	}
	resp.Data = data[0]
	resultByte, err := json.Marshal(resp)
	if err != nil {
		log.Println(fmt.Sprintf("json.Marshal failed:%s", err.Error()))
	}
	writer.Write(resultByte)
}

// 用户处理逻辑函数
type Action func(writer http.ResponseWriter, request *http.Request)

//封装用户处理逻辑函数，捕获panic异常，统一处理
func wrap(action Action) Action {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				apiName := request.RequestURI
				switch r.(type) {
				case types.Error:
					WriteResponse(writer, r, false)
					break
				case error:
					err := r.(error)
					e := types.Error{
						Code: types.ErrorCodeUnKnown.Code,
						Msg:  err.Error(),
					}
					WriteResponse(writer, e, false)
					break
				}
				log.Println(fmt.Sprintf("API[%s] execute failed,%+v", apiName, r))
			}
		}()
		action(writer, request)
	}
}

//注册api接口
// url : api路径
// method : api请求方法
// action : 用户处理逻辑
func RegisterApi(r *mux.Router, url, method string, action Action) {
	wrapperAction := wrap(action)
	r.HandleFunc(url, wrapperAction).Methods(method)
}
