package service

import (
	"testing"
)

func TestUpdateCurrentValueByKey(t *testing.T) {

	kv := map[string]interface{}{}

	kv["censorship_jail_duration"] = "12345678900123"
	kv["downtime_jail_duration"] = "0987654321123"
	kv["34234324"] = "234123412342134"
	err := GovParamsService{}.UpdateCurrentValueByKey(kv)

	if err != nil {
		t.Error(err)
	}
	t.Log("update  ok---------")
}
