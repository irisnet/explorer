package utils

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/irishub-sync/store"
	"io/ioutil"
	"log"
	"net/http"
)

var AddrNodeServer string
var AddrHubRpc string

func init() {
	AddrNodeServer = GetEnv("ADDR_NODE_SERVER", "http://127.0.0.1:1317")
	AddrHubRpc = GetEnv("ADDR_HUB_RPC", "http://192.168.150.7:26657")
}
func get(uri string) (int, []byte) {
	res, err := http.Get(uri)
	defer res.Body.Close()

	if err != nil {
		log.Println(err)
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	return res.StatusCode, resByte
}

type Account struct {
	Coins      []string `json:"coins"`
	AccountNum string   `json:"account_number"`
	Sequence   string   `json:"sequence"`
}

func GetBalance(address string) store.Coins {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	uri := fmt.Sprintf("%s/auth/accounts/%s", AddrNodeServer, address)
	statusCode, resBytes := get(uri)

	var account Account
	var resCoins []store.Coin
	if statusCode == 200 {
		if err := json.Unmarshal(resBytes, &account); err == nil {
			funBuildCoins := func(coins []string) []store.Coin {
				if len(coins) > 0 {
					for _, coinstr := range coins {
						coin := ParseCoin(coinstr)
						resCoins = append(resCoins, CovertCoin(coin, CoinTypeAtto))
					}
				}
				return resCoins
			}
			return funBuildCoins(account.Coins)
		}
	}
	return resCoins
}

func GetNodes() []byte {
	uri := fmt.Sprintf("%s/net_info", AddrHubRpc)
	statusCode, resBytes := get(uri)
	if statusCode != 200 {
		log.Println("query error,statusCode:", statusCode)
	}
	return resBytes
}
