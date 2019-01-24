package lcd

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

func NodeInfo() (result NodeInfoVo, err error) {
	resBytes, err := utils.Get(getUrl(nodeInfoUrl))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	return result, nil
}

func Genesis() (result GenesisVo, err error) {
	resBytes, err := utils.Get(fmt.Sprintf(genesisUrl, conf.Get().Server.HubNodeUrl))
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	return result, nil
}
