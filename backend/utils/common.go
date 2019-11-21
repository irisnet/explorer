package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"crypto/md5"
	"encoding/hex"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/conf"
)

const (
	StatusFailed = "failed"
	StatusFail   = "fail"
)

func RemoveDuplicationStrArr(list []string) []string {
	unique_set := make(map[string]bool, len(list))
	for _, x := range list {
		unique_set[x] = true
	}
	result := make([]string, 0, len(unique_set))
	for x := range unique_set {
		result = append(result, x)
	}
	return result
}

func ParseInt(text string) (i int64, b bool) {
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		return i, false
	}
	return i, true
}

func ParseIntWithDefault(text string, def int64) (i int64) {
	if text == "" {
		return def
	}
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		logger.Error("ParseIntWithDefault error", logger.String("str", text))
		return def
	}
	return i
}

func ParseUint(text string) (uint64, bool) {
	i, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return i, false
	}
	return i, true
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

func RoundString(x string) (int64, error) {
	f, err := strconv.ParseFloat(x, 0)
	if err != nil {
		logger.Error("RoundString error", logger.String("str", x))
	}
	return int64(math.Floor(f + 0.5)), err
}

func ParseStringToFloat(x string) (float64, error) {
	f, err := strconv.ParseFloat(x, 0)
	return f, err
}

func RoundToString(decimal string, bit int) (i string) {
	f, err := strconv.ParseFloat(decimal, bit)
	if err != nil {
		logger.Error("RoundFloatString error", logger.String("str", decimal))
	}
	return strconv.FormatFloat(f, 'f', bit, 64)
}

func Copy(src interface{}, dest interface{}) error {
	bz, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bz, dest)
	if err != nil {
		return err
	}
	return nil
}

// x / y  and returns *big.Rat.
func QuoByStr(xStr, yStr string) (*big.Rat, error) {
	xAsRat, ok := new(big.Rat).SetString(xStr)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Convert string(%v) to big.Rat fail \n", xStr))
	}
	yAsRat, ok := new(big.Rat).SetString(yStr)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Convert string(%v) to big.Rat fail \n", yStr))
	}

	if yAsRat.Cmp(new(big.Rat).SetInt64(0)) != 1 {
		return nil, errors.New("yStr must != 0")
	}

	return new(big.Rat).Quo(xAsRat, yAsRat), nil
}

func NewRatFromFloat64(f float64) *big.Rat {
	return new(big.Rat).SetFloat64(f)
}

func ParseStringFromFloat64(data float64) string {
	return strconv.FormatFloat(data, 'f', -1, 64)
}

func MarshalJsonIgnoreErr(v interface{}) []byte {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		logger.Error("marshal json fail", logger.String("err", err.Error()), logger.Any("v", v))
	}

	return jsonBytes
}

func Md5Encryption(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func GetValaddr(address string) string {
	prefix, _, _ := DecodeAndConvert(address)
	if prefix == conf.Get().Hub.Prefix.ValAddr {
		return address
	} else {
		return Convert(conf.Get().Hub.Prefix.ValAddr, address)
	}
}

func FailtoFailed(status string) string {

	if status == StatusFail {
		return StatusFailed
	}
	return status
}