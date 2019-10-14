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

// @Summary list
// @Description get proposals
// @Tags gov
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(10)
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.PageVo	"success"
// @Router /api/gov/proposals [get]
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

// @Summary deposit_voting_proposals
// @Description get deposit_voting_proposals
// @Tags gov
// @Accept  json
// @Produce  json
// @Param   needMoniker    query   string true    "needMoniker"
// @Success 200 {object} vo.ProposalNewStyleResponse    "success"
// @Router /api/gov/deposit_voting_proposals [get]
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

// @Summary detail
// @Description get proposal detail
// @Tags gov
// @Accept  json
// @Produce  json
// @Param   pid    path   string true    "proposal id"
// @Success 200 {object} vo.ProposalInfoVo    "success"
// @Router /api/gov/proposals/{pid} [get]
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

// @Summary params
// @Description get gov params
// @Tags gov
// @Accept  json
// @Produce  json
// @Success 200 {object} document.GovParamsList   "success"
// @Router /api/params [get]
func registerQueryGovParams(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryParams, "GET", func(request vo.IrisReq) interface{} {
		return govparams.QueryAll()
	})
	return nil
}

// @Summary voter_txs
// @Description get voter txs
// @Tags gov
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(10)
// @Param   id    path   string true    "proposal id"
// @Param   voterType   query   string false    "voter type" Enums(validator,delegator)
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.GetVoteTxResponse   "success"
// @Router /api/gov/proposals/{id}/voter_txs [get]
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

// @Summary depositor_txs
// @Description get depositor txs
// @Tags gov
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(10)
// @Param   id    path   string true    "proposal id"
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.TxPage   "success"
// @Router /api/gov/proposals/{id}/depositor_txs [get]
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

// @Summary deposit detail
// @Description get deposit detail
// @Tags gov
// @Accept  json
// @Produce  json
// @Param   id    path   string true    "proposal id"
// @Success 200 {object} vo.ProposalNewStyle   "success"
// @Router /api/gov/proposal_deposit/{id} [get]
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

// @Summary proposal voting
// @Description get proposal voting
// @Tags gov
// @Accept  json
// @Produce  json
// @Param   id    path   string true    "proposal id"
// @Success 200 {object} vo.ProposalNewStyle   "success"
// @Router /api/gov/proposal_voting/{id} [get]
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
