package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

func WriteResonse(writer http.ResponseWriter, data interface{}) {
	resultByte, err := json.Marshal(data)
	if err != nil {
		log.Println(fmt.Sprintf("json.Marshal failed:%s", err.Error()))
	}
	writer.Write(resultByte)
}
