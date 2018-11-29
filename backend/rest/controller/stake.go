package controller

import (
	"github.com/irisnet/explorer/backend/service"
	"net/http"

	"github.com/gorilla/mux"

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

func registerQueryValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryValidator, "GET", func(request *http.Request) interface{} {
		page, size := GetPage(request)
		result := service.GetStake().QueryValidators(page, size)
		return result
	})

	return nil
}
func registerQueryRevokedValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryRevokedValidator, "GET", func(request *http.Request) interface{} {
		page, size := GetPage(request)
		result := service.GetStake().QueryRevokedValidator(page, size)
		return result
	})
	return nil
}
func registerQueryCandidates(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidates, "GET", func(request *http.Request) interface{} {
		page, size := GetPage(request)
		result := service.GetStake().QueryCandidates(page, size)
		return result
	})

	return nil
}

func registerQueryCandidatesTop(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidatesTop, "GET", func(request *http.Request) interface{} {
		result := service.GetStake().QueryCandidatesTopN()
		return result
	})

	return nil
}

func registerQueryCandidate(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidate, "GET", func(request *http.Request) interface{} {
		address := Var(request, "address")
		result := service.GetStake().QueryCandidate(address)
		return result
	})

	return nil
}

func registerQueryCandidateUptime(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidateUptime, "GET", func(request *http.Request) interface{} {
		address := Var(request, "address")
		category := Var(request, "category")

		result := service.GetStake().QueryCandidateUptime(address, category)
		return result
	})

	return nil
}

func registerQueryCandidatePower(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidatePower, "GET", func(request *http.Request) interface{} {
		address := Var(request, "address")
		category := Var(request, "category")

		result := service.GetStake().QueryCandidatePower(address, category)
		return result
	})
	return nil
}

func registerQueryCandidateStatus(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidateStatus, "GET", func(request *http.Request) interface{} {
		address := Var(request, "address")

		result := service.GetStake().QueryCandidateStatus(address)
		return result
	})

	return nil
}

func registerQueryChain(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryChain, "GET", func(request *http.Request) interface{} {
		result := service.GetStake().QueryChainStatus()
		return result
	})

	return nil
}
