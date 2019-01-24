package lcd

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

func Account(address string) (result AccountVo, err error) {
	resBytes, err := utils.Get(getUrl(accountUrl, address))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	return result, nil
}
