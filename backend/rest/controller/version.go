package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/types"
	"net/http"
)

func RegisterQueryVersion(r *mux.Router) error {

	RegisterApi(r, types.UrlRegisterQueryApiVersion, "GET", func(writer http.ResponseWriter, request *http.Request) {
		WriteResponse(writer, map[string]string{"version": conf.Get().Server.ApiVersion})
	})

	return nil
}

func RegisterPing(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterPing, "GET", func(writer http.ResponseWriter, request *http.Request) {
	})
	return nil
}
