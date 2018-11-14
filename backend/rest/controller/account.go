package controller

import (
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"net/http"

	"github.com/gorilla/mux"
)

// mux.Router registrars
func RegisterAccount(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		RegisterQueryAccount,
		RegisterQueryAllAccount,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func RegisterQueryAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryAccount, func(writer http.ResponseWriter, request *http.Request) {
		address := Var(request, "address")

		result, err := service.GetAccount().Query(address)
		if err.Success() {
			WriteResonse(writer, result)
			return
		}
		WriteResonse(writer, err)

	}).Methods("GET")
	return nil
}

func RegisterQueryAllAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryAllAccount, func(writer http.ResponseWriter, request *http.Request) {
		page, size := GetPage(request)

		result := service.GetAccount().QueryList(page, size)
		WriteResonse(writer, result)

	}).Methods("GET")
	return nil
}
