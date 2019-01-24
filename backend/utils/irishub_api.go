package utils

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"log"
	"net/http"
)

type Account struct {
	Coins      []string `json:"coins"`
	AccountNum string   `json:"account_number"`
	Sequence   string   `json:"sequence"`
}

func GetBalance(address string) (amount document.Coins) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	uri := fmt.Sprintf(types.UrlIrisHubAccount, conf.Get().Server.HubLcdUrl, address)
	resBytes, err := Get(uri)
	if err == nil {
		var account Account
		if err := json.Unmarshal(resBytes, &account); err == nil {
			funBuildCoins := func(coins []string) []document.Coin {
				if len(coins) > 0 {
					for _, coinstr := range coins {
						coin := ParseCoin(coinstr)
						amount = append(amount, CovertCoin(coin, CoinTypeAtto))
					}
				}
				return amount
			}
			return funBuildCoins(account.Coins)
		}
		return amount
	}
	return
}

func GetNodes() (bz []byte) {
	uri := fmt.Sprintf(types.UrlIrisHubNetInfo, conf.Get().Server.HubNodeUrl)
	bz, err := Get(uri)
	if err != nil {
		panic(err)
	}
	return bz
}

func GetFaucetAccount(req *http.Request) (bz []byte, err error) {
	uri := fmt.Sprintf(types.UrlFaucetAccountService, conf.Get().Server.HubFaucetUrl)
	return Forward(req, uri)
}

func Apply(req *http.Request) (bz []byte, err error) {
	uri := fmt.Sprintf(types.UrlFaucetApplyService, conf.Get().Server.HubFaucetUrl)
	return Forward(req, uri)
}

func Genesis() model.Genesis {
	uri := fmt.Sprintf(types.UrlIrisHubGenesis, conf.Get().Server.HubNodeUrl)
	bz, err := Get(uri)
	if err != nil {
		panic(err)
	}
	var genesis model.Genesis
	if json.Unmarshal(bz, &genesis); err != nil {
		logger.Error("json Unmarshal genesis fail", logger.String("err", err.Error()))
	}
	return genesis
}
