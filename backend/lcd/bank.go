package lcd

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/utils"
	"fmt"
)

type TokenStats struct {
	LooseTokens  []*Coin `json:"loose_tokens"`
	BurnedTokens []*Coin `json:"burned_tokens"`
	BondedTokens []*Coin `json:"bonded_tokens"`
	TotalSupply  []*Coin `json:"total_supply"`
}

func GetBankTokenStats() (TokenStats, error) {

	var result TokenStats
	tokens, err := client.Bank().QueryTokenStats("")
	if err != nil {
		return result, err
	}
	data, _ := json.Marshal(tokens)
	fmt.Println(data)
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
	account, err := client.Bank().QueryAccount(CommunityTaxAddr)
	if err != nil {
		return Coin{}, err
	}
	data, _ := json.Marshal(account)
	fmt.Println(data)
	acc := Account01411{

	}

	return GetTokens(acc.Value.Coins), nil
}

//func GetTokenInitSupply() Coin {
//	return Coin{
//		Amount: conf.IniSupply,
//		Denom:  utils.CoinTypeIris,
//	}
//}
