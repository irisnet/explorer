package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/version"
	"net/http"
)

func RegisterQueryVersion(r *mux.Router) error {
	r.HandleFunc("/version", func(writer http.ResponseWriter, request *http.Request) {
		WriteResonse(writer, map[string]string{"version": version.GetVersion()})
	}).Methods("GET")
	return nil
}

func RegisterPing(r *mux.Router) error {
	r.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		return
	}).Methods("GET")
	return nil
}
