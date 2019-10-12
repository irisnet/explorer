package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"encoding/json"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

func TestGetDelegationsFromLcd(t *testing.T) {

	delegationPage := new(ValidatorService).GetDelegationsFromLcd("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx", 1, 5, false, true)
	t.Logf(" %v \n", delegationPage)
}

func TestQueryValidatorTopN(t *testing.T) {

	validatorPage := new(ValidatorService).QueryCandidatesTopN()

	t.Logf("total: %v \n", validatorPage.Count)
	t.Logf("power all: %v \n", validatorPage.PowerAll)

	for k, v := range validatorPage.Validators {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestConvert(t *testing.T) {

	conPubKeys := []string{"fcp1zcjduepqevwqk73gun8pp59wz6raddnsg2fczvs237cefl9ve7f94feh6lzsdr4qrf", "fcp1zcjduepqcjgmderfahnlyrse563r2hcc3d4vjpafw03axzn3e87kfuqznjcsur8xrq", "fcp1zcjduepq3dgenw6hha8kh0r4z5emqqv7u2k2w6qevytxwgdqkyugrsqmyykqq0zycl",
		"fcp1zcjduepq3hy0q3ltgktya5gzvjfmkrhkeqzqau9t927ne8wdts36en490q7s8lxjq2", "fcp1zcjduepqgs9ffj72djsf7vytvm8vq3d2gzavytw0pkl5lee0dp4maej97x9swzlt60", "fcp1zcjduepqujegylnmyuymekm7sk70g5up0xq6p5r8zaj5rpjey7tqxxp0z3zqjns5cy"}

	for k, v := range conPubKeys {

		hrp, bytes, err := utils.DecodeAndConvert(v)
		if err != nil {
			logger.Error("DecodeAndConvert", logger.String("err", err.Error()), logger.String("param", v))
			continue
		}
		pubkey := strings.ToUpper(hex.EncodeToString(bytes))

		// h := sha256.New()
		// h.Write([]byte(pubkey))
		// sum := h.Sum(nil)

		t.Logf("idx: %v hrp: %v conPubKey: %v  pubKey: %v  hex: %v  \n", k, hrp, v, pubkey, utils.GenHexAddrFromPubKey(v))
	}
}

func TestConvertConsensusPublicKey(t *testing.T) {

	consensusPubKey := "1624DE6420CB1C0B7A28E4CE10D0AE1687D6B670429381320A8FB194FCACCF925AA737D7C5"

	h := sha256.New()
	h.Write([]byte(consensusPubKey))
	sum := h.Sum(nil)

	fmt.Printf("%s \n", consensusPubKey)
	fmt.Printf("%v \n", sum)
	fmt.Printf("%x \n", sum)

	t.Logf("  consensus pubKey: %v  \n sum: %v  \n sumAsHex: %x \n", consensusPubKey, sum, sum)

}

func TestGetValidators(t *testing.T) {

	validatorList := new(ValidatorService).GetValidators("jailed", "browser", 0, 100, true)

	//res := validatorList.([]lcd.ValidatorVo)
	resBytes, _ := json.Marshal(validatorList)
	t.Log(string(resBytes))
}

func TestValidatorService_UpdateValidatorIcons(t *testing.T) {
	err := new(ValidatorService).UpdateValidatorIcons()
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidatorService_UpdateValidatorStaticData(t *testing.T) {
	if err := validatorService.UpdateValidatorStaticInfo(); err != nil {
		t.Fatal(err)
	} else {
		t.Log("success")
	}
}

func TestBuildValidators(t *testing.T) {
	res := buildValidators()
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}


func TestValidatorService_GetValidatorDetail(t *testing.T) {
	res := validatorService.GetValidatorDetail("fva10jv6pkdtjc39pwpjxnurqend0p09gphlyhzmel")
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}