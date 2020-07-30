package lcd_test

import (
	"testing"

	"github.com/irisnet/explorer/backend/lcd"
	"encoding/json"
)

func getModuleList() []string {

	return lcd.GovModuleList
}

func TestGetGovParamByModule(t *testing.T) {

	for k, v := range getModuleList() {
		t.Logf("k %d   ------- module name %s \n", k, v)
		resAsByte, err := lcd.GetGovModuleParam(v)
		if err != nil {
			t.Errorf("query module [%s] err, err: %v\n", v, err)
		}
		data,_ := json.Marshal(resAsByte)
		t.Logf("\nData:%v\n\n", string(data))
	}
}

func TestGetGovParamMapByModule(t *testing.T) {

	for k, v := range getModuleList() {
		t.Logf("k %d   ------- module name %s \n", k, v)
		resAsMap, err := lcd.GetGovModuleParamMap(v)
		if err != nil {
			t.Errorf("query module [%s] err, err: %v\n", v, err)
		}
		t.Logf("%v     alue type %T \n\n", resAsMap, resAsMap["value"])
	}
}

func printMap(m map[string]interface{}, t *testing.T) {
	for k, v := range m {
		t.Logf("k: %v  v: %v  type: %T\n", k, v, v)
	}
}

func TestGetGovParam(t *testing.T) {

	stakeMap, err := lcd.GetGovAssetParam()
	if err != nil {
		t.Errorf("GetGovAuthParam err: %v\n", err)
	}
	//t.Log(lcd.GovModuleStaking)
	data := stakeMap["params"].(map[string]interface{})["issue_token_base_fee"]
	t.Log(data.(map[string]interface{})["denom"])
	//printMap(stakeMap, t)
}

func TestGetGovGuardianParam(t *testing.T) {
	data,err := lcd.GetGovGuardianParam()
	if err != nil {
		t.Errorf("GetGovGuardianParam error:%v",err)
	}
	t.Log(data)
}
func TestGetGovAuthParam(t *testing.T) {

	t.Logf("GetGovAuthParam-------------")
	authMap, err := lcd.GetGovAuthParam()
	if err != nil {
		t.Errorf("GetGovAuthParam err: %v\n", err)
	}
	t.Log(lcd.GovModuleAuth)
	printMap(authMap, t)
}

func TestGetGovStakeParam(t *testing.T) {

	stakeMap, err := lcd.GetGovStakeParam()
	if err != nil {
		t.Errorf("GetGovAuthParam err: %v\n", err)
	}
	t.Log(lcd.GovModuleStaking)
	printMap(stakeMap, t)
}

func TestGetGovMintParam(t *testing.T) {

	mintMap, err := lcd.GetGovMintParam()
	if err != nil {
		t.Errorf("GetGovAuthParam err: %v\n", err)
	}
	t.Log(lcd.GovModuleMint)
	printMap(mintMap, t)
}

func TestGetGovDistrParam(t *testing.T) {

	distrMap, err := lcd.GetGovDistrParam()
	if err != nil {
		t.Errorf("GetGovAuthParam err: %v\n", err)
	}
	t.Log(lcd.GovModuleDistr)
	printMap(distrMap, t)
}

func TestGetGovSlashingParam(t *testing.T) {

	slashingMap, err := lcd.GetGovSlashingParam()
	if err != nil {
		t.Errorf("GetGovAuthParam err: %v\n", err)
	}
	t.Log(lcd.GovModuleSlashing)
	printMap(slashingMap, t)
}

func TestGetAllGovModuleParam(t *testing.T) {

	paramMap := lcd.GetAllGovModuleParam()

	t.Logf("total: %v \n", len(paramMap))
	printMap(paramMap, t)
}

func TestGetProposalVoters(t *testing.T) {
	data, err := lcd.GetProposalVoters(21)
	if err != nil {
		t.Errorf("query all gov porposal voter err: %v \n", err)
	}
	t.Log("len:", len(data))
	bytesmsg, _ := json.Marshal(data)
	t.Log(string(bytesmsg))
}
