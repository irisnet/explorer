package service

import (
	"fmt"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2"
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

func queryPage(collation string, data interface{}, m map[string]interface{}, sort string, page, size int) model.Page {
	return orm.QueryList(collation, data, m, sort, page, size)
}

func queryOne(collation string, data interface{}, m map[string]interface{}) error {
	return orm.QueryOne(collation, data, m)
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
