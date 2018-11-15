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
	RegisterApi(r, types.UrlRegisterQueryNodes, "GET", func(writer http.ResponseWriter, request *http.Request) {
		bz := utils.GetNodes()
		writer.Write(bz)
	})

	return nil
}

func RegisterQueryNodeLocation(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryIp, "POST", func(writer http.ResponseWriter, request *http.Request) {
		body, _ := ioutil.ReadAll(request.Body)
		var params map[string][]string
		json.Unmarshal(body, &params)
		ips := params["ipdata"]
		var ipMap = make([]string, len(ips))
		for i, ip := range ips {
			url := fmt.Sprintf(types.UrlNodeLocation, ip)
			resp, err := http.Get(url)
			if err != nil {
				panic(types.ErrorCodeExtSysFailed)
				return
			}
			body, _ := ioutil.ReadAll(resp.Body)
			ipMap[i] = string(body)
		}
		WriteResonse(writer, ipMap)
	})

	return nil
}
