package lcd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"errors"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"strings"
)

type (
	Coin struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	}

	Value struct {
		Coins      []*Coin   `json:"coins"`
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


func Account(address string) (result AccountVo, err error) {
	acc, err := AccountInfo(address)
	if err != nil {
		return result, err
	}
	result.Address = acc.Address
	result.Sequence = acc.Sequence
	result.AccountNumber = acc.AccountNumber
	result.PublicKey = acc.PublicKey
	ret, err := AccountBalances(address)
	if err != nil {
		return result, err
	}
	coinsStrArr := []string{}
	for _, v := range ret {
		coinsStrArr = append(coinsStrArr, v.Amount+v.Denom)
	}
	result.Address = address
	result.Coins = coinsStrArr
	return result, nil
}

func AccountInfo(address string) (AccountVo, error) {
	acc := AccountVo{}
	if !strings.HasPrefix(address, conf.Get().Hub.Prefix.AccAddr) {
		return acc, fmt.Errorf("address prefix is should %v", conf.Get().Hub.Prefix.AccAddr)
	}
	account, err := client.Bank().QueryAccount(address)
	if err != nil {
		return acc, err
	}
	acc = AccountVo{
		Address:       account.Address.String(),
		AccountNumber: account.AccountNumber,
		PublicKey:     string(account.PubKey),
		Sequence:      fmt.Sprint(account.Sequence),
	}

	return acc, nil
}

func AccountBalances(address string) ([]Coin, error) {
	balances, err := client.Bank().QueryBalances(address, "")
	if err != nil {
		return nil, err
	}
	var res []Coin
	for _, val := range balances {
		res = append(res, Coin{
			Denom:  val.Denom,
			Amount: val.Amount,
		})
	}
	return res, nil
}
func Faucet(req *http.Request) (bz []byte, err error) {
	uri := fmt.Sprintf(types.UrlFaucetAccountService, conf.Get().Server.FaucetUrl)
	return utils.Forward(req, uri)
}

func GetToken(req *http.Request) (bz []byte, err error) {
	uri := fmt.Sprintf(types.UrlFaucetApplyService, conf.Get().Server.FaucetUrl)
	return utils.Forward(req, uri)
}

func GetIconsByKey(key string) (string, error) {
	url := fmt.Sprintf(UrlLookupIconsByKeySuffix, key)
	resBytes, err := utils.Get(url)
	if err != nil {
		return "", err
	}
	var picdata vo.LookupIcons
	if err := json.Unmarshal(resBytes, &picdata); err != nil {
		//logger.Error("get icons error", logger.String("err", err.Error()))
		return "", err
	}

	if picdata.Status.Code != 0 {
		return "", errors.New("get icons failed")
	}

	if len(picdata.Them) == 0 {
		return "", nil
	}

	return picdata.Them[0].Pictures.Primary.Url, nil
}
