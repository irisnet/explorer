package task

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/types"
	"strconv"
)

type UpdateAccount struct {
	account document.Account
}

func (task UpdateAccount) Name() string {
	return "update_account"
}
func (task UpdateAccount) Start() {
	taskName := task.Name()
	timeInterval := conf.Get().Server.CronTimeAccountRewards

	utils.RunTimer(timeInterval, utils.Sec, func() {
		if err := tcService.runTask(taskName, timeInterval, task.DoTask); err != nil {
			logger.Error(err.Error())
		}
	})

}

func (task UpdateAccount) DoTask(fn func(string) chan bool) error {
	defer HeartQuit(fn(task.Name()))
	accounts, err := task.account.GetAllAccount()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	for i, val := range accounts {
		_, _, rewards, err := lcd.GetDistributionRewardsByValidatorAcc(val.Address)
		if err == nil && len(rewards) > 0 {
			newrewards := loadRewards(rewards)
			subvalue := newrewards.Amount - accounts[i].Rewards.Amount

			accounts[i].Total.Amount += subvalue

			accounts[i].Rewards = newrewards
		}
		if err := task.account.Update(accounts[i]); err != nil {
			logger.Warn("Account Update Rewards failed", logger.String("err", err.Error()))
		}
	}

	return nil
}

func loadRewards(coins utils.CoinsAsStr) utils.Coin {
	var retcoin utils.Coin
	for _, val := range coins {
		if val.Denom == types.IRISAttoUint {
			rewardsAmt, _ := strconv.ParseFloat(val.Amount, 64)
			return utils.Coin{Denom: val.Denom, Amount: rewardsAmt}
		}
	}
	return retcoin
}
