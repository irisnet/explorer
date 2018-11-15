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
	RegisterApi(r, types.UrlRegisterQueryAccount, "GET", func(writer http.ResponseWriter, request *http.Request) {
		address := Var(request, "address")

		result := service.GetAccount().Query(address)
		WriteResonse(writer, result)

	})

	return nil
}

func RegisterQueryAllAccount(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryAllAccount, "GET", func(writer http.ResponseWriter, request *http.Request) {
		page, size := GetPage(request)

		result := service.GetAccount().QueryList(page, size)
		WriteResonse(writer, result)
	})
	return nil
}
