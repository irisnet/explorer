package task

import (
	"testing"
	"encoding/json"
	"time"
)

var (
	task StaticValidatorByMonthTask
)

func TestStaticValidatorByMonthTask_getStaticValidator(t *testing.T) {

	terminalData, err := task.staticModel.GetDataByDate(time.Now().In(cstZone))
	if err != nil {
		t.Fatal(err.Error())
	}
	mapData, err := task.getCreateValidatorTx()
	if err != nil {
		t.Fatal(err.Error())
	}

	res := new(StaticValidatorByMonthTask).getStaticValidator(terminalData[0], mapData)

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