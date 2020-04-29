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

//func TestStaticDelegatorByMonthTask_getDelegationData(t *testing.T) {
//	res, err := new(StaticDelegatorByMonthTask).getDelegationData("2020-04-03T00:00:00")
//	if err != nil {
//		t.Fatal(err)
//	}
//	bytedata, _ := json.Marshal(res)
//	t.Log(string(bytedata))
//}
func TestStaticDelegatorByMonthTask_getPeriodTxByAddress(t *testing.T) {
	task := new(StaticDelegatorByMonthTask)
	starttime, err := time.Parse(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT17:15:00", 2020, 4, 28))
	if err != nil {
		t.Fatal(err.Error())
	}
	endtime, err := time.Parse(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT17:35:00", 2020, 4, 28))
	if err != nil {
		t.Fatal(err.Error())
	}
	txs, err := task.getPeriodTxByAddress(starttime, endtime, "faa1clxzr9f7sg2fqnr2zl2eswc5s82md85hsjl8vz")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(len(txs))
	fmt.Println(task.AddressCoin)
	fmt.Println(task.AddrPeriodCommission)

	bytedata, _ := json.Marshal(txs)
	t.Log(string(bytedata))
}

func Test_parseCoinAmountAndUnitFromStr(t *testing.T) {
	coin := parseCoinAmountAndUnitFromStr("39007580267975728iris-atto")
	//t.Log(coin.Amount)
	//t.Log(coin.Denom)
	coin1 := parseCoinAmountAndUnitFromStr("587702808887031233613iris-atto")
	coin2 := parseCoinAmountAndUnitFromStr("302757713438089471198iris-atto")
	t.Log(coin.Amount + coin1.Amount + coin2.Amount)
	t.Log(coin.Denom)
}

//func TestStaticDelegatorByMonthTask_getCoinflowByHash(t *testing.T) {
//	s := new(StaticDelegatorByMonthTask)
//	s.getCoinflowByHash("F40A07D5EDBDF2CF01F8EC746B7BBF7400BFA16B5DEBF57BFF2A15A82E5DBA29",nil)
//	t.Log(s.AddressCoin)
//}

func TestStaticDelegatorByMonthTask_getStaticDelegator(t *testing.T) {
	s := new(StaticDelegatorByMonthTask)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 7), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 8), cstZone)
	terminaldata := document.ExStaticDelegator{
		Date:    endtime,
		Address: "iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut",
		DelegationsRewards: []document.Rewards{
			{Iris: float64(0.014716424936353884), IrisAtto: ""},
		},
		Delegation: utils.Coin{
			Amount: float64(1500000000000000000000),
			Denom:  "iris-atto",
		},
		Commission: []document.Rewards{
			{Iris: float64(0.167533025538185)},
		},
		Total: []document.Rewards{
			{Iris: float64(0.182331089770524), IrisAtto: ""},
		},
	}
	txs, err := s.getPeriodTxByAddress(starttime, endtime, "iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut") //all address txs
	if err != nil {
		t.Fatal(err.Error())
	}
	delegate_times := s.getPeriodDelegationTimes("iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut", txs)
	t.Log(delegate_times)
	res, err := s.getStaticDelegator(starttime, terminaldata, txs)

	bytedata, _ := json.Marshal(res)
	t.Log(string(bytedata))

}

func TestStaticDelegatorByMonthTask_getIncrementDelegation(t *testing.T) {
	s := new(StaticDelegatorByMonthTask)
	value := utils.Coin{Amount: float64(112.3), Denom: "iris-atto"}
	value1 := document.Coin{Amount: float64(100), Denom: "iris-atto"}

	data := s.getIncrementDelegation(value, value1)
	t.Log(data.Amount)
}

func TestDoTask(t *testing.T) {
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 4), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 5), cstZone)
	delegatortask := new(StaticDelegatorByMonthTask)
	//delegatortask.SetCaculateAddress("iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut")
	delegatortask.SetCaculateScope(starttime, endtime)
	delegatortask.DoTask()
}
