package task

import (
	"testing"
	"github.com/irisnet/explorer/backend/utils"
	"time"
)

func TestValidatorStaticByDayTask_Start(t *testing.T) {
	new(ValidatorStaticByDayTask).Start()
}

func TestValidatorStaticByDayTask_getValidatorFromDb(t *testing.T) {
	res, err := new(ValidatorStaticByDayTask).getValidatorFromDb()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}

func TestValidatorStaticByDayTask_loadValidatorTokens(t *testing.T) {
	validators, err := new(ValidatorStaticByDayTask).getValidatorFromDb()
	if err != nil {
		t.Fatal(err.Error())
	}
	res1, err := new(ValidatorStaticByDayTask).loadValidatorTokens(validators[0], utils.TruncateTime(time.Now(), utils.Day))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res1)))
}
