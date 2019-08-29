package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"

	"github.com/irisnet/explorer/backend/types"
)

func RegisterStake(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryValidator,
		registerQueryCandidateStatus,
		registerQueryCandidatesTop,
		registerGetValidators,
		registerUpdateValidatorIcons,
		registerGetValidator,
		registerQueryDelegationsByValidator,
		registerQueryUnbondingDelegationsByValidator,
		registerQueryRedelegationsByValidator,
		registerQueryVoterTxsByValidatorAddr,
		registerQueryDepositorTxsByValidatorAddr,
		registerQueryWithdrawAddrByValidatorAddr,
		registerQueryRewardsByValidatorAddr,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

//type Stake struct {
//	*service.ValidatorService
//}
//
//var stake = Stake{
//	service.Get(service.Validator).(*service.ValidatorService),
//}

func registerQueryWithdrawAddrByValidatorAddr(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryWithdrawAddrByValidatorAddr, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		validatorAddr := Var(request, "validatorAddr")

		return stake.GetWithdrawAddrByValidatorAddr(validatorAddr)
	})
	return nil
}

func registerQueryRewardsByValidatorAddr(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryCommissionRewardsByValidatorAddr, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		validatorAddr := Var(request, "validatorAddr")
		return stake.GetDistributionRewardsByValidatorAddr(validatorAddr)
	})
	return nil

}

func registerQueryVoterTxsByValidatorAddr(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryValidatorVoteByValidatorAddr, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}

		return stake.GetVoteTxsByValidatorAddr(validatorAddr, page, size, istotal)
	})
	return nil
}

func registerQueryDepositorTxsByValidatorAddr(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryDepositorTxsByValidatorAddr, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}

		return stake.GetDepositedTxByValidatorAddr(validatorAddr, page, size, istotal)
	})
	return nil
}

func registerQueryDelegationsByValidator(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryValidatorsDelegations, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		needpage := QueryParam(request, "needpage")
		validatorAddr := Var(request, "validatorAddr")
		if page > 0 {
			page = page - 1
		}
		var ispage bool
		if needpage == "false" {
			ispage = false
		} else {
			ispage = true
		}
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}

		return stake.GetDelegationsFromLcd(validatorAddr, page, size, ispage, istotal)
	})
	return nil
}

func registerQueryUnbondingDelegationsByValidator(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryValidatorUnbondingDelegations, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")
		if page > 0 {
			page = page - 1
		}

		return stake.GetUnbondingDelegationsFromLcd(validatorAddr, page, size)
	})
	return nil
}

func registerQueryRedelegationsByValidator(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryValidatorRedelegations, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")

		if page > 0 {
			page = page - 1
		}
		return stake.GetRedelegationsFromLcd(validatorAddr, page, size)
	})
	return nil

}

func registerGetValidators(r *mux.Router) error {
	doApi(r, types.UrlRegisterGetValidators, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 100))
		typ := QueryParam(request, "type")
		origin := QueryParam(request, "origin")
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}
		result := stake.GetValidators(typ, origin, page, size, istotal)
		return result
	})
	return nil
}
func registerUpdateValidatorIcons(r *mux.Router) error {
	doApi(r, types.UrlRegisterUpdateIcons, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		if err := stake.UpdateValidatorIcons(); err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse("0", "success", nil)
	})
	return nil
}

func registerGetValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterGetValidator, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")
		result := stake.GetValidatorDetail(address)
		return result
	})
	return nil
}
func registerQueryCandidatesTop(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidatesTop, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		result := stake.QueryCandidatesTopN()
		return result
	})

	return nil
}
func registerQueryValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidate, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")
		result := stake.QueryValidator(address)
		return result
	})

	return nil
}

func registerQueryCandidateStatus(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidateStatus, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")

		result := stake.QueryCandidateStatus(address)
		return result
	})

	return nil
}
