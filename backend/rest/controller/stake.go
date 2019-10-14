package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"

	"github.com/irisnet/explorer/backend/types"
)

func RegisterStake(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerQueryValidator,
		registerQueryCandidateStatus,
		registerQueryCandidatesTop,
		registerGetValidators,
		registerUpdateValidatorIcons,
		registerGetValidator,
		registerQueryDelegationsByValidator,
		registerQueryUnbondingDelegationsByValidator,
		registerQueryRedelegationsByValidator,
		registerQueryVoterTxsByValidatorAddr,
		registerQueryDepositorTxsByValidatorAddr,
		registerQueryWithdrawAddrByValidatorAddr,
		registerQueryRewardsByValidatorAddr,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

//type Stake struct {
//	*service.ValidatorService
//}
//
//var stake = Stake{
//	service.Get(service.Validator).(*service.ValidatorService),
//}

// @Summary withdraw-addr
// @Description get withdraw-addr
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   validatorAddr    path   string true    "validatorAddr"
// @Success 200 {object} vo.WithdrawAddr	"success"
// @Router /api/stake/validators/{validatorAddr}/withdraw-addr [get]
func registerQueryWithdrawAddrByValidatorAddr(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryWithdrawAddrByValidatorAddr, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		validatorAddr := Var(request, "validatorAddr")

		return stake.GetWithdrawAddrByValidatorAddr(validatorAddr)
	})
	return nil
}

// @Summary commission-rewards
// @Description get commission-rewards
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   validatorAddr    path   string  true    "validatorAddr"
// @Success 200 {object} utils.CoinsAsStr	"success"
// @Router /api/stake/validators/{validatorAddr}/commission-rewards [get]
func registerQueryRewardsByValidatorAddr(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryCommissionRewardsByValidatorAddr, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		validatorAddr := Var(request, "validatorAddr")
		return stake.GetDistributionRewardsByValidatorAddr(validatorAddr)
	})
	return nil

}

// @Summary vote
// @Description get vote
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(5)
// @Param   validatorAddr    path   string  true    "validatorAddr"
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.ValidatorVotePage	"success"
// @Router /api/stake/validators/{validatorAddr}/vote [get]
func registerQueryVoterTxsByValidatorAddr(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryValidatorVoteByValidatorAddr, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}

		return stake.GetVoteTxsByValidatorAddr(validatorAddr, page, size, istotal)
	})
	return nil
}

// @Summary depositor_txs
// @Description get depositor txs
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(5)
// @Param   validatorAddr    path   string  true    "validatorAddr"
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.ValidatorDepositTxPage	"success"
// @Router /api/stake/validators/{validatorAddr}/depositor_txs [get]
func registerQueryDepositorTxsByValidatorAddr(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryDepositorTxsByValidatorAddr, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}

		return stake.GetDepositedTxByValidatorAddr(validatorAddr, page, size, istotal)
	})
	return nil
}

// @Summary delegations
// @Description get delegations
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(5)
// @Param   validatorAddr    path   string  true    "validatorAddr"
// @Param   needpage   query   bool false    "needpage" Enums(true,false)
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} vo.DelegationsPage	"success"
// @Router /api/stake/validators/{validatorAddr}/delegations [get]
func registerQueryDelegationsByValidator(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryValidatorsDelegations, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		needpage := QueryParam(request, "needpage")
		validatorAddr := Var(request, "validatorAddr")
		if page > 0 {
			page = page - 1
		}
		var ispage bool
		if needpage == "false" {
			ispage = false
		} else {
			ispage = true
		}
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}

		return stake.GetDelegationsFromLcd(validatorAddr, page, size, ispage, istotal)
	})
	return nil
}

// @Summary unbondingdelegations
// @Description get unbonding delegations
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(5)
// @Param   validatorAddr    path   string  true    "validatorAddr"
// @Success 200 {object} vo.UnbondingDelegationsPage	"success"
// @Router /api/stake/validators/{validatorAddr}/unbonding-delegations [get]
func registerQueryUnbondingDelegationsByValidator(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryValidatorUnbondingDelegations, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")
		if page > 0 {
			page = page - 1
		}

		return stake.GetUnbondingDelegationsFromLcd(validatorAddr, page, size)
	})
	return nil
}

// @Summary redelegations
// @Description get redelegations
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(5)
// @Param   validatorAddr    path   string  true    "validatorAddr"
// @Success 200 {object} vo.RedelegationPage	"success"
// @Router /api/stake/validators/{validatorAddr}/redelegations [get]
func registerQueryRedelegationsByValidator(r *mux.Router) error {

	doApi(r, types.UrlRegisterQueryValidatorRedelegations, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 5))
		validatorAddr := Var(request, "validatorAddr")

		if page > 0 {
			page = page - 1
		}
		return stake.GetRedelegationsFromLcd(validatorAddr, page, size)
	})
	return nil

}

// @Summary list
// @Description get validators
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   page    query   int true    "page num" Default(1)
// @Param   size   query   int true    "page size" Default(100)
// @Param   type   query   string false    "validators status" Enums(validator,candidate,jailed)
// @Param   origin   query   string false    "origin" Enums(browser,app)
// @Param   total   query   bool false    "total" Enums(true,false)
// @Success 200 {object} lcd.ValidatorsVoRespond	"success"
// @Router /api/stake/validators [get]
func registerGetValidators(r *mux.Router) error {
	doApi(r, types.UrlRegisterGetValidators, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		page := int(utils.ParseIntWithDefault(QueryParam(request, "page"), 1))
		size := int(utils.ParseIntWithDefault(QueryParam(request, "size"), 100))
		typ := QueryParam(request, "type")
		origin := QueryParam(request, "origin")
		total := QueryParam(request, "total")
		istotal := true
		if total == "false" {
			istotal = false
		}
		result := stake.GetValidators(typ, origin, page, size, istotal)
		return result
	})
	return nil
}

// @Summary update_icons
// @Description update validator icons
// @Tags stake
// @Accept  json
// @Produce  json
// @Success 200 {object} nil  "success"
// @Router /api/stake/validators/update_icons [get]
func registerUpdateValidatorIcons(r *mux.Router) error {
	doApi(r, types.UrlRegisterUpdateIcons, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		if err := stake.UpdateValidatorIcons(); err != nil {
			return vo.NewResponse("-1", err.Error(), nil)
		}
		return vo.NewResponse("0", "success", nil)
	})
	return nil
}

// @Summary validator detail
// @Description get validator detail
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   address   path   string  false    "address"
// @Success 200 {object} vo.ValidatorForDetail  "success"
// @Router /api/stake/validators/{address} [get]
func registerGetValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterGetValidator, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")
		result := stake.GetValidatorDetail(address)
		return result
	})
	return nil
}

// @Summary validator candidatesTop
// @Description get validator candidatesTop
// @Tags stake
// @Accept  json
// @Produce  json
// @Success 200 {object} vo.ValDetailVo  "success"
// @Router /api/stake/candidatesTop [get]
func registerQueryCandidatesTop(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidatesTop, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		result := stake.QueryCandidatesTopN()
		return result
	})

	return nil
}

// @Summary candidate detail
// @Description get validator candidate detail
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   address   path   string  false    "address"
// @Success 200 {object} vo.CandidatesInfoVo  "success"
// @Router /api/stake/candidate/{address} [get]
func registerQueryValidator(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidate, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")
		result := stake.QueryValidator(address)
		return result
	})

	return nil
}

// @Summary validator status
// @Description get validator status
// @Tags stake
// @Accept  json
// @Produce  json
// @Param   address   path   string  false    "address"
// @Success 200 {object} vo.ValStatus  "success"
// @Router /api/stake/candidate/{address}/status [get]
func registerQueryCandidateStatus(r *mux.Router) error {
	doApi(r, types.UrlRegisterQueryCandidateStatus, "GET", func(request vo.IrisReq) interface{} {
		stake.SetTid(request.TraceId)
		address := Var(request, "address")

		result := stake.QueryCandidateStatus(address)
		return result
	})

	return nil
}
