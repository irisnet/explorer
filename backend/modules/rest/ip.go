package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"io/ioutil"
)

const IPURI = "http://ip.taobao.com/service/getIpInfo.php?ip="

func RegisterQueryIp(r *mux.Router) error {
	r.HandleFunc("/api/ip/{ip}", queryIp).Methods("GET")
	return nil
}

func queryIp(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	ip := args["ip"]
	resp, err := http.Get(IPURI + ip)
	if err != nil {
		log.Printf("get request access err: %s", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
