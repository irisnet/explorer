package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/utils"

	"github.com/irisnet/explorer/backend/types"
)

func RegisterStake(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryValidator,
		registerQueryCandidateStatus,
		registerQueryCandidatesTop,
		registerQueryCandidateUptime,
		registerQueryCandidatePower,
		registerGetValidators,
		registerGetValidator,
		registerQueryDelegationsByValidator,
		registerQueryUnbondingDelegationsByValidator,
		registerQueryRedelegationsByValidator,
		registerQueryVoterTxsByValidatorAddr,
		registerQueryDepositorTxsByValidatorAddr,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

type Stake struct {
	*service.ValidatorService
}

var stake = Stake{
	service.Get(service.Validator).(*service.ValidatorService),
}

func registerQueryVoterTxsByValidatorAddr(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryValidatorVoteByValidatorAddr, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")

		if page > 0 {
			page = page - 1
		}

		return stake.GetVoteTxsByValidatorAddr(validatorAddr, page, size)
	})
	return nil
}

func registerQueryDepositorTxsByValidatorAddr(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryDepositorTxsByValidatorAddr, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")
		if page > 0 {
			page = page - 1
		}

		return stake.GetDepositedTxByValidatorAddr(validatorAddr, page, size)
	})
	return nil
}

func registerQueryDelegationsByValidator(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryValidatorsDelegations, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")
		if page > 0 {
			page = page - 1
		}

		return stake.GetDelegationsFromLcd(validatorAddr, page, size)
	})
	return nil
}

func registerQueryUnbondingDelegationsByValidator(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryValidatorUnbondingDelegations, "GET", func(request model.IrisReq) interface{} {
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

	doApi(r, types.UrlRegisterQueryValidatorRedelegations, "GET", func(request model.IrisReq) interface{} {
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
	doApi(r, types.UrlRegisterGetValidators, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 100))
		typ := QueryParam(request, "type")
		origin := QueryParam(request, "origin")
		result := stake.GetValidators(typ, origin, page, size)
		return result
	})
	return nil
}
func registerGetValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterGetValidator, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")
		result := stake.GetValidatorDetail(address)
		return result
	})
	return nil
}
func registerQueryCandidatesTop(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidatesTop, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		result := stake.QueryCandidatesTopN()
		return result
	})

	return nil
}
func registerQueryValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidate, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")
		result := stake.QueryValidator(address)
		return result
	})

	return nil
}
func registerQueryCandidateUptime(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidateUptime, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")
		category := Var(request, "category")

		result := stake.QueryCandidateUptime(address, category)
		return result
	})

	return nil
}
func registerQueryCandidatePower(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidatePower, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")
		category := Var(request, "category")

		result := stake.QueryCandidatePower(address, category)
		return result
	})
	return nil
}
func registerQueryCandidateStatus(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidateStatus, "GET", func(request model.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")

		result := stake.QueryCandidateStatus(address)
		return result
	})

	return nil
}
