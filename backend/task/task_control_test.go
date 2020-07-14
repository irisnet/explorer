package task

import "testing"

func Test_assetTaskShouldNotBeExecuted(t *testing.T) {
	service := TaskControlService{}

	taskName := "tn2"
	timeInterval := 10
	if res, err := service.assetTaskShouldNotBeExecuted(taskName, timeInterval, false); err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
}

func TestTaskControlService_runTask(t *testing.T) {
	task := TxNumGroupByDayTask{}
	taskName := task.Name()
	timeInterval := 1

	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		t.Log(err)
	} else {
		t.Log("success")
	}
}

func TestTaskControlService_Start(t *testing.T) {
	new(TaskControlMonitor).Start()
}

func TestTaskControlService_DoTask(t *testing.T) {
	new(TaskControlMonitor).DoTask(HeartBeat)
}

//func TestTaskControlMonitor_unlockAllTasks(t *testing.T) {
//	new(TaskControlMonitor).unlockAllTasks()
//}
