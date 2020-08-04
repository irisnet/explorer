package lcd

import (
	"github.com/irisnet/explorer/backend/logger"
	"fmt"
)

func GetAssetTokens() (result []AssetToken) {

	tokens, err := client.Asset().QueryTokens()
	if err != nil {
		logger.Error("Query Asset Token error", logger.String("err", err.Error()))
		return
	}

	baseToken, _ := QueryBaseDenom()

	for _, val := range tokens {
		if baseToken != "" && val.Value.Symbol == baseToken {
			continue
		}
		result = append(result, AssetToken{
			BaseToken: BaseToken{
				Name:          val.Value.Name,
				Scale:         int(val.Value.Scale),
				Symbol:        val.Value.Symbol,
				Mintable:      val.Value.Mintable,
				MaxSupply:     fmt.Sprint(val.Value.MaxSupply),
				Owner:         val.Value.Owner,
				InitialSupply: fmt.Sprint(val.Value.InitialSupply),
				MinUnitAlias:  val.Value.MinUnit,
			},
		})
	}

	return
}

//func GetAssetByAddr(address string) (result []AssetToken) {
//	tokens, err := client.Asset().QueryTokens(address)
//	if err != nil {
//		logger.Error("Query Asset Token by address error", logger.String("err", err.Error()))
//		return
//	}
//	for _, val := range tokens {
//		result = append(result, AssetToken{BaseToken: BaseToken{
//			Symbol:        val.Symbol,
//			Scale:         int(val.Scale),
//			InitialSupply: fmt.Sprint(val.InitialSupply),
//			MaxSupply:     fmt.Sprint(val.MaxSupply),
//			Mintable:      val.Mintable,
//			Owner:         val.Owner,
//		},
//		})
//	}
//	return
//}

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
