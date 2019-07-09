package controller

import "github.com/irisnet/explorer/backend/service"

var (
	account   service.AccountService
	block     service.BlockService
	common    service.CommonService
	gov       service.ProposalService
	govparams service.GovParamsService
	stake     service.ValidatorService
	tx        service.TxService
)
