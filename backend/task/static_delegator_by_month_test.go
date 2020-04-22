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
	task := new(StaticDelegatorByMonthTask)
	starttime, err := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 9), cstZone)
	if err != nil {
		t.Fatal(err.Error())
	}
	endtime, err := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 15), cstZone)
	if err != nil {
		t.Fatal(err.Error())
	}
	txs, err := task.getPeriodTxByAddress(starttime, endtime, "faa1eqvkfthtrr93g4p9qspp54w6dtjtrn279vcmpn")
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
	coin := parseCoinAmountAndUnitFromStr("154914833557376583102iris-atto")
	//t.Log(coin.Amount)
	//t.Log(coin.Denom)
	//coin1 := parseCoinAmountAndUnitFromStr("219559815884266200iris-atto")
	//coin2 := parseCoinAmountAndUnitFromStr("383449763866744iris-atto")
	t.Log(coin.Amount)
	t.Log(coin.Denom)
}

func TestStaticDelegatorByMonthTask_getCoinflowByHash(t *testing.T) {
	s := new(StaticDelegatorByMonthTask)
	s.getCoinflowByHash("F40A07D5EDBDF2CF01F8EC746B7BBF7400BFA16B5DEBF57BFF2A15A82E5DBA29")
	t.Log(s.AddressCoin)
}

func TestStaticDelegatorByMonthTask_getStaticDelegator(t *testing.T) {
	s := new(StaticDelegatorByMonthTask)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 9), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 13), cstZone)
	terminaldata := document.ExStaticDelegator{
		Date:    endtime,
		Address: "faa1k2f9qpnh0r3j4r83z34hen4xtx5yd42crazf50",
		DelegationsRewards: []document.Rewards{
			{Iris: float64(404.745631836595), IrisAtto: ""},
		},
		Delegation: utils.Coin{
			Amount: float64(1500000000000000000000),
			Denom:  "iris-atto",
		},
		Commission: []document.Rewards{},
		Total: []document.Rewards{
			{Iris: float64(404.745631836595), IrisAtto: ""},
		},
	}
	txs, err := s.getPeriodTxByAddress(starttime, endtime, "faa1k2f9qpnh0r3j4r83z34hen4xtx5yd42crazf50") //all address txs
	if err != nil {
		t.Fatal(err.Error())
	}
	delegate_times := s.getPeriodDelegationTimes("faa1k2f9qpnh0r3j4r83z34hen4xtx5yd42crazf50", txs)
	t.Log(delegate_times)
	res, err := s.getStaticDelegator(starttime, terminaldata, txs)

	bytedata, _ := json.Marshal(res)
	t.Log(string(bytedata))

}

func TestStaticDelegatorByMonthTask_getIncrementDelegation(t *testing.T) {
	s := new(StaticDelegatorByMonthTask)
	value := []utils.Coin{
		{Amount: float64(112.3), Denom: "iris-atto"},
		{Amount: float64(100), Denom: "iris-atto"},
	}
	data := s.getIncrementDelegation(value[0], value[1])
	t.Log(data.Amount)
}

//func TestStart(t *testing.T) {
//	delegatortask := new(StaticDelegatorByMonthTask)
//	validatortask := new(StaticValidatorByMonthTask)
//	delegatortask.Start()
//	fmt.Println(delegatortask.AddressCoin)
//	validatortask.SetAddressCoinMapData(delegatortask.AddressCoin)
//	validatortask.Start()
//}
