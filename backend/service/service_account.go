package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type AccountService struct {
	BaseService
}

func (service *AccountService) GetModule() Module {
	return Account
}

func (service *AccountService) Query(address string) interface{} {
	prefix, _, _ := utils.DecodeAndConvert(address)
	if prefix == conf.Get().Hub.Prefix.ValAddr {
		var valAccInfo model.ValAccVo
		valInfo := delegatorService.QueryDelegation(address)
		valAccInfo.Amount = document.Coins{valInfo.selfBond}
		valAccInfo.Deposits = valInfo.delegated
		valAccInfo.ValProfile = valInfo.ValProfile
		valAccInfo.WithdrawAddress = queryWithdrawAddr(address)
		valAccInfo.IsProfiler = isProfiler(address)
		valAccInfo.Address = address
		return valAccInfo
	}

	var accInfo model.AccountVo
	res, err := lcd.Account(address)
	if err == nil {
		var amount document.Coins
		for _, coinStr := range res.Coins {
			coin := utils.ParseCoin(coinStr)
			amount = append(amount, coin)
		}
		accInfo.Amount = amount
	}
	accInfo.Deposits = delegatorService.GetDeposits(address)
	accInfo.WithdrawAddress = queryWithdrawAddr(address)
	accInfo.IsProfiler = isProfiler(address)
	accInfo.Address = address
	return accInfo
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

func queryWithdrawAddr(address string) (result string) {
	var tx document.CommonTx
	var selector = bson.M{"to": 1}
	var accAddr = utils.Convert(conf.Get().Hub.Prefix.AccAddr, address)

	condition := bson.M{}
	condition[document.Tx_Field_From] = accAddr
	condition[document.Tx_Field_Type] = types.TxTypeSetWithdrawAddress

	if err := queryOne(document.CollectionNmCommonTx, selector, condition, &tx); err == nil {
		result = tx.To
	} else {
		result = accAddr
	}
	return
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
