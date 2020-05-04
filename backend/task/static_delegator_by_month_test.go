package task

import (
	"testing"
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/orm/document"
	"time"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/service"
	"gopkg.in/mgo.v2/bson"
	"github.com/irisnet/explorer/backend/vo"
	"math"
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
	starttime, err := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 1), cstZone)
	if err != nil {
		t.Fatal(err.Error())
	}
	endtime, err := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT23:59:59", 2020, 4, 10), cstZone)
	if err != nil {
		t.Fatal(err.Error())
	}
	txs, err := task.getPeriodTxByAddress(starttime, endtime, "")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(len(txs))
	fmt.Println(task.AddressCoin)
	fmt.Println(task.AddrPeriodCommission)

	//bytedata, _ := json.Marshal(txs)
	//t.Log(string(bytedata))
}

func TestStaticDelegatorByMonthTask_Caculate(t *testing.T) {
	task := new(StaticDelegatorByMonthTask)
	serv := new(service.CaculateService)

	starttime, err := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 1), cstZone)
	if err != nil {
		t.Fatal(err.Error())
	}
	endtime, err := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 10), cstZone)
	if err != nil {
		t.Fatal(err.Error())
	}
	txs, err := task.getPeriodTxByAddress(starttime, endtime, "")
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("txs num:", len(txs))
	fmt.Println("Rewards:", task.AddressCoin)
	fmt.Println("PeriodCommission:", task.AddrPeriodCommission)

	dbstarttime := fmt.Sprintf("%d.%02d.%02d", 2020, 4, 10)
	dbendtime := fmt.Sprintf("%d.%02d.%02d", 2020, 5, 1)

	cond := bson.M{}
	cond["date"] = bson.M{
		"$gte": dbstarttime,
		"$lt":  dbendtime,
	}
	var delegatorMonthData []vo.ExStaticDelegatorMonthVo
	var vaidatorMonthData []vo.ExStaticValidatorMonthVo
	page := 1
	limit := 100
	for {
		res, _, err := serv.GetDelegatorCaculateMonth(cond, page, limit, false)
		if err != nil {
			t.Log(err.Error())
			break
		}
		delegatorMonthData = append(delegatorMonthData, res...)
		if len(res) < limit {
			break
		}
		page++
	}

	for i, val := range delegatorMonthData {
		if data, ok := task.AddressCoin[val.Address]; ok {
			delegatorMonthData[i].PeriodWithdrawRewards += data.Amount / math.Pow10(18)
		}
	}
	bytedata, _ := json.Marshal(delegatorMonthData)
	t.Log(string(bytedata))

	page = 1
	for {
		res, _, err := serv.GetValidatorCaculateMonth(cond, page, limit, false)
		if err != nil {
			t.Log(err.Error())
			break
		}
		vaidatorMonthData = append(vaidatorMonthData, res...)
		if len(res) < limit {
			break
		}
		page++
	}

	for i, val := range vaidatorMonthData {
		if data, ok := task.AddrPeriodCommission[val.Address]; ok {
			vaidatorMonthData[i].PeriodCommission += data.Amount / math.Pow10(18)
		}
	}
	valbytedata, _ := json.Marshal(vaidatorMonthData)
	t.Log(string(valbytedata))

	//num, _ := task.getPeriodDelegationTimes(addr, txs)
}

func Test_parseCoinAmountAndUnitFromStr(t *testing.T) {
	coin := parseCoinAmountAndUnitFromStr("4.962968566351661iris-atto")
	//t.Log(coin.Amount)
	//t.Log(coin.Denom)
	coin1 := parseCoinAmountAndUnitFromStr("6.042877382409304iris-atto")
	coin2 := parseCoinAmountAndUnitFromStr("182.7501980494728iris-atto")
	//coin3 := parseCoinAmountAndUnitFromStr("5741584103447714iris-atto")
	//coin4 := parseCoinAmountAndUnitFromStr("18333217815328413661iris-atto")
	t.Log(coin.Amount + coin1.Amount + coin2.Amount)
	t.Log(coin.Denom)
}

func TestStaticDelegatorByMonthTask_getCoinflowByHash(t *testing.T) {
	s := new(StaticDelegatorByMonthTask)
	addr := "faa1eqvkfthtrr93g4p9qspp54w6dtjtrn279vcmpn"
	result := lcd.BlockCoinFlow("71D0D8A85EB74BD0FCFBF8F26D17784866ADF20F9BAA9A98A0E7190333FBF3EF")
	s.getCoinflow([]lcd.BlockCoinFlowVo{result})
	t.Log("71D0D8A85EB74BD0FCFBF8F26D17784866ADF20F9BAA9A98A0E7190333FBF3EF:", s.AddressCoin[addr])
	s.AddressCoin = nil
	result = lcd.BlockCoinFlow("6F99D1BD66A257E1C2F20E463BE53B651F7965EC0BF1D3A604D7B53CC9237646")
	s.getCoinflow([]lcd.BlockCoinFlowVo{result})
	t.Log("6F99D1BD66A257E1C2F20E463BE53B651F7965EC0BF1D3A604D7B53CC9237646:", s.AddressCoin[addr])
	s.AddressCoin = nil
	result = lcd.BlockCoinFlow("5415573E77894F38C25A8138A90FC00EE8DA8B28A28EF67CE65BF42492DCC310")
	s.getCoinflow([]lcd.BlockCoinFlowVo{result})
	t.Log("5415573E77894F38C25A8138A90FC00EE8DA8B28A28EF67CE65BF42492DCC310:", s.AddressCoin[addr])
}

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
