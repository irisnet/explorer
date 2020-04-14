package document

import (
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2"
	"testing"
	"encoding/json"
)

func TestTaskControl_QueryOneByTaskName(t *testing.T) {
	d := TaskControl{}

	taskName := "task_control"
	if res, err := d.QueryOneByTaskName(taskName); err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(utils.MarshalJsonIgnoreErr(res)))
	}
}

func TestTaskControl_Save(t *testing.T) {
	d := TaskControl{}

	record := TaskControl{
		TaskName:     "tn1",
		TimeInterval: 600,
		IsInProcess:  true,
	}

	if err := d.Save(record); err != nil {
		t.Fatal(err)
	} else {
		t.Log("success")
	}
}

func TestTaskControl_UpdateByPK(t *testing.T) {
	d := TaskControl{}

	taskName := "tn1"
	if task, err := d.QueryOneByTaskName(taskName); err != nil {
		if err != mgo.ErrNotFound {
			t.Fatal(err)
		} else {
			t.Log("can't find record")
			return
		}
	} else {
		task.IsInProcess = false
		task.TimeInterval = 901

		if err := d.UpdateByPK(task); err != nil {
			t.Fatal(err)
		} else {
			t.Log("success")

		}
	}
}

func TestTaskControl_LockTaskControl(t *testing.T) {
	d := TaskControl{}

	taskName := "tn1"
	instanceNo := "dsa"
	if err := d.LockTaskControl(taskName, instanceNo); err != nil {
		t.Fatal(err)
	} else {
		t.Log("success")
	}
}

func TestTaskControl_UnlockTaskControl(t *testing.T) {
	d := TaskControl{}
	taskName := "tn1"
	if err := d.UnlockTaskControl(taskName); err != nil {
		t.Fatal(err)
	} else {
		t.Log("success")
	}
}

func TestTaskControl_List(t *testing.T) {
	res, err := new(TaskControl).List(0, 100)
	if err != nil {
		t.Fatal(err.Error())
	}

	bytesdata, _ := json.Marshal(res)
	t.Log(string(bytesdata))
}
