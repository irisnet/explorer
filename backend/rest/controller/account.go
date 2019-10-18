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

// @Summary detail
// @Description get detail of account
// @Tags account
// @Accept  json
// @Produce  json
// @Param   address   path   string false    "account address"
// @Success 200 {object} vo.AccountVo	"success"
// @Router /api/account/{address} [get]
func registerQueryAccount(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryAccount, "GET", func(request vo.IrisReq) interface{} {
		address := Var(request, "address")
		account.SetTid(request.TraceId)
		result := account.Query(address)
		return result
	})

	return nil
}

// @Summary list
// @Description get list of account
// @Tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.AccountsInfoRespond	"success"
// @Router /api/accounts [get]
func registerQueryAccountList(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryAccounts, "GET", func(request vo.IrisReq) interface{} {
		return account.QueryRichList()
	})
	return nil
}

// @Summary list
// @Description get delegations of account
// @Tags account
// @Accept  json
// @Produce  json
// @Param   address   path   string false    "account address"
// @Success 200 {object} vo.AccountDelegationsRespond	"success"
// @Router /api/account/{address}/delegations [get]
func registerQueryAccountDelegations(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryAccountDelegations, "GET", func(request vo.IrisReq) interface{} {
		address := Var(request, "address")
		account.SetTid(request.TraceId)
		result := account.QueryDelegations(address)
		return result
	})

	return nil
}

// @Summary list
// @Description get unbondingdelegations of account
// @Tags account
// @Accept  json
// @Produce  json
// @Param   address   path   string false    "account address"
// @Success 200 {object} vo.AccountUnbondingDelegationsRespond	"success"
// @Router /api/account/{address}/unbonding_delegations [get]
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
// @Description get rewards of account
// @Tags account
// @Accept  json
// @Produce  json
// @Param   address   path   string false    "account address"
// @Success 200 {object} vo.AccountRewardsVo	"success"
// @Router /api/account/{address}/rewards [get]
func registerQueryAccountRewards(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryAccountRewards, "GET", func(request vo.IrisReq) interface{} {
		address := Var(request, "address")
		account.SetTid(request.TraceId)
		result := account.QueryRewards(address)
		return result
	})

	return nil
}