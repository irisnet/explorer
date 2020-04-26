package task

import (
	"github.com/irisnet/explorer/backend/utils"
	"testing"
	"time"
)

func TestValidatorStaticByDayTask_Start(t *testing.T) {
	new(StaticValidatorTask).Start()
}

func TestValidatorStaticByDayTask_getValidatorFromDb(t *testing.T) {
	res, err := new(StaticValidatorTask).getValidatorFromDb()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}

func TestValidatorStaticByDayTask_loadValidatorTokens(t *testing.T) {
	validators, err := new(StaticValidatorTask).getValidatorFromDb()
	if err != nil {
		t.Fatal(err.Error())
	}
	res1, err := new(StaticValidatorTask).loadValidatorTokens(validators[0],
		utils.TruncateTime(time.Now().In(cstZone), utils.Day))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res1)))
}

func TestStaticValidatorTask_getAllValidatorTokens(t *testing.T) {
	txsop, err := new(StaticValidatorTask).getAllValidatorTokens()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(txsop)
}

func TestStaticValidatorTask_DoTask(t *testing.T) {
	new(StaticValidatorTask).DoTask()
}
