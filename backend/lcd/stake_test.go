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

	jailedUntil, missedBlockCount, _, err := GetJailedUntilAndMissedBlocksCountByConsensusPublicKey("fcp1zcjduepqcjgmderfahnlyrse563r2hcc3d4vjpafw03axzn3e87kfuqznjcsur8xrq")

	if err != nil {
		t.Error(err)
	}

	t.Logf("jailedUntil: %v  missedBlockCount: %v \n", jailedUntil, missedBlockCount)
}

func TestValidator(t *testing.T) {
	address := "fva1k98nqnytl9u5ralns7ca7n8tpmacl8k25ymgyr"

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
	address := "faa192vef4442d07lqde59mx35dvmfv9v72wrsu84a"
	res := GetDelegationsByDelAddr(address)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestDelegationByValidator(t *testing.T) {
	address := "fva1k98nqnytl9u5ralns7ca7n8tpmacl8k25ymgyr"
	res := DelegationByValidator(address)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestStakePool(t *testing.T) {
	res := StakePool()
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestSignInfo(t *testing.T) {
	pubKey := "fcp1zcjduepqy5ygunecsppr7ye9fnjm3tsd82k7t3pmaqcpk3z9tegwz2kaxctsc3d4cj"
	res := SignInfo(pubKey)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestGetDelegationsByValidatorAddr(t *testing.T) {

	delegations := GetDelegationsByValidatorAddr("fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx")

	for k, v := range delegations {
		t.Logf("k: %v  v: %v \n", k, v)
	}
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
