package lcd

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"net/http"
)

func Account(address string) (result AccountVo, err error) {
	url := fmt.Sprintf(accountUrl, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
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
