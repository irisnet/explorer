package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
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
			logger.Error("convert delegations-Shares type (string to decimal) err", logger.Any("err", err.Error()), logger.String("delegation shares str", d.Shares))
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

	selfAsFloat64, _ := selfBondShares.Mul(rate).Mul(decimal.New(1, 18)).Float64()
	delegatedAsFloat64, _ := delSharesDec.Mul(rate).Mul(decimal.New(1, 18)).Float64()

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

func (service *DelegatorService) GetDeposits(address string) document.Coin {
	delegations := lcd.GetAllDelegationsByDelegatorAddr(address)
	var delegationMap = make(map[string]lcd.DelegationVo, len(delegations))
	var valAddrs []string
	for _, d := range delegations {
		delegationMap[d.ValidatorAddr] = d
		valAddrs = append(valAddrs, d.ValidatorAddr)
	}

	var validators []document.Validator
	var selector = bson.M{
		document.ValidatorFieldOperatorAddress: 1,
		"tokens":                               1,
		"delegator_shares":                     1}

	var condition = bson.M{}
	condition[document.ValidatorFieldOperatorAddress] = bson.M{
		"$in": valAddrs,
	}

	err := queryAll(document.CollectionNmValidator, selector, condition, "", 0, &validators)

	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return document.Coin{}
	}

	var totalAmtAsDecimal decimal.Decimal

	for _, v := range validators {
		delegation := delegationMap[v.OperatorAddress]

		tokenAsDecimal, err := decimal.NewFromString(v.Tokens)
		if err != nil {
			logger.Error("convert validator->tokens type  (string to decimal) err", logger.String("err", err.Error()), logger.String("tokensStr", v.Tokens))
			continue
		}

		valDelSharesAsDecimal, err := decimal.NewFromString(v.DelegatorShares)
		if err != nil {
			logger.Error("convert validator->DelegatorShares type (string to decimal) err", logger.String("err", err.Error()), logger.String("validator delegator shares", v.DelegatorShares))
			continue
		}

		delegationSharesAsDecimal, err := decimal.NewFromString(delegation.Shares)
		if err != nil {
			logger.Error("convert Delegation->Shares type (string to decimal) err", logger.String("err", err.Error()), logger.String("delegation share", delegation.Shares))
			continue
		}

		if tokenAsDecimal.GreaterThan(decimal.New(0, 0)) && valDelSharesAsDecimal.GreaterThan(decimal.New(0, 0)) {
			rate := tokenAsDecimal.Div(valDelSharesAsDecimal)
			totalAmtAsDecimal = totalAmtAsDecimal.Add(delegationSharesAsDecimal.Mul(rate))
		}
	}
	totalAmtAsFloat64, _ := totalAmtAsDecimal.Mul(decimal.New(1, 18)).Float64()
	return document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: totalAmtAsFloat64,
	}
}
