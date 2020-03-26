package task

import "testing"

func Test_assetTaskShouldNotBeExecuted(t *testing.T) {
	service := TaskControlService{}

	taskName := "tn2"
	timeInterval := 10
	if res, err := service.assetTaskShouldNotBeExecuted(taskName, timeInterval); err != nil {
		t.Fatal(err)
	} else {
		t.Log(res)
	}
}
