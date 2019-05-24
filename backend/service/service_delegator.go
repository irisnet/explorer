package service

import (
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

	validator := document.Validator{}
	selector := bson.M{
		"tokens":    1,
		"self_bond": 1}
	err := queryOne(document.CollectionNmValidator, selector, bson.M{document.ValidatorFieldOperatorAddress: valAddr}, &validator)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return document.Coin{}, document.Coin{}
	}

	logger.Info("query delegation by validator addres", logger.String("validatorAddr", valAddr), logger.String("tokens", validator.Tokens), logger.String("self_bond", validator.SelfBond))

	tokensAsRat, ok := new(big.Rat).SetString(validator.Tokens)
	if !ok {
		logger.Error("convert validator tokens type (string -> big.Rat) err", logger.Any("err", err.Error()), logger.String("token str", validator.Tokens))
		return document.Coin{}, document.Coin{}
	}

	selfBondAsRat, ok := new(big.Rat).SetString(validator.SelfBond)
	if !ok {
		logger.Error("convert validator selfBond type (string -> big.Rat) err", logger.Any("err", err.Error()), logger.String("self bond str", validator.SelfBond))
		return document.Coin{}, document.Coin{}
	}

	selfBondAsFloat64, exact := new(big.Rat).Mul(selfBondAsRat, new(big.Rat).SetFloat64(math.Pow10(18))).Float64()
	if !exact {
		logger.Info("convert selfBondAsRat type (big.Rat to float64) ",
			logger.Any("exact", exact),
			logger.Any("selfBondAsRat", selfBondAsRat))
	}

	otherBondAsRat := new(big.Rat)
	otherBondAsFloat64, exact := otherBondAsRat.Sub(tokensAsRat, selfBondAsRat).Mul(otherBondAsRat, new(big.Rat).SetFloat64(math.Pow10(18))).Float64()
	if !exact {
		logger.Info("convert otherBondAsRat type (big.Rat to float64) ",
			logger.Any("exact", exact),
			logger.Any("otherBondAsRat", otherBondAsRat))
	}

	return document.Coin{
			Denom:  utils.CoinTypeAtto,
			Amount: selfBondAsFloat64,
		}, document.Coin{
			Denom:  utils.CoinTypeAtto,
			Amount: otherBondAsFloat64,
		}
}

func (service *DelegatorService) GetDeposits(addressAsAccount string) document.Coin {
	delegations := lcd.GetDelegationsByDelAddr(addressAsAccount)
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

		rate, err := utils.QuoByStr(v.Tokens, v.DelegatorShares)
		if err != nil {
			logger.Error("validator.Tokens / validator.DelegatorShares", logger.String("err", err.Error()))
			continue
		}

		delegationSharesAsRat, ok := new(big.Rat).SetString(delegation.Shares)
		if !ok {
			logger.Info("convert Delegation.Shares type (string to big.Rat) ", logger.Any("result", ok), logger.String("delegation share", delegation.Shares))
			continue
		}

		totalAmtAsRat.Add(totalAmtAsRat, new(big.Rat).Mul(delegationSharesAsRat, rate))
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
