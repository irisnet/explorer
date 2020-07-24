package lcd

import (
	"encoding/json"
	"testing"
)

func TestGetRedelegationsByValidatorAddr(t *testing.T) {

	redelegations := GetRedelegationsByValidatorAddr("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx")

	for k, v := range redelegations {
		t.Logf("k: %v v: %v \n", k, v)
	}

}

func TestGetDistributionRewardsByValidatorAcc(t *testing.T) {

	rewards, _, _, err := GetDistributionRewardsByValidatorAcc("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx")
	if err != nil {
		t.Error(err)
	}

	for k, v := range rewards {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestGetJailedUntilAndMissedBlocksCountByConsensusPublicKey(t *testing.T) {
	jailedUntil, missedBlockCount, _, err := GetJailedUntilAndMissedBlocksCountByConsensusPublicKey("iva1x98k5n7xj0h3udnf5dcdzw85tsfa75qm682jtg")

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

func TestQueryWithdrawAddr(t *testing.T) {
	address := "faa1k98nqnytl9u5ralns7ca7n8tpmacl8k2p438ey"
	res := QueryWithdrawAddr(address)
	t.Log(res)
}

func TestGetDelegationsByDelAddr(t *testing.T) {
	address := "iaa18e2e9fxxrr88k78gg7fhuuqgccfv8self9ye65"
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
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestSignInfo(t *testing.T) {
	pubKey := "icp1zcjduepq6tfg639yglaxnxxst8mk058g86zuqlvnr0ql7d6setzj0x0flrysr6de8x"
	res := SignInfo(pubKey)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}



func TestGetUnbondingDelegationsByValidatorAddr(t *testing.T) {

	unbondingDelegations := GetUnbondingDelegationsByValidatorAddr("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx")
	for k, v := range unbondingDelegations {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestGetUnbondingDelegationsByDelegatorAddr(t *testing.T) {
	unbondingDelegations := GetUnbondingDelegationsByDelegatorAddr("faa1eqvkfthtrr93g4p9qspp54w6dtjtrn279vcmpn")
	for k, v := range unbondingDelegations {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}
