package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"io/ioutil"
	"net/http"
)

func RegisterNodes(r *mux.Router) error {
	apis := []func(*mux.Router) error{
		RegisterQueryNodes,
		RegisterQueryNodeLocation,
	}

	for _, fn := range apis {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func RegisterQueryNodes(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryNodes, func(w http.ResponseWriter, request *http.Request) {
		bz := utils.GetNodes()
		w.Write(bz)
	}).Methods("GET")
	return nil
}

func RegisterQueryNodeLocation(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryIp, func(writer http.ResponseWriter, request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)
		var params map[string][]string
		json.Unmarshal(body, &params)
		ips := params["ipdata"]
		var ipMap = make([]string, len(ips))
		var error types.Error
		for i, ip := range ips {
			url := fmt.Sprintf(types.UrlNodeLocation, ip)
			resp, err := http.Get(url)
			if err != nil {
				error = types.ErrorCodeExtSysFailed
				return
			}
			body, _ := ioutil.ReadAll(resp.Body)
			ipMap[i] = string(body)
		}
		if error.Success() {
			WriteResonse(writer, ipMap)
		} else {
			WriteResonse(writer, error)
		}

	}).Methods("POST")
	return nil
}
