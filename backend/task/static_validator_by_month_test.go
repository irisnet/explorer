package task

import (
	"testing"
	"encoding/json"
	"time"
	"github.com/irisnet/explorer/backend/types"
	"fmt"
	"github.com/irisnet/explorer/backend/orm/document"
)

var (
	task  StaticValidatorByMonthTask
	dtask StaticDelegatorByMonthTask
)

func init() {

	//starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 8), cstZone)
	//endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 9), cstZone)
	//terminaldata := document.ExStaticDelegator{
	//	Date:    endtime,
	//	Address: "iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut",
	//	DelegationsRewards: []document.Rewards{
	//		{Iris: float64(0.017621491086540352), IrisAtto: "17621491086540352"},
	//	},
	//	Delegation: utils.Coin{
	//		Amount: float64(1500000000000000000000),
	//		Denom:  "iris-atto",
	//	},
	//	Commission: []document.Rewards{
	//		{Iris: float64(0.188503181406925), IrisAtto: "188503181406924714"},
	//	},
	//	Total: []document.Rewards{
	//		{Iris: float64(0.206124672493465), IrisAtto: "206124672493465066"},
	//	},
	//}
	//txs, err := dtask.getPeriodTxByAddress(starttime, endtime, "iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut") //all address txs
	//if err != nil {
	//	fmt.Println(txs)
	//}
	//delegate_times := dtask.getPeriodDelegationTimes("iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut", txs)
	//fmt.Println(delegate_times)
	//res, err := dtask.getStaticDelegator(starttime, terminaldata, txs)
	//bytedatas, _ := json.Marshal(res)
	//fmt.Println(string(bytedatas))
}

func TestStaticValidatorByMonthTask_getStaticValidator(t *testing.T) {
	task.SetAddressCoinMapData(dtask.AddressCoin, dtask.AddrPeriodCommission, dtask.AddrTerminalCommission)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 8), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 9), cstZone)
	terminalData, err := task.staticModel.GetDataOneDay(starttime, "iva1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm64fepv")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(terminalData)
	if terminalData.OperatorAddress == "" {
		terminalData.OperatorAddress = "iva1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm64fepv"
	}
	terminalData.Date = endtime

	mapData, err := task.getCreateValidatorTx()
	if err != nil {
		t.Fatal(err.Error())
	}
	foundtionDelegation := task.getFoundtionDelegation([]document.ExStaticValidator{terminalData})

	res, err := task.getStaticValidator(starttime, terminalData, mapData, foundtionDelegation)
	if err != nil {
		t.Fatal(err.Error())
	}

	bytedatas, _ := json.Marshal(res)
	t.Log(string(bytedatas))
}

func TestStaticValidatorByMonthTask_getCreateValidatorTx(t *testing.T) {
	mapData, err := task.getCreateValidatorTx()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(mapData)
}

func TestStaticValidatorByMonthTask_Start(t *testing.T) {
	task.Start()
}

func TestStaticValidatorByMonthTask_DoTask(t *testing.T) {
	if err := task.DoTask(); err != nil {
		t.Fatal(err.Error())
	}
}

func TestMonthDoTask(t *testing.T) {
	//starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 24), cstZone)
	//endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 25), cstZone)
	delegatortask := new(StaticDelegatorByMonthTask)
	validatortask := new(StaticValidatorByMonthTask)
	//delegatortask.SetCaculateScope(starttime, endtime)
	//delegatortask.SetCaculateAddress("iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut")
	delegatortask.DoTask()
	fmt.Println("caculateWork have done,then validatortask work")
	//validatortask.SetCaculateAddress("iva1qq93sapmdcx36uz64vvw5gzuevtxsc7lcfxsat")
	//validatortask.SetCaculateScope(starttime, endtime)
	validatortask.SetAddressCoinMapData(delegatortask.AddressCoin, delegatortask.AddrTerminalCommission, delegatortask.AddrPeriodCommission)
	validatortask.DoTask()
}

func TestStaticValidatorByMonthTask_getCommissionRate(t *testing.T) {
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT18:40:00", 2020, 4, 27), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT18:49:00", 2020, 4, 27), cstZone)
	ratemax := task.getCommissionRate(starttime, endtime, "fva186qhtc62cf6ejlt3erw6zk28mgw8ne7grhmyfn", "-date")
	fmt.Println(ratemax)
	ratemin := task.getCommissionRate(starttime, endtime, "fva186qhtc62cf6ejlt3erw6zk28mgw8ne7grhmyfn", "date")
	fmt.Println(ratemin)
}
