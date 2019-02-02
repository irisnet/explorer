package service

import (
	"github.com/irisnet/explorer/backend/conf"
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

func (service *DelegatorService) QueryDelegation(valAddr string) (info ValInfo) {
	var dbInstance = getDb()
	defer dbInstance.Session.Close()

	var delegatorStore = dbInstance.C(document.CollectionNmStakeRoleDelegator)

	// query delegation info
	var accAddr = utils.Convert(conf.Get().Hub.Prefix.AccAddr, valAddr)
	var delegations []document.Delegator
	var selfBondShares float64
	var delegatedShares float64

	var query = bson.M{}
	query[document.Delegator_Field_ValidatorAddr] = valAddr
	err := delegatorStore.Find(query).All(&delegations)
	if err != nil {
		logger.Warn("validator not exist", logger.String("valAddr", valAddr))
		return
	}

	for _, d := range delegations {
		if d.Address == accAddr {
			selfBondShares = d.Shares
		} else {
			delegatedShares += d.Shares
		}
	}

	//query validator info
	var validatorStore = dbInstance.C(document.CollectionNmStakeRoleCandidate)
	var validator document.Candidate
	err = validatorStore.Find(bson.M{document.Candidate_Field_Address: valAddr}).One(&validator)
	if err != nil {
		logger.Error("validator not found", logger.Any("valAddr", valAddr))
		return
	}

	rate := validator.Tokens / validator.DelegatorShares

	selfBond := document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: selfBondShares * rate,
	}

	delegated := document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: delegatedShares * rate,
	}

	return ValInfo{
		valAddr:   valAddr,
		selfBond:  selfBond,
		delegated: delegated,
	}

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
	var valAddrs []string
	for _, d := range delegations {
		delegationMap[d.ValidatorAddr] = d
		valAddrs = append(valAddrs, d.ValidatorAddr)
	}

	var validatorStore = dbInstance.C(document.CollectionNmStakeRoleCandidate)
	var validators []document.Candidate

	var query = bson.M{}
	query[document.Candidate_Field_Address] = bson.M{
		"$in": valAddrs,
	}
	err = validatorStore.Find(query).All(&validators)
	if err != nil {
		logger.Error("validator not found", logger.Any("valAddrs", valAddrs))
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

type ValInfo struct {
	valAddr   string
	selfBond  document.Coin
	delegated document.Coin
}
