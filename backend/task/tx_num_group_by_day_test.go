package task

import (
	"github.com/robfig/cron/v3"
	"github.com/shopspring/decimal"
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

func TestTxNumGroupByDayTask_getTokenStat(t *testing.T) {
	task := TxNumGroupByDayTask{}
	task.DoTask(HeartBeat)
	//res := task.getTokenStat()
	//t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}

func TestTxNumGroupByDayTask_decimalShift(t *testing.T) {
	str := "314.15926"
	if d, err := decimal.NewFromString(str); err != nil {
		t.Fatal(err)
	} else {
		t.Log(str)
		t.Log(d.Shift(2))
		t.Log(d.Shift(-2))
	}

}
