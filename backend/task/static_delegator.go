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

type StaticDelegatorTask struct {
	account document.Account
}

func (task StaticDelegatorTask) Name() string {
	return "static_delegator"
}
func (task StaticDelegatorTask) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeStaticDelegator

	if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
		logger.Error(err.Error())
	}
}

func (task StaticDelegatorTask) DoTask() error {
	ops, err := task.getAllAccountRewards()
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return document.ExStaticDelegator{}.Batch(ops)
}

func (task StaticDelegatorTask) getRewardsFromLcd(address string) (utils.CoinsAsStr, []lcd.RewardsFromDelegations, utils.CoinsAsStr, error) {
	commissionrewards, delegationrewards, totalrewards, err := lcd.GetDistributionRewardsByValidatorAcc(address)
	if err != nil {
		logger.Error("GetDistributionRewardsByValidatorAcc have error", logger.String("err", err.Error()))
		return nil, nil, nil, err
	}
	return commissionrewards, delegationrewards, totalrewards, nil
}

func (task StaticDelegatorTask) getAllAccountRewards() ([]txn.Op, error) {

	accounts, err := task.getAccountFromDb()
	if err != nil {
		return nil, err
	}

	return task.saveExStaticRewardsOps(accounts)
}

func (task StaticDelegatorTask) saveExStaticRewardsOps(accs []document.Account) ([]txn.Op, error) {
	today := utils.TruncateTime(time.Now().In(cstZone), utils.Day)
	ops := make([]txn.Op, 0, len(accs))
	now := time.Now().Unix()
	if conf.Get().Server.CaculateDebug {
		today = time.Now().In(cstZone)
	}
	for _, acc := range accs {
		item, err := task.loadModelRewards(acc, today)
		item.CreateAt = now
		if err != nil {
			continue
		}
		op := txn.Op{
			C:      document.CollectionNameExStaticDelegator,
			Id:     bson.NewObjectId(),
			Insert: item,
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func (task StaticDelegatorTask) loadModelRewards(account document.Account, today time.Time) (document.ExStaticDelegator, error) {
	commisstionRds, _, totalRds, err := task.getRewardsFromLcd(account.Address)
	if err != nil {
		logger.Warn(err.Error())
		return document.ExStaticDelegator{}, err
	}

	item := document.ExStaticDelegator{
		Id:         bson.NewObjectId(),
		Address:    account.Address,
		Date:       today.In(cstZone),
		Total:      task.loadRewards(totalRds),
		Commission: task.loadRewards(commisstionRds),
		Delegation: account.Delegation,
		//DelegationsDetail: task.loadDelegationsRewardsDetail(delegationsRds),
	}
	if len(item.Total) > 0 {
		if len(item.Commission) == 0 {
			item.DelegationsRewards = item.Total
		} else {
			item.DelegationsRewards = task.loadDelegationsRewards(item.Total[0], item.Commission[0])
		}
	}
	return item, nil
}

func (task StaticDelegatorTask) loadRewards(rewards utils.CoinsAsStr) []document.Rewards {

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

func (task StaticDelegatorTask) loadDelegationsRewards(total, commission document.Rewards) []document.Rewards {
	subValue := funcSubStr(total.IrisAtto, commission.IrisAtto)
	if subValue == nil {
		return nil
	}
	subValueFloat64, _ := subValue.Float64()
	one := document.Rewards{
		Iris:     subValueFloat64 / math.Pow10(18),
		IrisAtto: subValue.FloatString(0),
	}

	return []document.Rewards{one}
}

func (task StaticDelegatorTask) getAccountFromDb() ([]document.Account, error) {
	size := 100
	offset := 0
	length := size
	var ret []document.Account
	for {
		accounts, err := task.account.GetDelegatores(offset, size)
		if err != nil {
			logger.Error(err.Error())
			return accounts, err
		}
		ret = append(ret, accounts...)
		length = len(accounts)
		if length < size {
			break
		}
		offset += length
	}

	return ret, nil
}
