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

	RegisterApi(r, types.UrlRegisterQueryProposals, "GET", func(writer http.ResponseWriter, request *http.Request) {
		page, size := GetPage(request)

		result := service.GetProposal().QueryList(page, size)
		WriteResponse(writer, result)
	})

	return nil
}

func registerQueryProposal(r *mux.Router) error {

	RegisterApi(r, types.UrlRegisterQueryProposal, "GET", func(writer http.ResponseWriter, request *http.Request) {
		pid, err := strconv.Atoi(Var(request, "pid"))
		if err != nil {
			panic(types.ErrorCodeInValidParam)
			return
		}

		result := service.GetProposal().Query(pid)
		WriteResponse(writer, result)
	})

	return nil
}
