package lcd

import (
	"github.com/irisnet/explorer/backend/types"
)

//type TokenStats struct {
//	LooseTokens  []*Coin `json:"loose_tokens"`
//	BurnedTokens []*Coin `json:"burned_tokens"`
//	BondedTokens []*Coin `json:"bonded_tokens"`
//	TotalSupply  []*Coin `json:"total_supply"`
//}

//func GetBankTokenStats() (TokenStats, error) {
//
//	var result TokenStats
//	tokens, err := client.Bank().QueryTokenStats("")
//	if err != nil {
//		return result, err
//	}
//	data, _ := json.Marshal(tokens)
//	fmt.Println(data)
//	return result, nil
//}

//func GetTokens(data []*Coin) Coin {
//
//	for _, val := range data {
//		if val.Denom == utils.CoinTypeStake {
//			return Coin{Denom: val.Denom, Amount: val.Amount}
//		}
//	}
//	return Coin{}
//}

//func GetTokenStatsCirculation() (Coin, error) {
//	resBytes, err := utils.Get(UrlTokenStatsCirculation)
//	if err != nil {
//		return Coin{}, err
//	}
//	return Coin{
//		Amount: string(resBytes),
//		Denom:  utils.CoinTypeStake,
//	}, nil
//}
//
//func GetTokenStatsSupply() (Coin, error) {
//	resBytes, err := utils.Get(UrlTokenStatsSupply)
//	if err != nil {
//		return Coin{}, err
//	}
//	return Coin{
//		Amount: string(resBytes),
//		Denom:  utils.CoinTypeStake,
//	}, nil
//}
func GetCommunityTax() (Coin, error) {
	balances, err := client.Bank().QueryBalances(CommunityTaxAddr, "")
	if err != nil {
		return Coin{}, err
	}

	var res Coin
	for _, val := range balances {
		if val.Denom != types.StakeUint {
			continue
		}
		res = Coin{
			Denom:  val.Denom,
			Amount: val.Amount,
		}
	}

	return res, nil
}

func GetTotalSupply() (Coin, error) {
	coins, err := client.Bank().QueryTotalSupply()
	if err != nil {
		return Coin{}, err
	}
	var res Coin
	for _, val := range coins {
		if val.Denom != types.StakeUint {
			continue
		}
		res = Coin{
			Denom:  val.Denom,
			Amount: val.Amount.String(),
		}
	}
	return res, nil
}

//func GetTokenInitSupply() Coin {
//	return Coin{
//		Amount: conf.IniSupply,
//		Denom:  utils.CoinTypeIris,
//	}
//}
