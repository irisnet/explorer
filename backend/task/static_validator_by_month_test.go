package task

import (
	"testing"
	"encoding/json"
)

var (
	task StaticValidatorByMonthTask
)

func TestStaticValidatorByMonthTask_getStaticValidator(t *testing.T) {
	terminaldate, err := task.staticModel.Terminaldate()
	if err != nil {
		t.Fatal(err.Error())
	}
	terminalData, err := task.staticModel.GetDataByDate(terminaldate)
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
