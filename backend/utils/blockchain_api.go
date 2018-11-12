package utils

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/irishub-sync/store"
	"log"
)

var AddrNodeServer string
var AddrHubRpc string

func init() {
	AddrNodeServer = GetEnv("ADDR_NODE_SERVER", "http://47.104.155.125:1317")
	AddrHubRpc = GetEnv("ADDR_HUB_RPC", "http://47.104.155.125:26657")
}

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

	uri := fmt.Sprintf("%s/bank/accounts/%s", AddrNodeServer, address)
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
	uri := fmt.Sprintf("%s/net_info", AddrHubRpc)
	bz, err := Get(uri)
	if err != nil {
		return
	}
	return bz
}
