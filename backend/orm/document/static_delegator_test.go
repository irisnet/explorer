package document

import (
	"testing"
	"encoding/json"
	"time"
	"github.com/irisnet/explorer/backend/types"
)

func TestExStaticDelegator_GetDataByDate(t *testing.T) {
	//date := time.Date(2020,time.April,13,14,0,0,0,time.UTC)
	//date,_ := new(ExStaticDelegator).Terminaldate()
	date, _ := time.ParseInLocation(types.TimeLayout, "2020-04-04T00:00:00", time.FixedZone("CST", 8*3600))
	res, err := new(ExStaticDelegator).GetDataByDate(date)
	if err != nil {
		t.Fatal(err.Error())
	}

	bytesdate, _ := json.Marshal(res)
	t.Log(string(bytesdate))
}

//func TestExStaticDelegator_Terminaldate(t *testing.T) {
//	time, err := new(ExStaticDelegator).Terminaldate()
//	if err != nil {
//		t.Fatal(err.Error())
//	}
//	t.Log(time.String())
//}

func TestExStaticDelegator_GetDataOneDay(t *testing.T) {
	date, _ := time.ParseInLocation(types.TimeLayout, "2020-04-05T00:00:00", time.FixedZone("CST", 8*3600))
	res, err := new(ExStaticDelegator).GetDataOneDay(date, "faa12t4gfg502wra9lhtjjvqudq82rrzu2sk5j2l09")
	if err != nil {
		t.Fatal(err.Error())
	}

	bytesdate, _ := json.Marshal(res)
	t.Log(string(bytesdate))
}
