package lcd

import (
	"github.com/irisnet/explorer/backend/logger"
	"fmt"
	"encoding/json"
)

func GetAssetTokens() (result []AssetToken) {

	tokens, err := client.Asset().QueryTokens("")
	if err != nil {
		logger.Error("Query Asset Token error", logger.String("err", err.Error()))
		return
	}
	data, _ := json.Marshal(tokens)
	fmt.Println(string(data))
	return
}

func GetAssetByAddr(address string) (result []AssetToken) {
	tokens, err := client.Asset().QueryTokens(address)
	if err != nil {
		logger.Error("Query Asset Token by address error", logger.String("err", err.Error()))
		return
	}
	for _, val := range tokens {
		result = append(result, AssetToken{BaseToken: BaseToken{
			Symbol:        val.Symbol,
			Name:          val.Name,
			Scale:         int(val.Scale),
			MinUnitAlias:  val.MinUnit,
			InitialSupply: fmt.Sprint(val.InitialSupply),
			MaxSupply:     fmt.Sprint(val.MaxSupply),
			Mintable:      val.Mintable,
			Owner:         val.Owner,
		},
		})
	}
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
