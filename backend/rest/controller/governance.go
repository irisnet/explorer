package controller

import (
	"strconv"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
)

func RegisterProposal(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryProposals,
		registerQueryProposal,
		registerQueryGovParams,
		registerQueryDepositAndVotingProposals,
		registerQueryProposalDepositorTxs,
		registerQueryProposalVoterTxs,
		registerQueryProposalDeposit,
		registerQueryProposalVoting,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

//
//type Gov struct {
//	*service.ProposalService
//	*service.GovParamsService
//}
//
//var gov = Gov{
//	service.Get(service.Proposal).(*service.ProposalService),
//	service.Get(service.GovParams).(*service.GovParamsService),
//}

func registerQueryProposals(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryProposals, "GET", func(request vo.IrisReq) interface{} {
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 10))

		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}
		result := gov.QueryList(page, size, istotal)
		return result
	})

	return nil
}

func registerQueryDepositAndVotingProposals(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryDepositVotingProposals, "GET", func(request vo.IrisReq) interface{} {
		need := QueryParam(request, "needMoniker")
		needMoniker := true
		if need == "false" {
			needMoniker = false
		}
		result := gov.QueryDepositAndVotingProposalList(needMoniker)
		return result
	})

	return nil
}

func registerQueryProposal(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryProposal, "GET", func(request vo.IrisReq) interface{} {
		pid, err := strconv.Atoi(Var(request, "pid"))
		if err != nil {
			panic(types.CodeInValidParam)
		}

		result := gov.Query(pid)
		return result
	})

	return nil
}

func registerQueryGovParams(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryParams, "GET", func(request vo.IrisReq) interface{} {
		return govparams.QueryAll()
	})
	return nil
}

func registerQueryProposalVoterTxs(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryProposalsVoterTxs, "GET", func(request vo.IrisReq) interface{} {
		id, err := strconv.ParseInt(Var(request, "id"), 10, 64)
		if err != nil {
			panic(types.CodeInValidParam)
		}
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 10))
		total := QueryParam(request, "total")
		voterType := QueryParam(request, "voterType")
		istotal := true
		if total == "false" {
			istotal = false
		}

		return gov.GetVoteTxs(id, page, size, istotal, voterType)
	})
	return nil
}

func registerQueryProposalDepositorTxs(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryProposalsDepositorTxs, "GET", func(request vo.IrisReq) interface{} {
		id, err := strconv.ParseInt(Var(request, "id"), 10, 64)
		if err != nil {
			panic(types.CodeInValidParam)
		}
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 10))

		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}
		return gov.GetDepositTxs(id, page, size, istotal)
	})
	return nil
}

func registerQueryProposalDeposit(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryProposalDeposit, "GET", func(request vo.IrisReq) interface{} {
		id, err := strconv.Atoi(Var(request, "id"))
		if err != nil {
			panic(types.CodeInValidParam)
		}

		result := gov.QueryDeposit(id)
		return result
	})
	return nil
}

func registerQueryProposalVoting(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryProposalVoting, "GET", func(request vo.IrisReq) interface{} {
		id, err := strconv.Atoi(Var(request, "id"))
		if err != nil {
			panic(types.CodeInValidParam)
		}

		result := gov.QueryVoting(id)
		return result
	})
	return nil
}
