package document

import (
	"testing"
	"encoding/json"
)

func TestExStaticDelegatorMonth_GetLatest(t *testing.T) {
	res, err := new(ExStaticDelegatorMonth).GetLatest("faa12t4gfg502wra9lhtjjvqudq82rrzu2sk5j2l09")
	if err != nil {
		t.Fatal(err.Error())
	}

	bytedata, _ := json.Marshal(res)
	t.Log(string(bytedata))
}
