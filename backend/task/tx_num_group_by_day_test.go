package task

import "testing"

func TestTxNumGroupByDayTask_Start(t *testing.T) {
	task := TxNumGroupByDayTask{}

	task.Start()
}

func TestTxNumGroupByDayTask_init(t *testing.T) {
	task := TxNumGroupByDayTask{}

	task.init()
}
