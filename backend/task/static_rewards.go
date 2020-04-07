package task

import (
	"time"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"math"
	"math/big"
)

type StaticRewardsTask struct {
	account document.Account
}

func (task StaticRewardsTask) Name() string {
	return "static_rewards"
}
func (task StaticRewardsTask) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeTxNumByDay

	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		logger.Error(err.Error())
	}
}

func (task StaticRewardsTask) DoTask() error {
	ops, err := task.getAllAccountRewards()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return document.ExStaticRewards{}.Batch(ops)
}

func (task StaticRewardsTask) getRewardsFromLcd(address string) (utils.CoinsAsStr, []lcd.RewardsFromDelegations, utils.CoinsAsStr, error) {
	commissionrewards, delegationrewards, totalrewards, err := lcd.GetDistributionRewardsByValidatorAcc(address)
	if err != nil {
		logger.Error("GetDistributionRewardsByValidatorAcc have error", logger.String("err", err.Error()))
		return nil, nil, nil, err
	}
	return commissionrewards, delegationrewards, totalrewards, nil
}

func (task StaticRewardsTask) getAllAccountRewards() ([]txn.Op, error) {

	accAddrs, err := task.getAccountFromDb()
	if err != nil {
		return nil, err
	}

	return task.saveExStaticRewardsOps(accAddrs)
}

func (task StaticRewardsTask) saveExStaticRewardsOps(accAddrs []string) ([]txn.Op, error) {
	today := utils.TruncateTime(time.Now().In(cstZone), utils.Day)
	ops := make([]txn.Op, 0, len(accAddrs))
	now := time.Now().Unix()
	for _, addr := range accAddrs {
		item, err := task.loadModelRewards(addr, today)
		item.CreateAt = now
		if err != nil {
			continue
		}
		op := txn.Op{
			C:      document.CollectionNameExStaticRewards,
			Id:     bson.NewObjectId(),
			Insert: item,
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func (task StaticRewardsTask) loadModelRewards(addr string, today time.Time) (document.ExStaticRewards, error) {
	commisstionRds, _, totalRds, err := task.getRewardsFromLcd(addr)
	if err != nil {
		logger.Warn(err.Error())
		return document.ExStaticRewards{}, err
	}

	item := document.ExStaticRewards{
		Id:         bson.NewObjectId(),
		Address:    addr,
		Date:       today.In(cstZone),
		Total:      task.loadRewards(totalRds),
		Commission: task.loadRewards(commisstionRds),
		//DelegationsDetail: task.loadDelegationsRewardsDetail(delegationsRds),
	}
	if len(item.Total) > 0 {
		if len(item.Commission) == 0 {
			item.Delegations = item.Total
		} else {
			item.Delegations = task.loadDelegationsRewards(item.Total[0], item.Commission[0])
		}
	}
	return item, nil
}

func (task StaticRewardsTask) loadRewards(rewards utils.CoinsAsStr) []document.Rewards {

	var ret []document.Rewards
	for _, val := range rewards {
		item := document.Rewards{
			IrisAtto: val.Amount,
		}
		quoVal, err := utils.QuoByStr(val.Amount, new(big.Rat).SetFloat64(math.Pow10(18)).FloatString(0))
		if err == nil {
			value, _ := quoVal.Float64()
			item.Iris = value

		}
		ret = append(ret, item)
	}
	return ret
}

//func (task StaticRewardsByDayTask) loadDelegationsRewardsDetail(delegations []lcd.RewardsFromDelegations) []document.DelegationsRewards {
//	ret := make([]document.DelegationsRewards, 0, len(delegations))
//
//	for _, val := range delegations {
//		item := document.DelegationsRewards{
//			Validator: val.Validator,
//		}
//		if len(val.Reward) == 0 {
//			continue
//		}
//		retWards := make([]document.Rewards, 0, len(val.Reward))
//		for _, im := range val.Reward {
//			ite := document.Rewards{IrisAtto: im.Amount}
//			amount, _ := new(big.Rat).SetString(im.Amount)
//			value, _ := amount.Quo(amount, new(big.Rat).SetFloat64(math.Pow10(18))).Float64()
//			ite.Iris = value
//			retWards = append(retWards, ite)
//		}
//		item.Reward = retWards
//		ret = append(ret, item)
//	}
//	return ret
//}

func funcSubStr(amount1, amount2 string) *big.Rat {
	Amount1, ok1 := new(big.Rat).SetString(amount1)
	Amount2, ok2 := new(big.Rat).SetString(amount2)
	if !ok1 || !ok2 {
		logger.Error("big.Rat SetString failed in funcSubStr")
		return nil
	}
	value := new(big.Rat).Sub(Amount1, Amount2)
	return value
}

func (task StaticRewardsTask) loadDelegationsRewards(total, commission document.Rewards) []document.Rewards {
	subValue := funcSubStr(total.IrisAtto, commission.IrisAtto)
	if subValue == nil {
		return nil
	}

	var ret []document.Rewards
	subValueFloat64, ok := subValue.Float64()
	if ok {
		one := document.Rewards{
			Iris:     subValueFloat64,
			IrisAtto: subValue.String(),
		}
		ret = []document.Rewards{one}
	}

	return ret
}

func (task StaticRewardsTask) getAccountFromDb() ([]string, error) {
	size := 100
	offset := 0
	length := size
	var accAddress []string
	for {
		accounts, err := task.account.GetDelegatores(offset, size)
		if err != nil {
			logger.Error(err.Error())
			return accAddress, err
		}

		for _, val := range accounts {
			accAddress = append(accAddress, val.Address)
		}
		length = len(accounts)
		if length < size {
			break
		}
		offset += length
	}

	return accAddress, nil
}
