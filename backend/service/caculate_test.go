package service

import (
	"testing"
	"time"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

var cstZone = time.FixedZone("CST", 8*3600)

func TestCaculateService_GetDelegatorCaculateMonth(t *testing.T) {
	//dbstarttime, err := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d.%02d.%02dT00:00:00", 2020, 4, 10), cstZone)
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//dbendtime, err := time.ParseInLocation(types.TimeLayout, fmt.Sprintf("%d.%02d.%02dT00:00:00", 2020, 5, 1), cstZone)
	//if err != nil {
	//	t.Fatal(err.Error())
	//}

	cond := bson.M{}
	cond["date"] = bson.M{
		"$gte": fmt.Sprintf("%d.%02d.%02d", 2020, 4, 10),
		"$lt":  fmt.Sprintf("%d.%02d.%02d", 2020, 5, 1),
	}
	res, _, err := new(CaculateService).GetDelegatorCaculateMonth(cond, 1, 100, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(res)

}

func TestCaculateService_GetValidatorCaculateMonth(t *testing.T) {

	cond := bson.M{}
	cond["date"] = bson.M{
		"$gte": fmt.Sprintf("%d.%02d.%02d", 2020, 4, 10),
		"$lt":  fmt.Sprintf("%d.%02d.%02d", 2020, 5, 1),
	}
	res, _, err := new(CaculateService).GetValidatorCaculateMonth(cond, 1, 100, false)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(res)
}
