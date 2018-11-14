package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
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
		page, size := GetPage(request)

		result := service.GetProposal().QueryList(page, size)
		WriteResonse(writer, result)
	}).Methods("GET")
	return nil
}

func registerQueryProposal(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryProposal, func(writer http.ResponseWriter, request *http.Request) {
		pid, err := strconv.Atoi(Var(request, "pid"))
		if err != nil {
			WriteResonse(writer, types.ErrorCodeInValidParam)
			return
		}

		result, error := service.GetProposal().Query(pid)

		if error.Success() {
			WriteResonse(writer, result)
			return
		}
		WriteResonse(writer, err)
	}).Methods("GET")
	return nil
}
