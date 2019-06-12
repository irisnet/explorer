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

func TestUpdateCurrentModuleParamValue(t *testing.T) {

	kv := map[string]interface{}{}

	kv["censorship_jail_duration"] = "12345678900123"
	kv["downtime_jail_duration"] = "0987654321123"
	kv["34234324"] = "234123412342134"
	err := GovParams{}.UpdateCurrentModuleParamValue(kv)

	if err != nil {
		t.Error(err)
	}
	t.Log("update  ok---------")
}
