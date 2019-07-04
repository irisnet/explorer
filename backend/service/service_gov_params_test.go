package service

import (
	"testing"
)

func TestQueryAll(t *testing.T) {

	moduleParamList := new(GovParamsService).QueryAll()

	for k, v := range moduleParamList {
		t.Logf("idx: %v param: %v\n", k, v)
	}

}
