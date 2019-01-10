package service

import (
	"github.com/irisnet/explorer/backend/logger"
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

func (service *AccountService) Query(address string) (result model.AccountResp) {
	db := getDb()
	c := db.C(document.CollectionNmAccount)
	defer db.Session.Close()
	err := c.Find(bson.M{document.Account_Field_Addres: address}).One(&result)
	if err != nil {
		error := types.CodeNotFound
		logger.Error("account don't found", service.GetTraceLog(), logger.String("address", address))
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
	query[document.Tx_Field_From] = address
	query[document.Tx_Field_Type] = types.TxTypeSetWithdrawAddress
	var tx document.CommonTx
	sort := desc(document.Tx_Field_Time)
	if err := txStore.Find(query).Sort(sort).One(&tx); err == nil {
		result.WithdrawAddress = tx.To
	}

	return
}

func (service *AccountService) QueryList(page, size int) model.Page {
	var result []document.Account
	sort := desc(document.Tx_Field_Time)
	return queryPage(document.CollectionNmAccount, &result, nil, sort, page, size)
}
