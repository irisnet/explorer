package document

import (
	"testing"
	"encoding/json"
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

//func TestValidator_QueryValidatorMonikerByAddrArr(t *testing.T) {
//	valaddr := []string{"iva1r2pq5y674llvd654tr9lm7s68wnumd0pf04v99"}
//	data, err := Validator{}.QueryValidatorMonikerByAddrArr(valaddr)
//	if err != nil {
//		t.Fatal(err)
//	}
//	t.Log(data)
//}

func TestValidator_UpdateByPk(t *testing.T) {
	validatorModel := Validator{}
	validators, err := validatorModel.GetAllValidator()
	if err != nil {
		t.Fatal(err)
	}
	validator := validators[0]
	validator.Uptime = float32(0.75)

	if err := validatorModel.UpdateByPk(validator); err != nil {
		t.Fatal(err)
	} else {
		t.Log("success")
	}
}

func TestValidator_QueryValidatorDescription(t *testing.T) {
	data := Validator{}.QueryValidatorDescription()
	datamsg, _ := json.Marshal(data)
	t.Log(string(datamsg))
}