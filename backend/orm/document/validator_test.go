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
