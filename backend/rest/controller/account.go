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
		registerQueryAccount,
		registerQueryAllAccount,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

type AccountController struct {
	*service.AccountService
}

var account = AccountController{
	service.Get(service.Account).(*service.AccountService),
}

func registerQueryAccount(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryAccount, "GET", func(request *http.Request) interface{} {
		address := Var(request, "address")

		result := account.Query(address)
		return result
	})

	return nil
}

func registerQueryAllAccount(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryAllAccount, "GET", func(request *http.Request) interface{} {
		page, size := GetPage(request)

		result := account.QueryList(page, size)
		return result
	})
	return nil
}
