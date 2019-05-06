package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/shopspring/decimal"
	"gopkg.in/mgo.v2/bson"
)

type DelegatorService struct {
	BaseService
}

func (service *DelegatorService) GetModule() Module {
	return Delegator
}

func (service *DelegatorService) QueryDelegation(valAddr string) (document.Coin, document.Coin) {

	var accAddr = utils.Convert(conf.Get().Hub.Prefix.AccAddr, valAddr)

	var validator document.Validator
	var selector = bson.M{
		"tokens":           1,
		"delegator_shares": 1}

	err := queryOne(document.CollectionNmValidator, selector, bson.M{document.ValidatorFieldOperatorAddress: valAddr}, &validator)

	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return document.Coin{}, document.Coin{}
	}

	delegations := lcd.DelegationByValidator(valAddr)
	var selfBondShares decimal.Decimal
	var delegatedShares decimal.Decimal
	for _, d := range delegations {
		tmp, err := decimal.NewFromString(d.Shares)
		if err != nil {
			logger.Error("convert string to decimal err", logger.Any("err", err.Error()))
			continue
		}
		if d.DelegatorAddr == accAddr {
			selfBondShares = selfBondShares.Add(tmp)
		} else {
			delegatedShares = delegatedShares.Add(tmp)
		}
	}

	valTokensDec, err := decimal.NewFromString(validator.Tokens)
	delSharesDec, err := decimal.NewFromString(validator.DelegatorShares)

	rate := valTokensDec.Div(delSharesDec)

	selfAsFloat64, _ := selfBondShares.Mul(rate).Mul(decimal.New(10, 17)).Float64()
	delegatedAsFloat64, _ := delSharesDec.Mul(rate).Mul(decimal.New(10, 17)).Float64()

	selfBond := document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: selfAsFloat64,
	}

	delegated := document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: delegatedAsFloat64,
	}

	return selfBond, delegated
}

func (service *DelegatorService) GetDeposits(delAddr string) (coin document.Coin) {

	var query = orm.NewQuery()
	defer query.Release()

	var delegations []document.Delegator
	query.SetCollection(document.CollectionNmStakeRoleDelegator).
		SetCondition(bson.M{document.Delegator_Field_Addres: delAddr}).
		SetResult(&delegations)

	if query.Exec() != nil {
		logger.Warn("delegator address not exist", logger.String("delAddr", delAddr))
		return
	}

	var delegationMap = make(map[string]document.Delegator, len(delegations))
	var valAddrs []string
	for _, d := range delegations {
		delegationMap[d.ValidatorAddr] = d
		valAddrs = append(valAddrs, d.ValidatorAddr)
	}

	var condition = bson.M{}
	condition[document.Candidate_Field_Address] = bson.M{
		"$in": valAddrs,
	}
	var validators []document.Candidate
	query.Reset().SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(condition).
		SetResult(&validators)

	if query.Exec() != nil {
		logger.Error("validator not found", logger.Any("valAddrs", valAddrs))
		return
	}

	var totalAmt float64
	for _, v := range validators {
		delegation := delegationMap[v.Address]
		if v.Tokens > 0 && v.DelegatorShares > 0 {
			rate := v.Tokens / v.DelegatorShares
			totalAmt += delegation.Shares * rate
		}
	}
	return document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: totalAmt,
	}
}

type ValInfo struct {
	model.ValProfile
	selfBond  document.Coin
	delegated document.Coin
}
