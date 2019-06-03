package lcd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
)

type (
	Coin struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	}

	Value struct {
		Coins      []Coin    `json:"coins"`
		Address    string    `json:"address"`
		PublicKey  PublicKey `json:"public_key"`
		AccountNum string    `json:"account_number"`
		Sequence   string    `json:"sequence"`
	}

	PublicKey struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}
	Account01411 struct {
		Type  string `json:"type"`
		Value Value  `json:"value"`
	}
)

func buildAccountVo(acc Account01411) AccountVo {

	res := AccountVo{}

	res.Address = acc.Value.Address
	res.Sequence = acc.Value.Sequence
	res.AccountNumber = acc.Value.AccountNum
	res.PublicKey.Type = acc.Value.PublicKey.Type
	res.PublicKey.Value = acc.Value.PublicKey.Value

	coinsStrArr := []string{}
	for _, v := range acc.Value.Coins {
		coinsStrArr = append(coinsStrArr, v.Amount+v.Denom)
	}
	res.Coins = coinsStrArr
	return res
}

func Account(address string) (result AccountVo, err error) {
	url := fmt.Sprintf(UrlAccount, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}
	acc := Account01411{}
	if err := json.Unmarshal(resBytes, &acc); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	result = buildAccountVo(acc)
	return result, nil
}

func Faucet(req *http.Request) (bz []byte, err error) {
	uri := fmt.Sprintf(types.UrlFaucetAccountService, conf.Get().Server.FaucetUrl)
	return utils.Forward(req, uri)
}

func GetToken(req *http.Request) (bz []byte, err error) {
	uri := fmt.Sprintf(types.UrlFaucetApplyService, conf.Get().Server.FaucetUrl)
	return utils.Forward(req, uri)
}
