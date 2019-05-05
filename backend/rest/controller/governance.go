package controller

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
)

func RegisterProposal(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryProposals,
		registerQueryProposal,
		registerQueryGovParams,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

type Gov struct {
	*service.ProposalService
	*service.GovParamsService
}

var gov = Gov{
	service.Get(service.Proposal).(*service.ProposalService),
	service.Get(service.GovParams).(*service.GovParamsService),
}

func registerQueryProposals(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryProposals, "GET", func(request model.IrisReq) interface{} {
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 10))

		result := gov.QueryList(page, size)
		return result
	})

	return nil
}

func registerQueryProposal(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryProposal, "GET", func(request model.IrisReq) interface{} {
		pid, err := strconv.Atoi(Var(request, "pid"))
		if err != nil {
			panic(types.CodeInValidParam)
			return nil
		}

		result := gov.Query(pid)
		return result
	})

	return nil
}

func registerQueryGovParams(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryGovParams, "GET", func(request model.IrisReq) interface{} {
		return gov.GovParamsService.QueryAll()
	})

	return nil
}
