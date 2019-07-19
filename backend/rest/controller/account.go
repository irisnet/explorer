package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
)

// mux.Router registrars
func RegisterAccount(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryAccount,
		registerQueryAccountList,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

//type AccountController struct {
//	*service.AccountService
//}
//
//var account = AccountController{
//	service.Get(service.Account).(*service.AccountService),
//}

func registerQueryAccount(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryAccount, "GET", func(request model.IrisReq) interface{} {
		address := Var(request, "address")
		account.SetTid(request.TraceId)
		result := account.Query(address)
		return result
	})

	return nil
}

func registerQueryAccountList(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryAccounts, "GET", func(request model.IrisReq) interface{} {
		return account.QueryRichList()
	})
	return nil
}
