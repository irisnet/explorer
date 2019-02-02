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

func (service *AccountService) Query(address string) (result model.AccountVo) {
	prefix, _, _ := utils.DecodeAndConvert(address)
	if prefix == conf.Get().Hub.Prefix.ValAddr {
		valinfo := delegatorService.QueryDelegation(address)
		result.Amount = document.Coins{valinfo.selfBond}
		result.Deposits = valinfo.delegated
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

	result.WithdrawAddress = queryWithdrawAddr(address)
	result.IsProfiler = isProfiler(address)
	result.Address = address
	return result
}

func (service *AccountService) QueryAll(page, size int) model.PageVo {
	var result []document.Account
	sort := desc(document.Tx_Field_Time)
	return queryPage(document.CollectionNmAccount, &result, nil, sort, page, size)
}

func queryWithdrawAddr(address string) (result string) {
	var accAddr = utils.Convert(conf.Get().Hub.Prefix.AccAddr, address)
	db := getDb()

	txStore := db.C(document.CollectionNmCommonTx)
	query := bson.M{}
	query[document.Tx_Field_From] = accAddr
	query[document.Tx_Field_Type] = types.TxTypeSetWithdrawAddress
	var tx document.CommonTx
	if err := txStore.Find(query).Sort(desc(document.Tx_Field_Time)).One(&tx); err == nil {
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
