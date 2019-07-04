package document

import (
	"testing"
	"time"
)

func TestCount(t *testing.T) {

	total, err := TxNumStat{}.Count()
	if err != nil {
		t.Error(err)
	}

	t.Logf("total: %v \n", total)
}

func TestQueryByDuration(t *testing.T) {
	txNumGroup, err := TxNumStat{}.QueryByDuration(time.Now().Add(time.Hour*-24), time.Now())

	if err != nil {
		t.Error(err)
	}

	for k, v := range txNumGroup {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}
