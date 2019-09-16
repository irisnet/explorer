package document

import (
	"encoding/json"
	"testing"
)

func TestQueryBlackList(t *testing.T) {

	blackListMap := BlackList{}.QueryBlackList()

	bytes, _ := json.Marshal(blackListMap)
	t.Log(string(bytes))
}
