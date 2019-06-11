package document

import (
	"testing"
)

func TestQueryPowerChangeList(t *testing.T) {

	powerHistory, err := PowerChange{}.QueryPowerChangeList()
	if err != nil {
		t.Error(err)
	}

	for k, v := range powerHistory {
		t.Logf("k: %v  v: %v \n", k, v)
	}
}
