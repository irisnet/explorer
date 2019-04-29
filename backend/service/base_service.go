package service

import (
	"fmt"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
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
	InitVoteValidatorNum            = 0
	InitValidatorNum                = 0
	InitPrecommitVotingPower        = 0
	InitTotalVotingPower            = 0
	DefaultPageSize                 = 10
	DefaultPageNum                  = 1
	_                        Module = iota
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

func (base *BaseService) QueryBlackList(database *mgo.Database) map[string]document.BlackList {
	var blackListStore = database.C(document.CollectionNmBlackList)
	var blackList []document.BlackList
	var blackListMap = make(map[string]document.BlackList)
	if err := blackListStore.Find(nil).All(&blackList); err == nil {
		for _, v := range blackList {
			blackListMap[v.OperatorAddr] = v
		}
	}
	return blackListMap
}

func queryAll(collation string, selector, condition bson.M, sort string, size int, result interface{}) error {
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(collation).
		SetCondition(condition).
		SetSelector(selector).
		SetSort(sort).
		SetSize(size).
		SetResult(result)

	err := query.Exec()
	if err != nil {
		logger.Error("queryAll error", logger.Any("query", condition), logger.String("err", err.Error()))
	}
	return err
}

func queryOne(collation string, selector, condition bson.M, result interface{}) error {
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(collation).
		SetCondition(condition).
		SetSelector(selector).
		SetResult(result)

	err := query.Exec()
	if err != nil {
		logger.Error("queryOne", logger.Any("query", condition), logger.String("err", err.Error()))
	}
	return err
}

func pageQuery(collation string, selector, condition bson.M, sort string, page, size int, result interface{}) (int, error) {
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(collation).
		SetCondition(condition).
		SetSelector(selector).
		SetSort(sort).
		SetPage(page).
		SetSize(size).
		SetResult(result)

	cnt, err := query.ExecPage()
	if err != nil {
		logger.Error("pageQuery", logger.Any("query", condition), logger.String("err", err.Error()))
	}

	return cnt, err
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
