package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
	"fmt"
)

type TokenStatsService struct {
	BaseService
}

func (service *TokenStatsService) QueryTokenStats() (model.TokenStatsVo, error) {

	var tokenStats model.TokenStatsVo

	banktokenstats, err := lcd.GetBankTokenStats()
	if err != nil {
		return tokenStats, err
	}
	supply, err := lcd.GetTokenStatsSupply()
	if err != nil {
		return tokenStats, err
	}

	circulation, err := lcd.GetTokenStatsCirculation()
	if err != nil {
		return tokenStats, err
	}

	initsupply := lcd.GetTokenInitSupply()
	burnedtokens := lcd.GetBuredTokens(banktokenstats.BurnedTokens)

	tokenStats.TotalsupplyTokens = LoadCoinVoFromLcdCoin(&supply)
	tokenStats.CirculationTokens = LoadCoinVoFromLcdCoin(&circulation)
	tokenStats.InitsupplyTokens = LoadCoinVoFromLcdCoin(&initsupply)
	tokenStats.DelegatedTokens = LoadCoinVoFromLcdCoin(banktokenstats.BondedTokens[0])
	tokenStats.BurnedTokens = LoadCoinVoFromLcdCoin(&burnedtokens)

	return tokenStats, nil
}

func LoadCoinVoFromLcdCoin(coin *lcd.Coin) *model.CoinVo {
	return &model.CoinVo{
		Denom:  coin.Denom,
		Amount: coin.Amount,
	}
}

func (service *TokenStatsService) QueryTokensAccountTotal() (map[string]model.TokenStatsSegment, error) {

	accounts, err := document.Account{}.GetAllAccount()
	if err != nil {
		logger.Error("QueryTokensAccountTotal have error", logger.String("err", err.Error()))
		return nil, err
	}

	var totalAmt = float64(0)

	for _, acc := range accounts {
		totalAmt += acc.Total.Amount
	}
	//fmt.Println("============>>:",len(accounts))

	if len(accounts) <= 5 {
		return computeSegment(accounts, totalAmt), nil
	}

	return computeSegment2(accounts,totalAmt), nil
}

func computeSegment(accounts []document.Account, totalAmt float64) map[string]model.TokenStatsSegment {
	accList := make(map[string]model.TokenStatsSegment)
	for index, acc := range accounts {
		rate, _ := utils.NewRatFromFloat64(acc.Total.Amount / totalAmt).Float64()
		accList[fmt.Sprint(index)] = model.TokenStatsSegment{
			Percent: rate,
			TotalAmount: &model.CoinVo{
				Denom:  acc.Total.Denom,
				Amount: utils.ParseStringFromFloat64(acc.Total.Amount),
			},
		}
	}
	return accList
}

func computeSegment2(accounts []document.Account, totalAmt float64) map[string]model.TokenStatsSegment {
	result := make(map[string]model.TokenStatsSegment)
	total := len(accounts)
	if total > 5 {
		result[fmt.Sprintf("%v-%v", 1, 5)] = CountNTotalAmount(0, 4, totalAmt, accounts)
		if total <= 10 {
			result[fmt.Sprintf("%v-%v", 6, total)] = CountNTotalAmount(5, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 6, 10)] = CountNTotalAmount(5, 9, totalAmt, accounts)
		if total <= 50 {
			result[fmt.Sprintf("%v-%v", 11, total)] = CountNTotalAmount(10, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 11, 50)] = CountNTotalAmount(10, 49, totalAmt, accounts)
		if total <= 100 {
			result[fmt.Sprintf("%v-%v", 51, total)] = CountNTotalAmount(50, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 51, 100)] = CountNTotalAmount(50, 99, totalAmt, accounts)
		if total <= 500 {
			result[fmt.Sprintf("%v-%v", 101, total)] = CountNTotalAmount(100, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 101, 500)] = CountNTotalAmount(100, 499, totalAmt, accounts)
		if total <= 1000 {
			result[fmt.Sprintf("%v-%v", 501, total)] = CountNTotalAmount(500, total-1, totalAmt, accounts)
			return result
		}

		result[fmt.Sprintf("%v-%v", 500, 1000)] = CountNTotalAmount(500, 999, totalAmt, accounts)
		result[fmt.Sprintf("%v-%v", 1001, total)] = CountNTotalAmount(1000, total-1, totalAmt, accounts)

	}
	return result
}

func CountNTotalAmount(start_index, end_index int, totalamt float64, account []document.Account) (result model.TokenStatsSegment) {

	var vaTotalAmt float64
	if totalamt <= 0 {
		return
	}

	retdata := &model.CoinVo{
		Denom: account[start_index].Total.Denom,
	}

	for pos := start_index; pos <= end_index; pos++ {
		vaTotalAmt += account[pos].Total.Amount
	}
	rate, _ := utils.NewRatFromFloat64(vaTotalAmt / totalamt).Float64()

	retdata.Amount = utils.ParseStringFromFloat64(vaTotalAmt)
	result.Percent = rate
	result.TotalAmount = retdata

	return
}
