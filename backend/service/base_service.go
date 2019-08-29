package service

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"go.uber.org/zap"
)

var (
	accountService   = &AccountService{}
	blockService     = &BlockService{}
	commonService    = &CommonService{}
	proposalService  = &ProposalService{}
	txService        = &TxService{}
	delegatorService = &DelegatorService{}
	govParamsService = &GovParamsService{}
	validatorService = &ValidatorService{}
	assetsService    = &AssetsService{}
)

const (
	_ Module = iota
	Account
	Block
	Common
	Proposal
	Tx
	Delegator
	GovParams
	Validator
	Asset
)

type Module int

func Get(m Module) Service {
	switch m {
	case Account:
		return accountService
	case Block:
		return blockService
	case Common:
		return commonService
	case Proposal:
		return proposalService
	// case Candidate:
	// 	return stakeService
	case Tx:
		return txService
	case Delegator:
		return delegatorService
	case GovParams:
		return govParamsService
	case Validator:
		return validatorService
	case Asset:
		return assetsService
	}
	return nil
}

type Service interface {
	GetModule() Module
}

type BaseService struct {
	tid string
}

func (base *BaseService) SetTid(traceId string) {
	base.tid = traceId
}

func (base *BaseService) GetTid() string {
	return base.tid
}

func (base *BaseService) GetTraceLog() zap.Field {
	return logger.String("traceId", base.GetTid())
}

func (b *BaseService) QueryBlackList() map[string]document.BlackList {
	return document.BlackList{}.QueryBlackList()
}
