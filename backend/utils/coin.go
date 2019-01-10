package utils

import (
	"fmt"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"regexp"
	"strconv"
	"strings"
)

const (
	CoinTypeIris  = "iris"
	CoinTypeAtto  = "iris-atto"
	CoinTypeFemto = "iris-femto"
	CoinTypePico  = "iris-pico"
	CoinTypeNano  = "iris-nano"
	CoinTypeMicro = "iris-micro"
	CoinTypeMilli = "iris-milli"
)

var coinsMap = make(map[string]float64)

func init() {
	coinsMap[CoinTypeIris] = float64(1)
	coinsMap[CoinTypeMilli] = float64(1000)
	coinsMap[CoinTypeMicro] = float64(1000000)
	coinsMap[CoinTypeNano] = float64(1000000000)
	coinsMap[CoinTypePico] = float64(1000000000000)
	coinsMap[CoinTypeFemto] = float64(1000000000000000)
	coinsMap[CoinTypeAtto] = float64(1000000000000000000)
}

func ParseCoin(coinStr string) (coin document.Coin) {
	var (
		reDnm  = `[A-Za-z\-]{2,15}`
		reAmt  = `[0-9]+[.]?[0-9]*`
		reSpc  = `[[:space:]]*`
		reCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
	)

	coinStr = strings.TrimSpace(coinStr)

	matches := reCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		logger.Error("invalid coin expression", logger.Any("coin", coinStr))
		return
	}
	denom, amount := matches[2], matches[1]

	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		logger.Error("Convert str to int failed", logger.Any("amount", amount))
	}

	return document.Coin{
		Denom:  denom,
		Amount: amt,
	}
}
func ParseCoins(coinsStr string) (coins document.Coins) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return
	}

	coinStrs := strings.Split(coinsStr, ",")
	for _, coinStr := range coinStrs {
		coin := ParseCoin(coinStr)
		coins = append(coins, coin)
	}

	return coins
}

func CovertCoin(srcCoin document.Coin, denom string) (destCoin document.Coin) {
	srcPreci := coinsMap[srcCoin.Denom]
	dstPreci := coinsMap[denom]

	dstAmt := srcCoin.Amount * (dstPreci / srcPreci)
	destCoin.Amount = dstAmt
	destCoin.Denom = denom
	return
}
