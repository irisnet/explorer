package service

import (
	"fmt"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type AccountService struct {
	*BaseService
}

func GetAccount() *AccountService {
	return &AccountService{
		baseService,
	}
}

func (service *AccountService) Query(address string) (result model.AccountResp) {

	db := service.GetDb()
	c := db.C(document.CollectionNmAccount)
	defer db.Session.Close()
	err := c.Find(bson.M{"address": address}).One(&result)
	if err != nil {
		error := types.ErrorCodeNotFound
		log.Println(fmt.Sprintf("account [%s] not found", address))
		panic(error)
		return
	}

	balance := utils.GetBalance(result.Address)
	if len(balance) > 0 {
		result.Amount = balance
	}
	result.WithdrawAddress = address

	txStore := db.C(document.CollectionNmCommonTx)
	query := bson.M{}
	query["from"] = address
	query["type"] = types.TxTypeSetWithdrawAddress
	var tx document.CommonTx
	if err := txStore.Find(query).Sort("-time").One(&tx); err == nil {
		result.WithdrawAddress = tx.To
	}

	return
}

func (service *AccountService) QueryList(page, size int) model.Page {
	var result []document.Account
	return service.QueryPage(document.CollectionNmAccount, &result, nil, "-time", page, size)
}
