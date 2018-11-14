package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"net/http"
	"strconv"
)

func RegisterProposal(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryProposals,
		registerQueryProposal,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerQueryProposals(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryProposals, func(writer http.ResponseWriter, request *http.Request) {
		page, size := utils.GetPage(request)

		result := service.GetProposal().QueryList(page, size)
		resultByte, _ := json.Marshal(result)
		writer.Write(resultByte)
	}).Methods("GET")
	return nil
}

func registerQueryProposal(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryProposal, func(writer http.ResponseWriter, request *http.Request) {
		args := mux.Vars(request)
		pid, _ := strconv.Atoi(args["pid"])

		result := service.GetProposal().Query(pid)
		resultByte, _ := json.Marshal(result)
		writer.Write(resultByte)
	}).Methods("GET")
	return nil
}
