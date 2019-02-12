package service

import (
	"fmt"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	accountService = &AccountService{}

	blockService = &BlockService{}

	commonService = &CommonService{}

	proposalService = &ProposalService{}

	stakeService = &CandidateService{}

	txService        = &TxService{}
	delegatorService = &DelegatorService{}
)

const (
	_ Module = iota
	Account
	Block
	Common
	Proposal
	Candidate
	Tx
	Delegator
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
	case Candidate:
		return stakeService
	case Tx:
		return txService
	case Delegator:
		return delegatorService
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

func queryRows(collation string, data interface{}, m map[string]interface{}, sort string, page, size int) model.PageVo {
	return orm.QueryRows(collation, data, m, sort, page, size)
}

func limitQuery(collation string, selector bson.M, m map[string]interface{}, sort string, size int, result interface{}) error {
	var query = orm.MQuery{
		C:        collation,
		Q:        m,
		Selector: selector,
		Result:   result,
		Sort:     sort,
		Size:     size,
	}
	err := orm.All(query)
	if err != nil {
		logger.Error("limitQuery error", logger.Any("query", m), logger.String("err", err.Error()))
	}
	return err
}

func one(collation string, selector bson.M, m map[string]interface{}, result interface{}) error {
	var query = orm.MQuery{
		C:        collation,
		Q:        m,
		Selector: selector,
		Result:   result,
	}
	return orm.One(query)
}

func all(collation string, selector bson.M, m map[string]interface{}, sort string, page, size int, result interface{}) (int, error) {
	var query = orm.MQuery{
		C:        collation,
		Q:        m,
		Selector: selector,
		Result:   result,
		Sort:     sort,
		Page:     page,
		Size:     size,
	}
	return orm.AllWithCount(query)
}

func getDb() *mgo.Database {
	return orm.GetDatabase()
}

func desc(field string) string {
	return fmt.Sprintf("-%s", field)
}

func asc(field string) string {
	return fmt.Sprintf("%s", field)
}
