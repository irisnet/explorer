package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/irisnet/explorer/backend/logger"
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

var (
	coinsMap = make(map[string]float64)

	reDnm  = `[A-Za-z\-]{2,15}`
	reAmt  = `[0-9]+[.]?[0-9]*`
	reSpc  = `[[:space:]]*`
	reCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
)

func init() {
	coinsMap[CoinTypeIris] = float64(1)
	coinsMap[CoinTypeMilli] = float64(1000)
	coinsMap[CoinTypeMicro] = float64(1000000)
	coinsMap[CoinTypeNano] = float64(1000000000)
	coinsMap[CoinTypePico] = float64(1000000000000)
	coinsMap[CoinTypeFemto] = float64(1000000000000000)
	coinsMap[CoinTypeAtto] = float64(1000000000000000000)
}

func ParseCoin(coinStr string) (coin Coin) {
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

	return Coin{
		Denom:  denom,
		Amount: amt,
	}
}
func ParseCoins(coinsStr string) (coins Coins) {
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

func Parse(coinsStr string) (string, string) {
	coinsStr = strings.TrimSpace(coinsStr)

	matches := reCoin.FindStringSubmatch(coinsStr)
	if matches == nil {
		logger.Error("invalid coin expression", logger.Any("coin", coinsStr))
		return "", ""
	}
	return matches[2], matches[1]
}

func CovertCoin(srcCoin Coin, denom string) (destCoin Coin) {
	srcPreci := coinsMap[srcCoin.Denom]
	dstPreci := coinsMap[denom]

	dstAmt := srcCoin.Amount * (dstPreci / srcPreci)
	destCoin.Amount = dstAmt
	destCoin.Denom = denom
	return
}

type Coin struct {
	Denom  string  `json:"denom"`
	Amount float64 `json:"amount"`
}

func (c Coin) String() string {
	return fmt.Sprintf(
		`Denom: %v Amount: %v`, c.Denom, c.Amount)
}

func (c Coin) Add(a Coin) Coin {
	if c.Denom == a.Denom {
		return Coin{
			Denom:  c.Denom,
			Amount: c.Amount + a.Amount,
		}
	}
	return c
}

type Coins []Coin

type Fee struct {
	Amount Coins `json:"amount"`
	Gas    int64 `json:"gas"`
}

type ActualFee struct {
	Denom  string  `json:"denom"`
	Amount float64 `json:"amount"`
}
