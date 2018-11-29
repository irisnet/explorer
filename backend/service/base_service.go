package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2"
)

var (
	baseService = &BaseService{}

	blockService = &BlockService{
		baseService,
	}

	commonService = &CommonService{
		baseService,
	}

	proposalService = &ProposalService{
		baseService,
	}

	stakeService = &StakeService{
		baseService,
	}

	txService = &TxService{
		baseService,
	}
)

type BaseService struct {
}

func (service *BaseService) QueryPage(collation string, data interface{}, m map[string]interface{}, sort string, page, size int) model.Page {
	return orm.QueryList(collation, data, m, sort, page, size)
}

func (service *BaseService) GetDb() *mgo.Database {
	return orm.GetDatabase()
}
