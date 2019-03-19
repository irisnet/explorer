package utils

import (
	"fmt"
	"github.com/irisnet/explorer/backend/logger"
	"math"
	"math/big"
	"strconv"
)

func ParseInt(text string) (i int64, b bool) {
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		return i, false
	}
	return i, true
}

func ParseBigInt(str string) *big.Int {
	f, ok := ParseInt(str)
	if !ok {
		logger.Error("ParseBigInt error", logger.String("str", str))
	}
	return big.NewInt(f)
}

func ParseBigFloat(str string, bit int) *big.Float {
	f, err := strconv.ParseFloat(str, bit)
	if err != nil {
		logger.Error("ParseBigFloat error", logger.String("str", str))
	}
	return big.NewFloat(f)
}

func ParseIntWithDefault(text string, def int64) (i int64) {
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		logger.Error("ParseIntWithDefault error", logger.String("str", text))
		return def
	}
	return i
}

func ParseUint(text string) (i int64, b bool) {
	i, ok := ParseInt(text)
	if ok {
		return i, i > 0
	}
	return i, ok
}

func RoundFloat(num float64, bit int) (i float64, b bool) {
	format := "%" + fmt.Sprintf("0.%d", bit) + "f"
	s := fmt.Sprintf(format, num)
	i, err := strconv.ParseFloat(s, 0)
	if err != nil {
		return i, false
	}
	return i, true
}

func Round(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

func RoundString(x string) int64 {
	f, err := strconv.ParseFloat(x, 0)
	if err != nil {
		logger.Error("RoundString error", logger.String("str", x))
	}
	return int64(math.Floor(f + 0.5))
}

func RoundToString(decimal string, bit int) (i string) {
	f, err := strconv.ParseFloat(decimal, bit)
	if err != nil {
		logger.Error("RoundFloatString error", logger.String("str", decimal))
	}
	return strconv.FormatFloat(f, 'f', bit, 64)
}
