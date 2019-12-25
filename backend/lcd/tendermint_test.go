package lcd

import (
	"encoding/json"
	"testing"
)

func TestNodeInfo(t *testing.T) {
	if res, err := NodeInfo(""); err != nil {
		t.Fatal(err)
	} else {
		resBytes, _ := json.MarshalIndent(res, "", "\t")
		t.Log(string(resBytes))
	}
}

func TestNodeVersion(t *testing.T) {
	if res, err := NodeVersion(""); err != nil {
		t.Fatal(err)
	} else {
		resBytes, _ := json.MarshalIndent(res, "", "\t")
		t.Log(string(resBytes))
	}
}

//func TestGenesis(t *testing.T) {
//	if res, err := Genesis(); err != nil {
//		t.Fatal(err)
//	} else {
//		resBytes, _ := json.MarshalIndent(res, "", "\t")
//		t.Log(string(resBytes))
//	}
//}

func TestGetGenesisGovModuleParamMap(t *testing.T) {

	moduleParamMap, err := GetGenesisGovModuleParamMap()
	if err != nil {
		t.Error(err)
	}

	for k, v := range moduleParamMap {
		t.Logf("k: %v   v %v   %T \n\n\n", k, v, v)
	}
}

func TestGetGenesisAppStateGovParam(t *testing.T) {
	moduleParamMap, err := GetGenesisAppStateGovParam()
	if err != nil {
		t.Error(err)
	}

	for k, v := range moduleParamMap {
		t.Logf("k: %v   v %v   %T \n\n\n", k, v, v)
	}
}

func TestBlock(t *testing.T) {
	height := int64(1034)
	res := Block(height)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestBlockLatest(t *testing.T) {
	res := BlockLatest()
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestValidatorSet(t *testing.T) {
	height := int64(1034)
	res := ValidatorSet(height)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestLatestValidatorSet(t *testing.T) {
	res := LatestValidatorSet()
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}

func TestBlockResult(t *testing.T) {
	height := int64(1034)
	res := BlockResult(height)
	resBytes, _ := json.MarshalIndent(res, "", "\t")
	t.Log(string(resBytes))
}
