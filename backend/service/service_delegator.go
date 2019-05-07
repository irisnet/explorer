package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"math"
	"math/big"
)

type DelegatorService struct {
	BaseService
}

func (service *DelegatorService) GetModule() Module {
	return Delegator
}

func (service *DelegatorService) QueryDelegation(valAddr string) (document.Coin, document.Coin) {

	accAddr := utils.Convert(conf.Get().Hub.Prefix.AccAddr, valAddr)

	validator := document.Validator{}
	selector := bson.M{
		"tokens":           1,
		"delegator_shares": 1}
	err := queryOne(document.CollectionNmValidator, selector, bson.M{document.ValidatorFieldOperatorAddress: valAddr}, &validator)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return document.Coin{}, document.Coin{}
	}

	valTokenAsRat, ok := new(big.Rat).SetString(validator.Tokens)
	if !ok {
		logger.Error("convert validator.tokens type (string to big.Rat) err", logger.Any("err", err.Error()), logger.String("validator.Tokens str", validator.Tokens))
		return document.Coin{}, document.Coin{}
	}
	valDelegatorSharesAsRat, ok := new(big.Rat).SetString(validator.DelegatorShares)
	if !ok {
		logger.Error("convert validator.DelegatorShares type (string to big.Rat) err", logger.Any("err", err.Error()), logger.String("delegation shares str", validator.DelegatorShares))
		return document.Coin{}, document.Coin{}
	}

	rate := new(big.Rat).Quo(valTokenAsRat, valDelegatorSharesAsRat)

	delegations := lcd.DelegationByValidator(valAddr)

	selfBondShares := new(big.Rat)
	otherBondShares := new(big.Rat)
	for _, d := range delegations {
		tmp, ok := new(big.Rat).SetString(d.Shares)
		if !ok {
			logger.Info("convert delegation->Shares type (string to big.Rat) ", logger.Any("ok", ok), logger.String("delegation shares str", d.Shares))
			continue
		}
		if d.DelegatorAddr == accAddr {
			selfBondShares.Add(selfBondShares, tmp)
		} else {
			otherBondShares.Add(otherBondShares, tmp)
		}
	}

	selfShareAsFloat64, exact := selfBondShares.Mul(selfBondShares, rate).Mul(selfBondShares, new(big.Rat).SetFloat64(math.Pow10(18))).Float64()
	if !exact {
		logger.Info("convert selfBondShares type (big.Rat to float64)", logger.Any("exact", exact))
	}
	otherShareAsFloat64, exact := otherBondShares.Mul(otherBondShares, rate).Mul(otherBondShares, new(big.Rat).SetFloat64(math.Pow10(18))).Float64()
	if !exact {
		logger.Info("convert otherBondShares type (big.Rat to float64)", logger.Any("exact", exact))
	}

	return document.Coin{
			Denom:  utils.CoinTypeAtto,
			Amount: selfShareAsFloat64,
		}, document.Coin{
			Denom:  utils.CoinTypeAtto,
			Amount: otherShareAsFloat64,
		}
}

func (service *DelegatorService) GetDeposits(addressAsAccount string) document.Coin {
	delegations := lcd.GetAllDelegationsByDelegatorAddr(addressAsAccount)
	delegationMap := make(map[string]lcd.DelegationVo, len(delegations))
	valAddrs := []string{}
	for _, d := range delegations {
		delegationMap[d.ValidatorAddr] = d
		valAddrs = append(valAddrs, d.ValidatorAddr)
	}

	validators := []document.Validator{}
	selector := bson.M{
		document.ValidatorFieldOperatorAddress: 1,
		"tokens":                               1,
		"delegator_shares":                     1}
	condition := bson.M{
		document.ValidatorFieldOperatorAddress: bson.M{"$in": valAddrs},
	}
	err := queryAll(document.CollectionNmValidator, selector, condition, "", 0, &validators)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return document.Coin{}
	}

	totalAmtAsRat := new(big.Rat)
	for _, v := range validators {
		delegation := delegationMap[v.OperatorAddress]

		tokenAsRat, ok := new(big.Rat).SetString(v.Tokens)
		if !ok {
			logger.Info("convert validator->tokens type  (string to big.Rat)",
				logger.Any("ok", ok),
				logger.String("tokensStr", v.Tokens))
			continue
		}

		valDelSharesAsRat, ok := new(big.Rat).SetString(v.DelegatorShares)
		if !ok {
			logger.Info("convert validator->DelegatorShares type (string to big.Rat) ",
				logger.Any("ok", ok),
				logger.String("validator delegator shares", v.DelegatorShares))
			continue
		}

		delegationSharesAsRat, ok := new(big.Rat).SetString(delegation.Shares)
		if !ok {
			logger.Info("convert Delegation->Shares type (string to big.Rat) ",
				logger.String("err", err.Error()),
				logger.String("delegation share", delegation.Shares))
			continue
		}

		if tokenAsRat.Cmp(new(big.Rat).SetInt64(0)) == 1 && valDelSharesAsRat.Cmp(new(big.Rat).SetInt64(0)) == 1 {
			rate := new(big.Rat).Quo(tokenAsRat, valDelSharesAsRat)
			totalAmtAsRat.Add(totalAmtAsRat, new(big.Rat).Mul(delegationSharesAsRat, rate))
		}
	}
	totalAmtAsFloat64, exact := new(big.Rat).Mul(totalAmtAsRat, new(big.Rat).SetFloat64(math.Pow10(18))).Float64()
	if !exact {
		logger.Info("convert totalAmtAsFloat64 type (big.Rat to float64) ",
			logger.Any("exact", exact),
			logger.Any("totalAmtAsRat", totalAmtAsRat))
	}

	return document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: totalAmtAsFloat64,
	}
}
