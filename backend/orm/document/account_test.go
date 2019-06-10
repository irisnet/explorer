package document

import (
	"testing"
)

func TestGetAccountList(t *testing.T) {

	accList, err := Account{}.GetAccountList()
	if err != nil {
		t.Error(err)
	}

	for k, v := range accList {
		t.Logf("k: %v   acc: %v \n", k, v)
	}

}
