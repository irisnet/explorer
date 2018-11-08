package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"io/ioutil"
	"log"
	"net/http"
)

const IPURI = "http://opendata.baidu.com/api.php?query=%s&resource_id=6006&ie=utf8&oe=utf8"

func RegisterQueryIp(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryIp, queryIp).Methods("POST")
	return nil
}

func queryIp(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var params map[string][]string
	json.Unmarshal(body, &params)
	ips := params["ipdata"]
	var ipMap = make([]string, len(ips))
	for i, ip := range ips {
		url := fmt.Sprintf(IPURI, ip)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("get request access err: %s", err.Error())
			return
		}
		body, _ := ioutil.ReadAll(resp.Body)
		ipMap[i] = string(body)
		fmt.Println(ipMap[i])
		//time.Sleep(100 * time.Millisecond)
	}
	res, err := json.Marshal(ipMap)
	if err == nil {
		w.Write(res)
	}
}
