package task

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"math"
	"testing"
	"time"
)

func TestStaticRewardsByDayTask_Start(t *testing.T) {
	new(StaticDelegatorTask).Start()
}

func TestStaticRewardsByDayTask_getRewardsFromLcd(t *testing.T) {
	res1, res2, res3, err := new(StaticDelegatorTask).getRewardsFromLcd("faa1ljemm0yznz58qxxs8xyak7fashcfxf5lssn6jm")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res1)))
	t.Log(string(utils.MarshalJsonIgnoreErr(res2)))
	t.Log(string(utils.MarshalJsonIgnoreErr(res3)))
}

func TestStaticRewardsByDayTask_getAllAccountRewards(t *testing.T) {

	res, err := new(StaticDelegatorTask).getAllAccountRewards()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}

func TestStaticRewardsByDayTask_loadRewardsModel(t *testing.T) {
	res, _ := document.Account{}.GetAllAccount()
	res1, err := new(StaticDelegatorTask).loadModelRewards(res[0],
		utils.TruncateTime(time.Now().In(cstZone), utils.Day))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res1)))

}
func TestStaticRewardsByDayTask_loadRewards(t *testing.T) {
	res := new(StaticDelegatorTask).loadRewards(utils.CoinsAsStr{
		{Amount: "18770397509925229288209"},
	})
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))

}
func TestStaticRewardsByDayTask_loadDelegationsRewards(t *testing.T) {
	var total, commission document.Rewards
	total.IrisAtto = "182249450474538571"
	total.Iris = 0.182249450474539
	commission.IrisAtto = "167533025538184688"
	commission.Iris = 0.167533025538185
	res := new(StaticDelegatorTask).loadDelegationsRewards(total, commission)

	t.Log(res)
}
func TestStaticRewardsByDayTask_getAccountFromDb(t *testing.T) {
	res, err := new(StaticDelegatorTask).getAccountFromDb()
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

func TestStaticRewardsByDayTask_DoTask(t *testing.T) {
	new(StaticDelegatorTask).DoTask(HeartBeat)
}

func TestStaticRewardsTask_Common(t *testing.T) {
	today := utils.TruncateTime(time.Now(), utils.Day)
	t.Log(today.String())
	t.Log(today.Unix())
	t.Log(today.In(cstZone).String())
	t.Log(today.In(cstZone).Unix())
}

func TestStaticDelegatorTask_funcSubStr(t *testing.T) {
	amt1 := "5846551511526896055"
	amt2 := "5846551500047593555"

	v := funcSubStr(amt1, amt2)
	f, _ := v.Float64()
	t.Log(f / math.Pow10(18))
	t.Log(v.FloatString(2))
}
