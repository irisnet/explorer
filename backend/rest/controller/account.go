package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
)

// mux.Router registrars
func RegisterAccount(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryAccount,
		registerQueryAccountList,
		registerQueryAccountDelegations,
		registerQueryAccountUnbondingDelegations,
		registerQueryAccountRewards,
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

	doApi(r, types.UrlRegisterQueryAccount, "GET", func(request vo.IrisReq) interface{} {
		address := Var(request, "address")
		account.SetTid(request.TraceId)
		result := account.Query(address)
		return result
	})

	return nil
}

func registerQueryAccountList(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryAccounts, "GET", func(request vo.IrisReq) interface{} {
		return account.QueryRichList()
	})
	return nil
}

func registerQueryAccountDelegations(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryAccountDelegations, "GET", func(request vo.IrisReq) interface{} {
		address := Var(request, "address")
		account.SetTid(request.TraceId)
		result := account.QueryDelegations(address)
		return result
	})

	return nil
}

func registerQueryAccountUnbondingDelegations(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryAccountUnbondingDelegations, "GET", func(request vo.IrisReq) interface{} {
		address := Var(request, "address")
		account.SetTid(request.TraceId)
		result := account.QueryUnbondingDelegations(address)
		return result
	})

	return nil
}

// @Summary list
// @Description get list of notice
// @Tags account
// @Accept  json
// @Produce  json
// @Param   account   query   string false    "notice status" Enums(unpublished, published, expired)
// @Success 200 {object} vo.AccountRewardsVo	"success"
// @Router /account/{address}/rewards [get]
func registerQueryAccountRewards(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryAccountRewards, "GET", func(request vo.IrisReq) interface{} {
		address := Var(request, "address")
		account.SetTid(request.TraceId)
		result := account.QueryRewards(address)
		return result
	})

	return nil
}