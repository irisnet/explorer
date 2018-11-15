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
	RegisterApi(r, types.UrlRegisterQueryValidator, "GET", func(writer http.ResponseWriter, request *http.Request) {
		page, size := GetPage(request)
		result := service.GetStake().QueryValidators(page, size)
		WriteResonse(writer, result)
	})

	return nil
}
func registerQueryRevokedValidator(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryRevokedValidator, "GET", func(writer http.ResponseWriter, request *http.Request) {
		page, size := GetPage(request)
		result := service.GetStake().QueryRevokedValidator(page, size)
		WriteResonse(writer, result)
	})
	return nil
}
func registerQueryCandidates(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryCandidates, "GET", func(writer http.ResponseWriter, request *http.Request) {
		page, size := GetPage(request)
		result := service.GetStake().QueryCandidates(page, size)
		WriteResonse(writer, result)
	})

	return nil
}

func registerQueryCandidatesTop(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryCandidatesTop, "GET", func(writer http.ResponseWriter, request *http.Request) {
		result := service.GetStake().QueryCandidatesTopN()
		WriteResonse(writer, result)
	})

	return nil
}

func registerQueryCandidate(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryCandidate, "GET", func(writer http.ResponseWriter, request *http.Request) {
		address := Var(request, "address")
		result := service.GetStake().QueryCandidate(address)
		WriteResonse(writer, result)
	})

	return nil
}

func registerQueryCandidateUptime(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryCandidateUptime, "GET", func(writer http.ResponseWriter, request *http.Request) {
		address := Var(request, "address")
		category := Var(request, "category")

		result := service.GetStake().QueryCandidateUptime(address, category)
		WriteResonse(writer, result)
	})

	return nil
}

func registerQueryCandidatePower(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryCandidatePower, "GET", func(writer http.ResponseWriter, request *http.Request) {
		address := Var(request, "address")
		category := Var(request, "category")

		result := service.GetStake().QueryCandidatePower(address, category)
		WriteResonse(writer, result)
	})
	return nil
}

func registerQueryCandidateStatus(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryCandidateStatus, "GET", func(writer http.ResponseWriter, request *http.Request) {
		address := Var(request, "address")

		result := service.GetStake().QueryCandidateStatus(address)
		WriteResonse(writer, result)
	})

	return nil
}

func registerQueryChain(r *mux.Router) error {
	RegisterApi(r, types.UrlRegisterQueryChain, "GET", func(writer http.ResponseWriter, request *http.Request) {
		result := service.GetStake().QueryChainStatus()
		WriteResonse(writer, result)
	})

	return nil
}
