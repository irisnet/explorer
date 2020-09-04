package lcd

import (
	"encoding/json"
	"testing"
	"strconv"
	"fmt"
)

func TestGetRedelegationsByValidatorAddr(t *testing.T) {

	redelegations := GetRedelegationsByValidatorAddr("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx")

	for k, v := range redelegations {
		t.Logf("k: %v v: %v \n", k, v)
	}

}

func TestGetDistributionRewardsByValidatorAcc(t *testing.T) {

	rewards,err := GetDistributionCommissionRewardsByAddress("iva1x98k5n7xj0h3udnf5dcdzw85tsfa75qm682jtg")
	if err != nil {
		t.Error(err)
	}

	for k, v := range rewards {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestGetJailedUntilAndMissedBlocksCountByConsensusPublicKey(t *testing.T) {
	jailedUntil, missedBlockCount, _, err := GetJailedUntilAndMissedBlocksCountByConsensusPublicKey("icp1zcjduepqwuhmlewmrr0hvfcjkccm7qk7kektujr8ten0tve0rd9eqexdee2qujw3vv")

	if err != nil {
		t.Error(err)
	}

	t.Logf("jailedUntil: %v  missedBlockCount: %v \n", jailedUntil, missedBlockCount)
}

func TestValidator(t *testing.T) {
	address := "iva1x98k5n7xj0h3udnf5dcdzw85tsfa75qm682jtg"

	if res, err := Validator(address); err != nil {
		t.Fatal(err)
	} else {
		resBytes, _ := json.MarshalIndent(res, "", "\t")
		t.Log(string(resBytes))
	}
}

func TestValidators(t *testing.T) {
	res := Validators(1, 10)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}



func TestGetDelegationsByDelAddr(t *testing.T) {
	address := "iaa17vkfw608qg2fpkrguk7rmwl04hgf7j7fv8740q"
	res := GetDelegationsByDelAddr(address)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestDelegationByValidator(t *testing.T) {
	address := "fva1k98nqnytl9u5ralns7ca7n8tpmacl8k25ymgyr"
	res := GetDelegationByValidator(address)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestStakePool(t *testing.T) {
	res := StakePool()
	a, _ := strconv.ParseFloat(res.BondedTokens, 64)
	b, _ := strconv.ParseFloat(res.LooseTokens, 64)
	fmt.Println(a + b)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestSignInfo(t *testing.T) {
	pubKey := "icp1wuhmlewmrr0hvfcjkccm7qk7kektujr8ten0tve0rd9eqexdee2qrm74re"
	res := SignInfo(pubKey)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}



func TestGetUnbondingDelegationsByValidatorAddr(t *testing.T) {

	unbondingDelegations := GetUnbondingDelegationsByValidatorAddr("iva1x98k5n7xj0h3udnf5dcdzw85tsfa75qm682jtg")
	for k, v := range unbondingDelegations {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestGetUnbondingDelegationsByDelegatorAddr(t *testing.T) {
	unbondingDelegations := GetUnbondingDelegationsByDelegatorAddr("iaa1ryyqz07d75fx20hk9npmnan9p0jn9yw35vpk6k")
	for k, v := range unbondingDelegations {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}
