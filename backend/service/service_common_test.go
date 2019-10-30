package service

import (
	"testing"
)

//func TestGetGenesis(t *testing.T) {
//
//	genesis := CommonService{}.GetGenesis()
//	t.Logf("genesis: %v \n", genesis)
//}

func TestGetConfig(t *testing.T) {

	configList := CommonService{}.GetConfig()

	for k, config := range configList {
		t.Logf("idx: %v  config: %v \n", k, config)
	}
}
