package lcd

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

func GetAssetTokens() (result []AssetTokens) {
	url := fmt.Sprintf(UrlAssetTokens, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get AssetTokens error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal AssetTokens error", logger.String("err", err.Error()))
		return result
	}
	return result
}
