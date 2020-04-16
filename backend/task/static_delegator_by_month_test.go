package task

import (
	"testing"
	"encoding/json"
)

func TestStaticDelegatorByMonthTask_Start(t *testing.T) {
	new(StaticDelegatorByMonthTask).Start()
}

func TestStaticDelegatorByMonthTask_DoTask(t *testing.T) {
	if err := new(StaticDelegatorByMonthTask).DoTask(); err != nil {
		t.Fatal(err.Error())
	}
}

func TestStaticDelegatorByMonthTask_getDelegationData(t *testing.T) {
	res, err := new(StaticDelegatorByMonthTask).getDelegationData("2020-04-03T00:00:00")
	if err != nil {
		t.Fatal(err)
	}
	bytedata, _ := json.Marshal(res)
	t.Log(string(bytedata))
}
func TestStaticDelegatorByMonthTask_getPeriodTxByAddress(t *testing.T) {
	txs, err := new(StaticDelegatorByMonthTask).getPeriodTxByAddress(2019, 11, "faa174qyl02cupyqq77cqqtdl0frda6dl3rpjcrgnp")
	if err != nil {
		t.Fatal(err.Error())
	}

	bytedata, _ := json.Marshal(txs)
	t.Log(string(bytedata))
}

func Test_parseCoinAmountAndUnitFromStr(t *testing.T) {
	coin := parseCoinAmountAndUnitFromStr("666575029847018882iris-atto")
	t.Log(coin.Amount)
	t.Log(coin.Denom)
}

func TestStaticDelegatorByMonthTask_getCoinflowByHash(t *testing.T) {
	s := new(StaticDelegatorByMonthTask)
	s.getCoinflowByHash("DD5011DA37A00DB4EBB1F60A3F7DA8422F1553BA6E7C5C4FC9EDC38D22C5BB70")
	t.Log(s.AddressCoin)
}

//func TestStart(t *testing.T) {
//	delegatortask := new(StaticDelegatorByMonthTask)
//	validatortask := new(StaticValidatorByMonthTask)
//	delegatortask.Start()
//	fmt.Println(delegatortask.AddressCoin)
//	validatortask.SetAddressCoinMapData(delegatortask.AddressCoin)
//	validatortask.Start()
//}
