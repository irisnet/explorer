package task

import (
	"github.com/irisnet/explorer/backend/utils"
	"testing"
	"time"
)

func TestStaticRewardsByDayTask_Start(t *testing.T) {
	new(StaticRewardsTask).Start()
}

func TestStaticRewardsByDayTask_getRewardsFromLcd(t *testing.T) {
	res1, res2, res3, err := new(StaticRewardsTask).getRewardsFromLcd("faa1ljemm0yznz58qxxs8xyak7fashcfxf5lssn6jm")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res1)))
	t.Log(string(utils.MarshalJsonIgnoreErr(res2)))
	t.Log(string(utils.MarshalJsonIgnoreErr(res3)))
}

func TestStaticRewardsByDayTask_getAllAccountRewards(t *testing.T) {

	res, err := new(StaticRewardsTask).getAllAccountRewards()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}

func TestStaticRewardsByDayTask_loadRewardsModel(t *testing.T) {
	res1, err := new(StaticRewardsTask).loadModelRewards("faa1ljemm0yznz58qxxs8xyak7fashcfxf5lssn6jm",
		utils.TruncateTime(time.Now().In(cstZone), utils.Day))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res1)))

}
func TestStaticRewardsByDayTask_loadRewards(t *testing.T) {
	res := new(StaticRewardsTask).loadRewards(utils.CoinsAsStr{
		{Amount: "18770397509925229288209"},
	})
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))

}
func TestStaticRewardsByDayTask_loadDelegationsRewardsDetail(t *testing.T) {

}
func TestStaticRewardsByDayTask_getAccountFromDb(t *testing.T) {
	res, err := new(StaticRewardsTask).getAccountFromDb()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
	//res1, err1 := new(StaticRewardsByDayTask).saveExStaticRewardsOps(res)
	//if err1 != nil {
	//	t.Fatal(err1.Error())
	//}
	//t.Log(string(utils.MarshalJsonIgnoreErr(res1)))
}

func TestStaticRewardsByDayTask_funcSubStr(t *testing.T) {
	new(StaticRewardsTask).DoTask()
}

func TestStaticRewardsTask_Common(t *testing.T) {
	today := utils.TruncateTime(time.Now(), utils.Day)
	t.Log(today.String())
	t.Log(today.Unix())
	t.Log(today.In(cstZone).String())
	t.Log(today.In(cstZone).Unix())
}
