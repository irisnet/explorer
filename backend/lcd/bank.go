package lcd

import (
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/logger"
	"encoding/json"
)

type TokenStats struct {
	LooseTokens  []*Coin `json:"loose_tokens"`
	BurnedTokens []*Coin `json:"burned_tokens"`
	BondedTokens []*Coin `json:"bonded_tokens"`
}

func GetBankTokenStats() (TokenStats, error) {

	var result TokenStats
	url := fmt.Sprintf(UrlBankTokenStats, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("GetBankTokenStats have error", logger.String("err", err.Error()))
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("GetBankTokenStats Unmarshal error", logger.String("err", err.Error()))
		return result, err
	}
	return result, nil
}

func GetTokens(data []*Coin) Coin {

	for _, val := range data {
		if val.Denom == "iris-atto" {
			return Coin{Denom: val.Denom, Amount: val.Amount}
		}
	}
	return Coin{}
}

func GetTokenStatsCirculation() (Coin, error) {
	resBytes, err := utils.Get(UrlTokenStatsCirculation)
	if err != nil {
		logger.Error("GetTokenStatsCirculation have error", logger.String("err", err.Error()))
		return Coin{}, err
	}
	return Coin{
		Amount: string(resBytes),
		Denom:  "IRIS",
	}, nil
}

func GetTokenStatsSupply() (Coin, error) {
	resBytes, err := utils.Get(UrlTokenStatsSupply)
	if err != nil {
		logger.Error("GetTokenStatsSupply Unmarshal error", logger.String("err", err.Error()))
		return Coin{}, err
	}
	return Coin{
		Amount: string(resBytes),
		Denom:  "IRIS",
	}, nil
}

func GetTokenInitSupply() Coin {
	return Coin{
		Amount: conf.IniSupply,
		Denom:  "IRIS",
	}
}
