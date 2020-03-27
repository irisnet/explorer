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

func TestAccount_CountAll(t *testing.T) {
	d := Account{}

	if res, err := d.CountAll(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
}

func TestAccount_CountDelegatorNum(t *testing.T) {
	d := Account{}

	if res, err := d.CountDelegatorNum(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
}
