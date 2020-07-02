package document

import (
	"testing"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"time"
	"github.com/irisnet/explorer/backend/types"
	"encoding/json"
)

func TestExStaticValidator_GetCommissionRate(t *testing.T) {
	selector := bson.M{
		"commission.rate": 1,
	}
	year, month := 2020, 4
	datestr := fmt.Sprintf("%d-%02d-01T00:00:00", year, month)
	date, _ := time.ParseInLocation(types.TimeLayout, datestr, time.FixedZone("CST", 8*3600))
	datestr = fmt.Sprintf("%d-%02d-01T00:00:00", year, month+1)
	date1, _ := time.ParseInLocation(types.TimeLayout, datestr, time.FixedZone("CST", 8*3600))
	cond := bson.M{
		ExStaticValidatorMonthDateTag: bson.M{
			"$gte": date,
			"$lt":  date1,
		},
	}
	sorts := "-commission.rate"
	data, err := new(ExStaticValidator).GetCommissionRate(selector, cond, sorts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data.Commission.Rate)
	sorts = "commission.rate"
	data, err = new(ExStaticValidator).GetCommissionRate(selector, cond, sorts)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data.Commission.Rate)
}

func TestExStaticValidator_GetDataOneDay(t *testing.T) {
	address := "iva16tvg03gnrud9hzec7nsf9k8da74fqa9sk8xfqv"
	datestr := fmt.Sprintf("%d-%02d-01T00:00:00", 2020, 6)
	date, _ := time.ParseInLocation(types.TimeLayout, datestr, time.FixedZone("CST", 8*3600))
	data, err := new(ExStaticValidator).GetDataOneDay(date, address)
	if err != nil {
		t.Fatal(err)
	}
	bytestr, _ := json.Marshal(data)
	t.Log(string(bytestr))
}
