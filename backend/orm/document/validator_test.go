package document

import (
	"testing"
)

func TestGetAllValidator(t *testing.T) {
	validatorList, err := Validator{}.GetAllValidator()
	if err != nil {
		t.Error(err)
	}

	for k, v := range validatorList {
		t.Logf("idx: %v  v: %v \n", k, v)
	}

}

func TestValidatorGetCandidatesTopN(t *testing.T) {

	validators, power, upTimeMap, err := Validator{}.GetCandidatesTopN()

	if err != nil {
		t.Error(err)
	}

	t.Logf("power: %v \n", power)

	t.Logf("validators------------")
	for k, v := range validators {
		t.Logf("k: %v  v: %v \n", k, v)
	}

	t.Log("uptime map-------------")

	for k, v := range upTimeMap {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestQueryCandidateUptimeWithHour(t *testing.T) {

	//db.getCollection('ex_uptime_change').find({"address":"35609FEA8D3401EFDC143D42FD9D4C95FBFC6A92","time":{"$gte":"2019-05-17 06"}})
	history, err := Validator{}.QueryCandidateUptimeWithHour("BF87BDA76737C1820D3D4DF5D753A57CC0D6837F")

	if err != nil {
		t.Error(err)
	}

	for k, v := range history {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}
