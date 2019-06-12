package service

import (
	"testing"

	"github.com/irisnet/explorer/backend/lcd"
)

func TestGetValidators(t *testing.T) {

	validatorList := new(CandidateService).GetValidators("", "", 0, 100)

	if validatorListAsLcd, ok := validatorList.([]lcd.ValidatorVo); ok {
		for k, v := range validatorListAsLcd {
			t.Logf("k: %v    v: %v \n", k, v)
		}
	}
}

func TestGetValidator(t *testing.T) {

	validator := new(CandidateService).GetValidator("fva1amca8h9msuee5j8pg54q3sc9n9p2ss060wvhd5")

	t.Logf("validator: %v \n", validator)

}

func TestQueryValidatorByConAddr(t *testing.T) {
	validator := new(CandidateService).QueryValidatorByConAddr("BF87BDA76737C1820D3D4DF5D753A57CC0D6837F")
	t.Logf("validator: %v \n", validator)
}

func TestQueryCandidate(t *testing.T) {
	candidator := new(CandidateService).QueryCandidate("fva1amca8h9msuee5j8pg54q3sc9n9p2ss060wvhd5")
	t.Logf("candidator: %v \n", candidator)
}

func TestQueryCandidatesTopN(t *testing.T) {

	val := new(CandidateService).QueryCandidatesTopN()

	t.Logf("validator: %v \n", val)
}

func TestQueryCandidateUptime(t *testing.T) {

	unitsOfTime := []string{"hour", "week", "month"}

	for _, unit := range unitsOfTime {

		uptimeHistory := new(CandidateService).QueryCandidateUptime("fva1d7rxfhyhqzgudv5nmne548t3xrk4tvazpg94cu", unit)

		for k, v := range uptimeHistory {
			t.Logf("unit: %v k: %v  v: %v  \n", unit, k, v)
		}
	}
}

func TestQueryCandidatePower(t *testing.T) {
	unitsOfTime := []string{"months", "week", "month"}

	for _, unit := range unitsOfTime {
		powerHistory := new(CandidateService).QueryCandidatePower("fva1d7rxfhyhqzgudv5nmne548t3xrk4tvazpg94cu", unit)

		for k, v := range powerHistory {
			t.Logf("unit: %v  k: %v  v: %v", unit, k, v)
		}
	}
}

func TestQueryCandidateStatus(t *testing.T) {

	status := new(CandidateService).QueryCandidateStatus("fva1d7rxfhyhqzgudv5nmne548t3xrk4tvazpg94cu")

	t.Logf("status: %v \n", status)

}
