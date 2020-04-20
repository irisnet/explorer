package task

import (
	"testing"
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/orm/document"
	"time"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
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
	txs, err := new(StaticDelegatorByMonthTask).getPeriodTxByAddress(2020, 4, "faa12t4gfg502wra9lhtjjvqudq82rrzu2sk5j2l09")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(len(txs))

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

func TestStaticDelegatorByMonthTask_getStaticDelegator(t *testing.T) {
	s := new(StaticDelegatorByMonthTask)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-01T00:00:00", 2020, 4), cstZone)
	terminaldate := document.ExStaticDelegator{
		Date:    starttime,
		Address: "faa12t4gfg502wra9lhtjjvqudq82rrzu2sk5j2l09",
		DelegationsRewards: []document.Rewards{
			{Iris: float64(237.866927663832), IrisAtto: "237866927663831973898"},
		},
		Delegation: utils.Coin{
			Amount: float64(1150000000000000000000),
			Denom:  "iris-atto",
		},
		Commission: []document.Rewards{},
		Total: []document.Rewards{
			{IrisAtto: "237866927663831973898", Iris: float64(237.866927663832)},
		},
	}
	txs, err := s.getPeriodTxByAddress(2020, 4, "faa12t4gfg502wra9lhtjjvqudq82rrzu2sk5j2l09") //all address txs
	if err != nil {
		t.Fatal(err.Error())
	}
	delegate_times := s.getPeriodDelegationTimes("faa12t4gfg502wra9lhtjjvqudq82rrzu2sk5j2l09", txs)
	t.Log(delegate_times)
	res := s.getStaticDelegator(terminaldate, txs)

	bytedata, _ := json.Marshal(res)
	t.Log(string(bytedata))

}

//func TestStart(t *testing.T) {
//	delegatortask := new(StaticDelegatorByMonthTask)
//	validatortask := new(StaticValidatorByMonthTask)
//	delegatortask.Start()
//	fmt.Println(delegatortask.AddressCoin)
//	validatortask.SetAddressCoinMapData(delegatortask.AddressCoin)
//	validatortask.Start()
//}
