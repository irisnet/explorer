package service

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"gopkg.in/mgo.v2/bson"
)

type DelegatorService struct {
	BaseService
}

func (service *DelegatorService) GetModule() Module {
	return Delegator
}

func (service *DelegatorService) GetTotalDeposits(delAddr string) float64 {
	delegatorStore := getDb().C(document.CollectionNmStakeRoleDelegator)
	defer delegatorStore.Database.Session.Close()

	query := bson.M{}
	query[document.Delegator_Field_Addres] = delAddr

	pipe := delegatorStore.Pipe(
		[]bson.M{
			{"$match": query},
			{"$group": bson.M{
				"_id":   document.Delegator_Field_Addres,
				"count": bson.M{"$sum": "$shares"},
			}},
		},
	)
	var result model.CountVo
	if err := pipe.One(&result); err != nil {
		logger.Warn("delegator address not exist", logger.String("delAddr", delAddr))
	}
	return result.Count
}
