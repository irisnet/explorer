package service

import (
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type DelegatorService struct {
	BaseService
}

func (service *DelegatorService) GetModule() Module {
	return Delegator
}

func (service *DelegatorService) GetDeposits(delAddr string) (coin document.Coin) {
	var dbInstance = getDb()
	defer dbInstance.Session.Close()

	var delegatorStore = dbInstance.C(document.CollectionNmStakeRoleDelegator)
	var delegations []document.Delegator

	err := delegatorStore.Find(bson.M{document.Delegator_Field_Addres: delAddr}).All(&delegations)
	if err != nil {
		logger.Warn("delegator address not exist", logger.String("delAddr", delAddr))
		return
	}

	var delegationMap = make(map[string]document.Delegator, len(delegations))
	var valAddr []string
	for _, d := range delegations {
		delegationMap[d.ValidatorAddr] = d
		valAddr = append(valAddr, d.ValidatorAddr)
	}

	var validatorStore = dbInstance.C(document.CollectionNmStakeRoleCandidate)
	var validators []document.Candidate

	err = validatorStore.Find(bson.M{"$in": valAddr}).All(&validators)
	if err != nil {
		logger.Error("validator not found", logger.Any("valAddrs", valAddr))
		return
	}

	var totalAmt float64
	for _, v := range validators {
		delegation := delegationMap[v.Address]
		rate := v.Tokens / v.DelegatorShares
		totalAmt += delegation.Shares * rate
	}
	return document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: totalAmt,
	}
}
