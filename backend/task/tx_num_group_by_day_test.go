package task

import (
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func TestTxNumGroupByDayTask_Start(t *testing.T) {
	task := TxNumGroupByDayTask{}

	c := cron.New()
	c.AddFunc("0/1 * * * *", func() {
		task.Start()
	})
	c.Start()
	time.Sleep(time.Duration(1) * time.Hour)
}

func TestTxNumGroupByDayTask_init(t *testing.T) {
	task := TxNumGroupByDayTask{}

	task.init()
}
