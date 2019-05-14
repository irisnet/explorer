package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
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
		result.Amount = document.Coins{self}
		result.Deposits = delegated

	} else {
		res, err := lcd.Account(address)
		if err == nil {
			var amount document.Coins
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

func (service *AccountService) QueryAll(page, size int) model.PageVo {
	var result []document.Account
	sort := desc(document.Tx_Field_Time)
	cnt, _ := pageQuery(document.CollectionNmAccount, nil, nil, sort, page, size, &result)
	return model.PageVo{
		Count: cnt,
		Data:  result,
	}
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
