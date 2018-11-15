package service

import (
	"fmt"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type AccountService struct {
	*BaseService
}

func GetAccount() *AccountService {
	return accountService
}

func (service *AccountService) Query(address string) (result document.Account) {

	c := service.GetDb().C(document.CollectionNmAccount)
	defer c.Database.Session.Close()
	err := c.Find(bson.M{"address": address}).Sort("-amount.amount").One(&result)
	if err != nil {
		log.Println(fmt.Sprintf("account [%s] not found", address))
		return
	}

	balance := utils.GetBalance(result.Address)
	if len(balance) >= 0 {
		result.Amount = balance
	}
	return
}

func (service *AccountService) QueryList(page, size int) model.Page {
	var result []document.Account
	return service.QueryPage(document.CollectionNmAccount, &result, nil, "-time", page, size)
}
