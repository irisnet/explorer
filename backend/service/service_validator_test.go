package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/logger"
)

func TestGetDelegationsFromLcd(t *testing.T) {

	delegationPage := new(ValidatorService).GetDelegationsFromLcd("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx", 1, 5)
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

func TestQueryCandidateUptime(t *testing.T) {

	unitsOfTime := []string{"hour", "week", "month"}

	for _, unit := range unitsOfTime {

		uptimeHistory := new(ValidatorService).QueryCandidateUptime("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx", unit)

		for k, v := range uptimeHistory {
			t.Logf("unit: %v k: %v  v: %v  \n", unit, k, v)
		}
	}
}

// func TestGetValidators(t *testing.T) {
//
// 	validatorList := new(ValidatorService).GetValidators("", "", 0, 100)
//
// 	if validatorListAsLcd, ok := validatorList.([]lcd.ValidatorVo); ok {
// 		for k, v := range validatorListAsLcd {
// 			t.Logf("k: %v    v: %v \n", k, v)
// 		}
// 	}
// }
//
// func TestGetValidator(t *testing.T) {
//
// 	validator := new(ValidatorService).GetValidator("fva1amca8h9msuee5j8pg54q3sc9n9p2ss060wvhd5")
//
// 	t.Logf("validator: %v \n", validator)
//
// }
//
// func TestQueryValidatorByConAddr(t *testing.T) {
// 	validator := new(ValidatorService).QueryValidatorByConAddr("BF87BDA76737C1820D3D4DF5D753A57CC0D6837F")
// 	t.Logf("validator: %v \n", validator)
// }
//
// func TestQueryCandidate(t *testing.T) {
// 	candidator := new(ValidatorService).QueryCandidate("fva1amca8h9msuee5j8pg54q3sc9n9p2ss060wvhd5")
// 	t.Logf("candidator: %v \n", candidator)
// }
//
// func TestQueryCandidatesTopN(t *testing.T) {
//
// 	val := new(ValidatorService).QueryCandidatesTopN()
//
// 	t.Logf("validator: %v \n", val)
// }
//

//
// func TestQueryCandidatePower(t *testing.T) {
// 	unitsOfTime := []string{"months", "week", "month"}
//
// 	for _, unit := range unitsOfTime {
// 		powerHistory := new(ValidatorService).QueryCandidatePower("fva1d7rxfhyhqzgudv5nmne548t3xrk4tvazpg94cu", unit)
//
// 		for k, v := range powerHistory {
// 			t.Logf("unit: %v  k: %v  v: %v", unit, k, v)
// 		}
// 	}
// }
//
// func TestQueryCandidateStatus(t *testing.T) {
//
// 	status := new(ValidatorService).QueryCandidateStatus("fva1d7rxfhyhqzgudv5nmne548t3xrk4tvazpg94cu")
//
// 	t.Logf("status: %v \n", status)
//
// }
