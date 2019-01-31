package service

import (
	"fmt"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/irishub-sync/logger"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	accountService = &AccountService{}

	blockService = &BlockService{}

	commonService = &CommonService{}

	proposalService = &ProposalService{}

	stakeService = &StakeService{}

	txService = &TxService{}
)

const (
	_ Module = iota
	Account
	Block
	Common
	Proposal
	Stake
	Tx
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
	case Stake:
		return stakeService
	case Tx:
		return txService
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

func queryPage(collation string, data interface{}, m map[string]interface{}, sort string, page, size int) model.Page {
	return orm.QueryList(collation, data, m, sort, page, size)
}

func QueryListField(collation string, selector bson.M, m map[string]interface{}, sort string, page, size int) (int, []map[string]interface{}) {
	var query = orm.MQuery{
		C:        collation,
		Q:        m,
		Selector: selector,
		Sort:     sort,
		Page:     page,
		Size:     size,
	}
	count, data, err := orm.QueryListField(query)
	if err != nil {
		logger.Error("QueryListField error", logger.Any("query", m), logger.String("err", err.Error()))
	}
	return count, data
}

func LimitQuery(collation string, selector bson.M, m map[string]interface{}, sort string, size int) ([]map[string]interface{}, error) {
	var query = orm.MQuery{
		C:        collation,
		Q:        m,
		Selector: selector,
		Sort:     sort,
		Size:     size,
	}
	data, err := orm.LimitQuery(query)
	if err != nil {
		logger.Error("QueryListField error", logger.Any("query", m), logger.String("err", err.Error()))
	}
	return data, err
}

func queryOne(collation string, data interface{}, m map[string]interface{}) error {
	return orm.QueryOne(collation, data, m)
}

func queryOneField(collation string, selector bson.M, m map[string]interface{}) map[string]interface{} {
	var query = orm.MQuery{
		C:        collation,
		Q:        m,
		Selector: selector,
	}
	return orm.QueryOneField(query)
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
