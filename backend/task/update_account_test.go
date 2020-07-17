package task

import (
	"testing"
	"github.com/irisnet/explorer/backend/orm/document"
	"encoding/json"
)

func TestUpdateAccount_Start(t *testing.T) {
	new(UpdateAccount).Start()
}

var account document.Account

func TestMain(m *testing.M) {
	accs, err := new(document.Account).GetAllAccount()
	if err != nil {
		panic(err)
	}

	account = accs[0]
	m.Run()
}
func TestUpdateAccount_DoTask(t *testing.T) {
	new(UpdateAccount).DoTask(HeartBeat)

}

func TestUpdateAccount_updateAccountInfo(t *testing.T) {
	account, _ := getAccountInfo(&account)
	bytedata, _ := json.Marshal(account)
	t.Log(string(bytedata))
}

func TestUpdateAccount_updateAccount(t *testing.T) {
	updateAccount(&account)
	bytedata, _ := json.Marshal(account)
	t.Log(string(bytedata))
}

func TestUpdateAccount_getBalance(t *testing.T) {
	v, acc, _ := getBalance(&account)
	bytedata, _ := json.Marshal(acc)
	t.Log(string(bytedata))
	t.Log(v)
}



func TestUpdateAccount_getDelegationInfo(t *testing.T) {
	ret, _ := getDelegationInfo(account.Address)
	t.Log(ret)
}

func TestUpdateAccount_getUnbondingDelegationInfo(t *testing.T) {
	ret, _ := getUnbondingDelegationInfo(account.Address)
	t.Log(ret)
}
