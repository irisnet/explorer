package document

import (
	"testing"
)

func TestQueryAll(t *testing.T) {

	moduleParamList, err := GovParams{}.QueryAll()
	if err != nil {
		t.Error(err)
	}

	for k, v := range moduleParamList {

		t.Logf("k: %v v: %v \n", k, v)
	}

}
