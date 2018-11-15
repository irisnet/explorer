package rest

import (
	"github.com/irisnet/explorer/backend/service"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"

	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
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
	r.HandleFunc(types.UrlRegisterQueryValidator, func(writer http.ResponseWriter, request *http.Request) {
		page, size := utils.GetPage(request)
		result := service.GetStake().QueryValidators(page, size)
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}
	}).Methods("GET")
	return nil
}
func registerQueryRevokedValidator(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryRevokedValidator, func(writer http.ResponseWriter, request *http.Request) {
		page, size := utils.GetPage(request)
		result := service.GetStake().QueryRevokedValidator(page, size)
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}
	}).Methods("GET")
	return nil
}
func registerQueryCandidates(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryCandidates, func(writer http.ResponseWriter, request *http.Request) {
		page, size := utils.GetPage(request)
		result := service.GetStake().QueryCandidates(page, size)
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}
	}).Methods("GET")
	return nil
}

func registerQueryCandidatesTop(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryCandidatesTop, func(writer http.ResponseWriter, request *http.Request) {
		result := service.GetStake().QueryCandidatesTopN()
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}
	}).Methods("GET")
	return nil
}

func registerQueryCandidate(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryCandidate, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		address := args["address"]
		result := service.GetStake().QueryCandidate(address)
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}
	}).Methods("GET")
	return nil
}

func registerQueryCandidateUptime(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryCandidateUptime, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		address := args["address"]
		category := args["category"]

		result := service.GetStake().QueryCandidateUptime(address, category)
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}

	}).Methods("GET")
	return nil
}

func registerQueryCandidatePower(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryCandidatePower, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		address := args["address"]
		category := args["category"]

		result := service.GetStake().QueryCandidatePower(address, category)
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}
	}).Methods("GET")
	return nil
}

func registerQueryCandidateStatus(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryCandidateStatus, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		address := args["address"]

		result := service.GetStake().QueryCandidateStatus(address)
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}
	}).Methods("GET")
	return nil
}

func registerQueryChain(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryChain, func(writer http.ResponseWriter, request *http.Request) {
		result := service.GetStake().QueryChainStatus()
		resp, err := json.Marshal(result)
		if err == nil {
			writer.Write(resp)
		}
	}).Methods("GET")
	return nil
}
