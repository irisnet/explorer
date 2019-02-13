package task

import (
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	UpTimeChangeTask{}.Start()
	time.Sleep(1 * time.Hour)
}

func TestStartTask(t *testing.T) {
	TxNumStatTask{}.Start()
	time.Sleep(1 * time.Hour)
}
