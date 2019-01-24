package lcd

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

func Validator(address string) (result ValidatorVo, err error) {
	url := fmt.Sprintf(validatorUrl, conf.Get().Server.HubLcdUrl, address)
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
