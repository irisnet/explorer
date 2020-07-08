package task

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/types"
	"strconv"
	"math/big"
	"time"
	"math"
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
		if err := updateAccount(&accounts[i]); err != nil {
			logger.Warn("get AccountInfo failed", logger.String("err", err.Error()))
		}
	}

	return nil
}

func updateAccount(account *document.Account) error {
	if res, err := updateAccountInfo(account); err == nil {
		account = res
	} else {
		logger.Warn("get AccountInfo failed", logger.String("err", err.Error()))
	}
	_, _, rewards, err := lcd.GetDistributionRewardsByValidatorAcc(account.Address)
	if err == nil && len(rewards) > 0 {
		newrewards := loadRewards(rewards)
		subvalue := newrewards.Amount - account.Rewards.Amount
		account.Total.Amount += subvalue
		account.Rewards = newrewards
	}
	if err := account.Update(*account); err != nil {
		logger.Warn("Account Update Rewards failed", logger.String("err", err.Error()))
		return err
	}
	return nil
}

func updateAccountInfo(account *document.Account) (*document.Account, error) {
	balance, res, err := getBalance(account)
	if err != nil {
		logger.Error("lcd getBalance Info have error", logger.String("err", err.Error()))
		return account, err
	} else {
		account = res
	}

	delegate, err := getDelegationInfo(account.Address)
	if err != nil {
		logger.Warn("update Delegation Info have error", logger.String("err", err.Error()))
	}
	unbondingdelegator, err := getUnbondingDelegationInfo(account.Address)
	if err != nil {
		logger.Warn("update UnbondingDelegation Info have error", logger.String("err", err.Error()))
	}
	starttime := time.Unix(account.TotalUpdateAt, 0).In(cstZone)
	if res, err := updateHeightTimeStamp(starttime, time.Now().In(cstZone), account); err == nil {
		account = res
	} else {
		logger.Warn("update HeightTimeStamp  have error", logger.String("err", err.Error()))
	}

	account.Delegation = utils.Coin{
		Denom:  balance.Denom,
		Amount: delegate,
	}
	account.UnbondingDelegation = utils.Coin{
		Denom:  balance.Denom,
		Amount: unbondingdelegator,
	}

	account.CoinIris = utils.Coin{
		Denom:  balance.Denom,
		Amount: balance.Amount,
	}

	account.Total = utils.Coin{
		Denom:  balance.Denom,
		Amount: balance.Amount + delegate + unbondingdelegator,
	}
	return account, nil
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

func getBalance(account *document.Account) (utils.Coin, *document.Account, error) {
	var balance utils.Coin
	ret, err := lcd.AccountInfo(account.Address)
	if err != nil {
		logger.Warn("lcd AccountInfo Info have error", logger.String("err", err.Error()))
		return balance, account, err
	}

	for _, val := range ret.Value.Coins {
		if val.Denom == types.IRISAttoUint {
			balance.Denom = val.Denom
			amount, _ := strconv.ParseFloat(val.Amount, 64)
			if amount > 0 {
				balance.Amount = amount
			}
			break
		}
	}
	if number, _ := strconv.ParseUint(ret.Value.AccountNum, 10, 64); number > 0 {
		account.AccountNumber = number
	}
	return balance, account, nil
}

func updateHeightTimeStamp(starttime, endtime time.Time, account *document.Account) (*document.Account, error) {
	var txModel document.CommonTx
	txs, err := txModel.GetTxsByDurationAddress(starttime, endtime, "")
	if err != nil {
		return nil, err
	}
	var height, timestamp int64
	for _, val := range txs {

		switch val.Type {
		case types.TxTypeStakeDelegate, types.TxTypeStakeBeginUnbonding, types.TxTypeBeginRedelegate,
			types.TxTypeStakeCreateValidator:
			if val.Height > height && val.From == account.Address {
				height = val.Height
				timestamp = val.Time.Unix()
			}
		}

	}
	if height > account.TotalUpdateHeight && timestamp > account.TotalUpdateAt {
		account.CoinIrisUpdateHeight = height
		account.CoinIrisUpdateAt = timestamp
		account.TotalUpdateHeight = height
		account.TotalUpdateAt = timestamp
	}
	return account, nil
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
	token = new(big.Rat).Mul(token, new(big.Rat).SetFloat64(math.Pow10(18)))
	return strconv.ParseFloat(token.FloatString(18), 64)
}

func getUnbondingDelegationInfo(address string) (float64, error) {

	unbondingDelegations := lcd.GetUnbondingDelegationsByDelegatorAddr(address)
	token, _ := new(big.Rat).SetString("0")
	if len(unbondingDelegations) > 0 {
		for _, v := range unbondingDelegations {
			coin := utils.ParseCoin(v.InitialBalance)
			token = new(big.Rat).Add(token, new(big.Rat).SetFloat64(coin.Amount))
		}
	}

	return strconv.ParseFloat(token.FloatString(18), 64)
}
