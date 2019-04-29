package controller

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
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

type Proposal struct {
	*service.ProposalService
}

var proposal = Proposal{
	service.Get(service.Proposal).(*service.ProposalService),
}

func registerQueryProposals(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryProposals, "GET", func(request model.IrisReq) interface{} {
		proposal.SetTid(request.TraceId)
		page, size := GetPage(request)

		result := proposal.QueryList(page, size)
		return result
	})

	return nil
}

func registerQueryProposal(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryProposal, "GET", func(request model.IrisReq) interface{} {
		proposal.SetTid(request.TraceId)
		pid, err := strconv.Atoi(Var(request, "pid"))
		if err != nil {
			panic(types.CodeInValidParam)
			return nil
		}

		result := proposal.Query(pid)
		return result
	})

	return nil
}
