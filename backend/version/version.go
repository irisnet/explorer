package version

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

const Version = "0.1.0"

func RegisterQueryVersion(r *mux.Router) error {
	r.HandleFunc("/version", queryVersion).Methods("GET")
	return nil
}

func queryVersion(w http.ResponseWriter, r *http.Request) {
	accByte, _ := json.Marshal(map[string]string{"version": Version})
	w.Write(accByte)
}
