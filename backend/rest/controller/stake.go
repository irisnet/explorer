package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/service"

	"github.com/irisnet/explorer/backend/types"
)

func RegisterStake(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryValidator,
		registerQueryRevokedValidator,
		registerQueryCandidates,
		registerQueryCandidate,
		registerQueryCandidateStatus,
		registerQueryCandidatesTop,
		registerQueryCandidateUptime,
		registerQueryCandidatePower,
		registerQueryChain,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

type Stake struct {
	*service.StakeService
}

var stake = Stake{
	service.Get(service.Stake).(*service.StakeService),
}

func registerQueryValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryValidator, "GET", func(request IrisReq) interface{} {
		page, size := GetPage(request)
		result := stake.QueryValidators(page, size)
		return result
	})

	return nil
}
func registerQueryRevokedValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryRevokedValidator, "GET", func(request IrisReq) interface{} {
		page, size := GetPage(request)
		result := stake.QueryRevokedValidator(page, size)
		return result
	})
	return nil
}
func registerQueryCandidates(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidates, "GET", func(request IrisReq) interface{} {
		page, size := GetPage(request)
		result := stake.QueryCandidates(page, size)
		return result
	})

	return nil
}

func registerQueryCandidatesTop(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidatesTop, "GET", func(request IrisReq) interface{} {
		result := stake.QueryCandidatesTopN()
		return result
	})

	return nil
}

func registerQueryCandidate(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidate, "GET", func(request IrisReq) interface{} {
		address := Var(request, "address")
		result := stake.QueryCandidate(address)
		return result
	})

	return nil
}

func registerQueryCandidateUptime(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidateUptime, "GET", func(request IrisReq) interface{} {
		address := Var(request, "address")
		category := Var(request, "category")

		result := stake.QueryCandidateUptime(address, category)
		return result
	})

	return nil
}

func registerQueryCandidatePower(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidatePower, "GET", func(request IrisReq) interface{} {
		address := Var(request, "address")
		category := Var(request, "category")

		result := stake.QueryCandidatePower(address, category)
		return result
	})
	return nil
}

func registerQueryCandidateStatus(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidateStatus, "GET", func(request IrisReq) interface{} {
		address := Var(request, "address")

		result := stake.QueryCandidateStatus(address)
		return result
	})

	return nil
}

func registerQueryChain(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryChain, "GET", func(request IrisReq) interface{} {
		result := stake.QueryChainStatus()
		return result
	})

	return nil
}
