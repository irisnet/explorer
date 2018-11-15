package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"net/http"
)

func RegisterTextSearch(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryText,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryText(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryText, "GET", func(writer http.ResponseWriter, request *http.Request) {
		text := Var(request, "text")

		result := service.GetCommon().QueryText(text)
		WriteResponse(writer, result)
	})

	return nil
}
