package document

import (
	"testing"
)

func TestGetValidatorTokenAndSelfBond(t *testing.T) {
	operatorAddr := "fva1kmjs6m4sfyacdm6sscmfgzpnslsw2mhc7wxgyn"

	tokens, selfBond, err := Delegator{}.GetValidatorTokenAndSelfBond(operatorAddr)
	if err != nil {
		t.Error(err)
	}

	t.Logf("tokens: %v  selfBond: %v \n", tokens, selfBond)

}

func TestGetDepositValidator(t *testing.T) {
	operatorAddr := "fva1kmjs6m4sfyacdm6sscmfgzpnslsw2mhc7wxgyn"

	validatorList, err := Delegator{}.GetDepositValidator([]string{operatorAddr})
	if err != nil {
		t.Error(err)
	}

	for k, v := range validatorList {
		t.Logf("k: %v  validator: %v \n", k, v)
	}

}
