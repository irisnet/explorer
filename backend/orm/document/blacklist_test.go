package document

import (
	"testing"
)

func TestQueryBlackList(t *testing.T) {

	blackListMap := BlackList{}.QueryBlackList()
	for k, v := range blackListMap {
		t.Logf("k: %v \nv: %v\n", k, v)
	}
}
