package task

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/vo"
	"testing"
	"time"
)

var (
	task  StaticValidatorByMonthTask
	dtask StaticDelegatorByMonthTask
)

func init() {

	//starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 1), cstZone)
	//endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 30), cstZone)
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
	//terminalData, err := dtask.staticModel.GetDataByDate(endtime)
	//if err != nil {
	//	panic(err.Error())
	//}
	//var terminaldata document.ExStaticDelegator
	//for _, val := range terminalData {
	//	if val.Address == address {
	//		terminaldata = val
	//		fmt.Println(terminaldata.Commission)
	//	}
	//}
	////terminaldata := terminalData[0]
	//txs, err := dtask.getPeriodTxByAddress(starttime, endtime, address) //all address txs
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(dtask.AddressCoin)
	//fmt.Println(dtask.AddrPeriodCommission)
	//delegate_times := dtask.getPeriodDelegationTimes(address, txs)
	//fmt.Println(delegate_times)
	//res, err := dtask.getStaticDelegator(starttime, terminaldata, txs)
	//bytedatas, _ := json.Marshal(res)
	//fmt.Println(string(bytedatas))
}

func TestStaticValidatorByMonthTask_getStaticValidator(t *testing.T) {
	task.SetAddressCoinMapData(dtask.AddressCoin, dtask.AddrPeriodCommission, dtask.AddrBeginCommission, dtask.AddrTerminalCommission)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 5, 1), cstZone)
	datetime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 5, 31), cstZone)
	endtime, createat, err := document.Getdate(task.staticModel.Name(), starttime, datetime, "-"+document.ExStaticDelegatorDateTag)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(time.Unix(createat, 0))
	fmt.Println(endtime)
	terminalData, err := task.staticModel.GetDataByDate(endtime)
	if err != nil {
		t.Fatal(err.Error())
	}

	mapData, err := task.getCreateValidatorTx()
	if err != nil {
		t.Fatal(err.Error())
	}
	foundtionDelegation := task.getFoundtionDelegation()

	for _, val := range terminalData {

		res, err := task.getStaticValidator(starttime, val, mapData, foundtionDelegation)
		if err != nil {
			t.Fatal(err.Error())
		}

		bytedatas, _ := json.Marshal(res)
		t.Log(string(bytedatas))
	}

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
	if err := task.DoTask(HeartBeat); err != nil {
		t.Fatal(err.Error())
	}
}

func TestMonthDoTask(t *testing.T) {
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 6, 1), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 7, 1), cstZone)
	delegatortask := new(StaticDelegatorByMonthTask)
	validatortask := new(StaticValidatorByMonthTask)
	delegatortask.SetCaculateScope(starttime, endtime)
	//delegatortask.SetCaculateAddress("iaa1xf5jaw09klqg9hzxfks3ycjvqgnpyjcm0yrkut")
	data, err := delegatortask.calculateWork()
	if err != nil {
		t.Fatal(err.Error())
	}
	var vomodel []vo.ExStaticDelegatorMonthVo

	for _, val := range data {
		ite := service.LoadFromDelModel(val)
		vomodel = append(vomodel, ite)
	}
	bytedata, _ := json.Marshal(vomodel)
	fmt.Println(string(bytedata))
	fmt.Println("calculateWork have done,then validatortask work")
	//validatortask.SetCaculateAddress("iva1qq93sapmdcx36uz64vvw5gzuevtxsc7lcfxsat")
	validatortask.SetCaculateScope(starttime, endtime)
	validatortask.SetAddressCoinMapData(delegatortask.AddressCoin, delegatortask.AddrPeriodCommission, delegatortask.AddrBeginCommission, delegatortask.AddrTerminalCommission)
	valdata, err := validatortask.caculateWork()
	if err != nil {
		t.Fatal(err.Error())
	}
	var vamodel []vo.ExStaticValidatorMonthVo

	for _, val := range valdata {
		ite := service.LoadFromValModel(val)
		vamodel = append(vamodel, ite)
	}
	bytedataval, _ := json.Marshal(vamodel)
	fmt.Println(string(bytedataval))
}

func TestStaticValidatorByMonthTask_getCommissionRate(t *testing.T) {
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT18:40:00", 2020, 4, 27), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT18:49:00", 2020, 4, 27), cstZone)
	ratemax := task.getCommissionRate(starttime, endtime, "fva186qhtc62cf6ejlt3erw6zk28mgw8ne7grhmyfn", "-date")
	fmt.Println(ratemax)
	ratemin := task.getCommissionRate(starttime, endtime, "fva186qhtc62cf6ejlt3erw6zk28mgw8ne7grhmyfn", "date")
	fmt.Println(ratemin)
}

func TestStaticValidatorByMonthTask_SetCaculateScope(t *testing.T) {
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT18:40:00", 2020, 4, 27), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT18:49:00", 2020, 4, 27), cstZone)
	fmt.Println(starttime)
	fmt.Println(endtime)
	task.SetCaculateScope(starttime, endtime)
	fmt.Println(task.startTime)
	fmt.Println(task.endTime)
}

func TestStaticValidatorByMonthTask_getValidatorsInPeriod(t *testing.T) {
	task := new(StaticValidatorByMonthTask)
	starttime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 5, 1), cstZone)
	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 5, 31), cstZone)
	data, err := task.getValidatorsInPeriod(starttime, endtime)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(len(data))
	fmt.Println(data)
}
