package service

import (
	"fmt"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/module/logger"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type AccountService struct{}

func (service *AccountService) GetModule() Module {
	return Account
}

func (service *AccountService) Query(address string) (result model.AccountResp) {

	db := getDb()
	c := db.C(document.CollectionNmAccount)
	defer db.Session.Close()
	err := c.Find(bson.M{document.Account_Field_Addres: address}).One(&result)
	if err != nil {
		error := types.ErrorCodeNotFound
		logger.Error("account don't found", logger.String("address", address))
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
