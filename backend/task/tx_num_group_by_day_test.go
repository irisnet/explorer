package task

import (
	"testing"
	"time"
)

func TestTxNumGroupByDayTask_DoTask(t *testing.T) {
	new(TxNumGroupByDayTask).DoTask()
	time.Sleep(time.Minute)
}
