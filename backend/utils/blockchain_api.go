package utils

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/irishub-sync/store"
	"log"
)

type Account struct {
	Coins      []string `json:"coins"`
	AccountNum string   `json:"account_number"`
	Sequence   string   `json:"sequence"`
}

func GetBalance(address string) (amoumt store.Coins) {
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
			funBuildCoins := func(coins []string) []store.Coin {
				if len(coins) > 0 {
					for _, coinstr := range coins {
						coin := ParseCoin(coinstr)
						amoumt = append(amoumt, CovertCoin(coin, CoinTypeAtto))
					}
				}
				return amoumt
			}
			return funBuildCoins(account.Coins)
		}
		return amoumt
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
