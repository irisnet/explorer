package document

import (
	"testing"
	"time"

	"github.com/irisnet/explorer/backend/types"
)

func TestGetValidatorListByPage(t *testing.T) {

	for k, v := range types.RoleList {
		total, validatorList, err := Candidate{}.GetValidatorListByPage(v, 0, 100)
		if err != nil {
			t.Error(err)
		}

		t.Logf("\n\nk: %v role: %v total: %v \n", k, v, total)
		for valK, val := range validatorList {
			t.Logf("validator key: %v   validator:\n %v \n", valK, val)
		}
	}
}

func TestGetCandidatesTopN(t *testing.T) {

	candidates, power, upTimeMap, err := Candidate{}.GetCandidatesTopN()
	if err != nil {
		t.Error(err)
	}

	t.Logf("power: %v \n", power)

	for k, v := range candidates {
		t.Logf("k: %v  candidator: \n %v \n", k, v)
	}

	for k, v := range upTimeMap {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}

func TestGetCandidatePubKeyAddrByAddr(t *testing.T) {

	candidatorAddr := "fva1xtstdchjyzkddaptgyug62g23cta7eyzq49svq"

	publicKey, err := Candidate{}.GetCandidatePubKeyAddrByAddr(candidatorAddr)
	if err != nil {
		t.Error(err)
	}
	t.Logf("candidator addr: %v  public key: %v \n", candidatorAddr, publicKey)

}

func TestQueryCandidateUptimeWithHour(t *testing.T) {

	candidatorAddr := "fva1xtstdchjyzkddaptgyug62g23cta7eyzq49svq"
	uptime, err := Candidate{}.QueryCandidateUptimeWithHour(candidatorAddr)

	if err != nil {
		t.Error(err)
	}

	for k, v := range uptime {
		t.Logf("k: %v v: %v \n", k, v)
	}
}

func TestQueryCandidateUptimeByWeekOrMonth(t *testing.T) {

	candidatorAddr := "fva1xtstdchjyzkddaptgyug62g23cta7eyzq49svq"

	uptimeList, err := Candidate{}.QueryCandidateUptimeByWeekOrMonth(candidatorAddr, "")

	if err != nil {
		t.Error(err)
	}

	for k, v := range uptimeList {
		t.Logf("k: %v   v: %v \n", k, v)
	}

	uptimeList, err = Candidate{}.QueryCandidateUptimeByWeekOrMonth(candidatorAddr, "month")

	if err != nil {
		t.Error(err)
	}

	for k, v := range uptimeList {
		t.Logf("k: %v   v: %v \n", k, v)
	}
}

func TestQueryCandidatePower(t *testing.T) {

	candidatorAddr := "fva1xtstdchjyzkddaptgyug62g23cta7eyzq49svq"

	votingPowerHistory, err := Candidate{}.QueryCandidatePower(candidatorAddr, time.Now().Add(time.Hour*(-24)).String())

	if err != nil {
		t.Error(err)
	}

	for k, v := range votingPowerHistory {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestQueryCandidateStatus(t *testing.T) {

	candidatorAddr := "fva1xtstdchjyzkddaptgyug62g23cta7eyzq49svq"
	precommitCount, uptime, err := Candidate{}.QueryCandidateStatus(candidatorAddr)
	if err != nil {
		t.Error(err)
	}

	t.Logf("precommitCount: %v   uptime: %v \n", precommitCount, uptime)

}

func TestQueryCandidateListByAddrList(t *testing.T) {

	candidatorList, err := Candidate{}.QueryCandidateListByAddrList([]string{"fva1xtstdchjyzkddaptgyug62g23cta7eyzq49svq"})

	if err != nil {
		t.Error(err)
	}

	for k, v := range candidatorList {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestQueryValidatorByConsensusAddr(t *testing.T) {

	validator, err := Candidate{}.QueryValidatorByConsensusAddr("BF87BDA76737C1820D3D4DF5D753A57CC0D6837F")
	if err != nil {
		t.Error(err)
	}
	t.Logf("validator: %v \n", validator)
}

func TestQueryPowerWithBonded(t *testing.T) {

	count, err := Candidate{}.QueryPowerWithBonded()

	if err != nil {
		t.Error(err)
	}

	t.Logf("count: %v \n", count)
}
