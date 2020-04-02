package lcd

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/utils"
)

type TokenStats struct {
	LooseTokens  []*Coin `json:"loose_tokens"`
	BurnedTokens []*Coin `json:"burned_tokens"`
	BondedTokens []*Coin `json:"bonded_tokens"`
	TotalSupply  []*Coin `json:"total_supply"`
}

func GetBankTokenStats() (TokenStats, error) {

	var result TokenStats
	url := fmt.Sprintf(UrlBankTokenStats, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		return result, err
	}
	return result, nil
}

func GetTokens(data []*Coin) Coin {

	for _, val := range data {
		if val.Denom == utils.CoinTypeAtto {
			return Coin{Denom: val.Denom, Amount: val.Amount}
		}
	}
	return Coin{}
}

func GetTokenStatsCirculation() (Coin, error) {
	resBytes, err := utils.Get(UrlTokenStatsCirculation)
	if err != nil {
		return Coin{}, err
	}
	return Coin{
		Amount: string(resBytes),
		Denom:  utils.CoinTypeIris,
	}, nil
}

func GetTokenStatsSupply() (Coin, error) {
	resBytes, err := utils.Get(UrlTokenStatsSupply)
	if err != nil {
		return Coin{}, err
	}
	return Coin{
		Amount: string(resBytes),
		Denom:  utils.CoinTypeIris,
	}, nil
}
func GetCommunityTax() (Coin, error) {
	url := fmt.Sprintf(UrlAccount, conf.Get().Hub.LcdUrl, CommunityTaxAddr)
	resBytes, err := utils.Get(url)
	if err != nil {
		return Coin{}, err
	}
	acc := Account01411{}
	if err := json.Unmarshal(resBytes, &acc); err != nil {
		return Coin{}, err
	}

	return GetTokens(acc.Value.Coins), nil
}

//func GetTokenInitSupply() Coin {
//	return Coin{
//		Amount: conf.IniSupply,
//		Denom:  utils.CoinTypeIris,
//	}
//}
