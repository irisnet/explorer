package lcd

import (
	"github.com/irisnet/explorer/backend/logger"
	"fmt"
)

func GetAssetTokens() (result []AssetTokens) {
	//url := fmt.Sprintf(UrlAssetTokens, conf.Get().Hub.LcdUrl)
	//resBytes, err := utils.Get(url)
	//if err != nil {
	//	logger.Error("get AssetTokens error", logger.String("err", err.Error()))
	//	return result
	//}
	//
	//if err := json.Unmarshal(resBytes, &result); err != nil {
	//	logger.Error("Unmarshal AssetTokens error", logger.String("err", err.Error()))
	//	return result
	//}
	tokens, err := client.Asset().QueryTokens("")
	if err != nil {
		logger.Error("Query AssetTokens error", logger.String("err", err.Error()))
		return
	}
	fmt.Println(tokens)
	return
}

//func GetAssetGateways() (result []AssetGateways) {
//	url := fmt.Sprintf(UrlAssetGateways, conf.Get().Hub.LcdUrl)
//	resBytes, err := utils.Get(url)
//	if err != nil {
//		logger.Error("get GetAssetGateways error", logger.String("err", err.Error()))
//		return result
//	}
//
//	if err := json.Unmarshal(resBytes, &result); err != nil {
//		logger.Error("Unmarshal GetAssetGateways error", logger.String("err", err.Error()))
//		return result
//	}
//	return result
//}
