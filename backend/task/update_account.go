package task

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/types"
	"math/big"
	"strings"
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
	stop := fn(task.Name())
	defer HeartQuit(stop)
	accounts, err := task.account.GetAllAccount()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	for i := range accounts {
		if !strings.HasPrefix(accounts[i].Address, conf.Get().Hub.Prefix.AccAddr) {
			continue
		}
		if err := updateAccount(&accounts[i]); err != nil {
			logger.Warn("get AccountInfo failed", logger.String("err", err.Error()))
		}
	}

	return nil
}

func updateAccount(account *document.Account) error {
	if res, err := getAccountInfo(account); err == nil {
		account = res
	} else {
		logger.Warn("get AccountInfo failed", logger.String("err", err.Error()))
	}

	if err := account.Update(*account); err != nil {
		logger.Warn("Account Update Rewards failed", logger.String("err", err.Error()))
		return err
	}
	return nil
}

func getAccountInfo(account *document.Account) (*document.Account, error) {
	balance, res, err := getBalance(account)
	if err != nil {
		return account, err
	}
	account = res

	delegation, err := getDelegationInfo(account.Address)
	if err != nil {
		logger.Warn("update Delegation Info have error", logger.String("err", err.Error()))
	}
	unbondingDelegation, err := getUnbondingDelegationInfo(account.Address)
	if err != nil {
		logger.Warn("update UnbondingDelegation Info have error", logger.String("err", err.Error()))
	}
	//rewards
	_, rewards, err := lcd.GetDistributionRewardsByValidatorAcc(account.Address)
	if err != nil {
		logger.Warn("update GetDistributionRewards Info have error", logger.String("err", err.Error()))
	}

	if len(rewards) > 0 {
		newrewards := loadRewards(rewards)
		account.Rewards = newrewards
	} else {
		account.Rewards = utils.Coin{Denom: account.Rewards.Denom}
	}

	account.Delegation = utils.Coin{
		Denom:  balance.Denom,
		Amount: delegation,
	}
	account.UnbondingDelegation = utils.Coin{
		Denom:  balance.Denom,
		Amount: unbondingDelegation,
	}

	account.CoinIris = utils.Coin{
		Denom:  balance.Denom,
		Amount: balance.Amount,
	}

	account.Total = utils.Coin{
		Denom:  balance.Denom,
		Amount: balance.Amount + delegation + unbondingDelegation + account.Rewards.Amount,
	}

	return account, nil
}

func loadRewards(coins utils.CoinsAsStr) utils.Coin {
	var retcoin utils.Coin
	for _, val := range coins {
		if val.Denom == types.StakeUint {
			rewardsAmt, _ := utils.ParseStringToFloat(val.Amount)
			return utils.Coin{Denom: val.Denom, Amount: rewardsAmt}
		}
	}
	return retcoin
}

func getBalance(account *document.Account) (utils.Coin, *document.Account, error) {
	var balance utils.Coin
	ret, err := lcd.AccountInfo(account.Address)
	if err != nil {
		return balance, account, err
	}

	coins, err := lcd.AccountBalances(account.Address)
	if err != nil {
		logger.Warn("lcd Account Balances Info have error", logger.String("err", err.Error()))
	}

	for _, val := range coins {
		if val.Denom == types.StakeUint {
			balance.Denom = val.Denom
			amount, _ := utils.ParseStringToFloat(val.Amount)
			if amount > 0 {
				balance.Amount = amount
			}
			break
		}
	}
	account.AccountNumber = ret.AccountNumber
	return balance, account, nil
}

func getDelegationInfo(address string) (float64, error) {
	delegations := lcd.GetDelegationsByDelAddr(address)
	token, _ := new(big.Rat).SetString("0")
	if len(delegations) > 0 {
		for _, v := range delegations {
			if validator, err := lcd.Validator(v.ValidatorAddr); err != nil {
				logger.Error("get validator fail", logger.String("validatorAddr", v.ValidatorAddr),
					logger.String("err", err.Error()))
				continue
			} else {
				value := utils.CovertShareTokens(validator.Tokens, validator.DelegatorShares, v.Shares)
				tmp, ok := new(big.Rat).SetString(value)
				if ok {
					token = new(big.Rat).Add(tmp, token)
				}

			}
		}
	}
	serveconf := conf.Get().Server
	if serveconf.NtScale > 0 {
		token = new(big.Rat).Mul(token, new(big.Rat).SetInt64(int64(serveconf.NtScale)))
	}
	return utils.ParseStringToFloat(token.FloatString(18))
}

func getUnbondingDelegationInfo(address string) (float64, error) {

	unbondingDelegations := lcd.GetUnbondingDelegationsByDelegatorAddr(address)
	token, _ := new(big.Rat).SetString("0")
	if len(unbondingDelegations) > 0 {
		for _, v := range unbondingDelegations {
			if v.InitialBalance != "" {
				balances, _ := new(big.Rat).SetString(v.InitialBalance)
				token = new(big.Rat).Add(token, balances)
			}

		}
	}

	return utils.ParseStringToFloat(token.FloatString(18))
}
