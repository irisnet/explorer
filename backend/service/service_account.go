package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/logger"
)

type AccountService struct {
	BaseService
}

func (service *AccountService) GetModule() Module {
	return Account
}

func (service *AccountService) Query(address string) (result model.AccountVo) {
	prefix, _, _ := utils.DecodeAndConvert(address)
	if prefix == conf.Get().Hub.Prefix.ValAddr {
		self, delegated := delegatorService.QueryDelegation(address)
		result.Amount = utils.Coins{self}
		result.Deposits = delegated

	} else {
		res, err := lcd.Account(address)
		if err == nil {
			var amount utils.Coins
			for _, coinStr := range res.Coins {
				coin := utils.ParseCoin(coinStr)
				amount = append(amount, coin)
			}
			result.Amount = amount
		}
		result.Deposits = delegatorService.GetDeposits(address)
	}

	result.WithdrawAddress = lcd.QueryWithdrawAddr(address)
	result.IsProfiler = isProfiler(address)
	result.Address = address
	return result
}

func (service *AccountService) QueryRichList() interface{} {

	result, err := document.Account{}.GetAccountList()

	if err != nil {
		logger.Error("GetAccountList have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	var accList []model.AccountInfo
	var totalAmt = float64(0)

	for _, acc := range result {
		totalAmt += acc.Total.Amount
	}

	for index, acc := range result {
		rate, _ := utils.NewRatFromFloat64(acc.Total.Amount / totalAmt).Float64()
		accList = append(accList, model.AccountInfo{
			Rank:    index + 1,
			Address: acc.Address,
			Balance: utils.Coins{
				acc.Total,
			},
			Percent:  rate,
			UpdateAt: acc.TotalUpdateAt,
		})
	}
	return accList
}

func isProfiler(address string) bool {
	genesis := commonService.GetGenesis()
	for _, profiler := range genesis.Result.Genesis.AppState.Guardian.Profilers {
		if profiler.Address == address {
			return true
		}
	}
	return false
}
