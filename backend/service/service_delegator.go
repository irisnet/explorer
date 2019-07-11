package service

import (
	"math"
	"math/big"

	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
)

type DelegatorService struct {
	BaseService
}

func (service *DelegatorService) GetModule() Module {
	return Delegator
}

func (service *DelegatorService) QueryDelegation(valAddr string) (utils.Coin, utils.Coin) {
	validator, err := lcd.Validator(valAddr)
	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return utils.Coin{}, utils.Coin{}
	}
	rate, err := utils.QuoByStr(validator.Tokens, validator.DelegatorShares)
	if err != nil {
		logger.Error("validator.Tokens / validator.DelegatorShares", logger.String("err", err.Error()))
		return utils.Coin{}, utils.Coin{}
	}

	tokens, selfBond, err := document.Delegator{}.GetValidatorTokenAndSelfBond(valAddr)

	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return utils.Coin{}, utils.Coin{}
	}

	logger.Info("query delegation by validator addres", logger.String("validatorAddr", valAddr), logger.String("tokens", tokens), logger.String("self_bond", selfBond))

	tokensAsRat, ok := new(big.Rat).SetString(tokens)
	if !ok {
		logger.Error("convert validator tokens type (string -> big.Rat) err", logger.Any("err", err.Error()), logger.String("token str", tokens))
		return utils.Coin{}, utils.Coin{}
	}

	selfBondAsRat, ok := new(big.Rat).SetString(selfBond)
	if !ok {
		logger.Error("convert validator selfBond type (string -> big.Rat) err", logger.Any("err", err.Error()), logger.String("self bond str", selfBond))
		return utils.Coin{}, utils.Coin{}

	}

	selfBondAsFloat64, exact := new(big.Rat).Mul(selfBondAsRat, new(big.Rat).SetFloat64(math.Pow10(18))).Float64()
	if !exact {
		logger.Info("convert selfBondAsRat type (big.Rat to float64) ",
			logger.Any("exact", exact),
			logger.Any("selfBondAsRat", selfBondAsRat))
	}

	selfBondTokensAsRat := new(big.Rat).Mul(selfBondAsRat, rate)
	BondStakeAsRat := new(big.Rat).Sub(tokensAsRat, selfBondTokensAsRat)
	BondStakeAsFloat64, exact := new(big.Rat).Mul(BondStakeAsRat, new(big.Rat).SetFloat64(math.Pow10(18))).Float64()
	if !exact {
		logger.Info("convert otherBondAsRat type (big.Rat to float64) ",
			logger.Any("exact", exact),
			logger.Any("BondStakeAsRat", BondStakeAsRat))
	}

	return utils.Coin{
			Denom:  utils.CoinTypeAtto,
			Amount: selfBondAsFloat64,
		}, utils.Coin{
			Denom:  utils.CoinTypeAtto,
			Amount: BondStakeAsFloat64,
		}
}

func (service *DelegatorService) GetDeposits(addressAsAccount string) utils.Coin {
	delegations := lcd.GetDelegationsByDelAddr(addressAsAccount)
	delegationMap := make(map[string]lcd.DelegationVo, len(delegations))
	valAddrs := []string{}
	for _, d := range delegations {
		delegationMap[d.ValidatorAddr] = d
		valAddrs = append(valAddrs, d.ValidatorAddr)
	}

	validators, err := document.Delegator{}.GetDepositValidator(valAddrs)

	if err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return utils.Coin{}
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

	return utils.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: totalAmtAsFloat64,
	}
}
