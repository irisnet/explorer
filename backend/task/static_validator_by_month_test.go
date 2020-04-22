package task

import (
	"testing"
	"encoding/json"
	"time"
	"github.com/irisnet/explorer/backend/types"
	"fmt"
)

var (
	task StaticValidatorByMonthTask
)

func TestStaticValidatorByMonthTask_getStaticValidator(t *testing.T) {

	endtime, _ := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d-%02d-%02dT00:00:00", 2020, 4, 22), cstZone)
	terminalData, err := task.staticModel.GetDataOneDay(endtime, "faa186qhtc62cf6ejlt3erw6zk28mgw8ne7gkx3t55")
	if err != nil {
		t.Fatal(err.Error())
	}
	mapData, err := task.getCreateValidatorTx()
	if err != nil {
		t.Fatal(err.Error())
	}

	res, err := new(StaticValidatorByMonthTask).getStaticValidator(terminalData, mapData)
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
