package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"strconv"
	"math/big"
	"github.com/irisnet/irishub-sync/util/constant"
)

type AccountService struct {
	BaseService
}

func (service *AccountService) GetModule() Module {
	return Account
}

func (service *AccountService) Query(address string) (result vo.AccountVo) {
	prefix, _, _ := utils.DecodeAndConvert(address)
	if prefix == conf.Get().Hub.Prefix.ValAddr {
		self, delegated := delegatorService.QueryDelegation(address)
		result.Amount = utils.Coins{self}
		result.Deposits = delegated

	} else {
		res, err := lcd.Account(address)
		if err == nil {
			var amount utils.Coins
			for _, coinStr := range res.Coins {
				coin := utils.ParseCoin(coinStr)
				amount = append(amount, coin)
			}
			result.Amount = amount
		}
		//result.Deposits = delegatorService.GetDeposits(address)
		valaddress := utils.Convert(conf.Get().Hub.Prefix.ValAddr, address)
		validator, err := document.Validator{}.QueryValidatorDetailByOperatorAddr(valaddress)
		if err == nil {
			result.Status = strconv.Itoa(validator.Status)
			result.Moniker = validator.Description.Moniker
			result.OperatorAddress = valaddress
		}
	}


	result.WithdrawAddress = lcd.QueryWithdrawAddr(address)
	result.IsProfiler = isProfiler(address)
	result.Address = address
	return result
}

func (service *AccountService) QueryRichList() interface{} {

	result, err := document.Account{}.GetAccountList()

	if err != nil {
		logger.Error("GetAccountList have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	var accList []vo.AccountInfo
	var totalAmt = float64(0)

	for _, acc := range result {
		totalAmt += acc.Total.Amount
	}

	for index, acc := range result {
		rate, _ := utils.NewRatFromFloat64(acc.Total.Amount / totalAmt).Float64()
		accList = append(accList, vo.AccountInfo{
			Rank:    index + 1,
			Address: acc.Address,
			Balance: utils.Coins{
				acc.Total,
			},
			Percent:  rate,
			UpdateAt: acc.TotalUpdateAt,
		})
	}
	return accList
}

func isProfiler(address string) bool {
	genesis := commonService.GetGenesis()
	for _, profiler := range genesis.Result.Genesis.AppState.Guardian.Profilers {
		if profiler.Address == address {
			return true
		}
	}
	return false
}

func (service *AccountService) QueryDelegations(address string) (result []*vo.AccountDelegationsVo) {
	delegations := lcd.GetDelegationsByDelAddr(address)
	var  valaddrlist  []string
	for _,val := range delegations {
		valaddrlist = append(valaddrlist, val.ValidatorAddr)
	}
	validatorMap := getValidators(valaddrlist)

	result = make([]*vo.AccountDelegationsVo, 0, len(delegations))
	for _, val := range delegations {
		data := vo.AccountDelegationsVo{
			Address: val.ValidatorAddr,
			Shares:  val.Shares,
			Height:  val.Height,
		}
		if validatorMap != nil {
			if valdator,ok := validatorMap[val.ValidatorAddr];ok {
				data.Moniker = valdator.Description.Moniker
				data.Amount = computeVotingPower(valdator, val.Shares)
			}
		}
		result = append(result, &data)
	}

	return result
}

func getValidators(valaddrlist  []string)(validatorMap map[string]document.Validator) {

	valdators, err := document.Validator{}.QueryValidatorListByAddrList(valaddrlist)
	if err == nil {
		validatorMap = make(map[string]document.Validator,len(valdators))
		for _,val := range valdators {
			validatorMap[val.OperatorAddress] = val
		}
	}

	return validatorMap
}

func computeVotingPower(validator document.Validator, shares string) utils.Coin {
	rate, err := utils.QuoByStr(validator.Tokens, validator.DelegatorShares)
	if err != nil {
		logger.Error("validator.Tokens / validator.DelegatorShares", logger.String("err", err.Error()))
		rate, _ = new(big.Rat).SetString("1")
	}
	sharesAsRat, ok := new(big.Rat).SetString(shares)
	if !ok {
		logger.Error("convert validator.Tokens type (string to big.Rat) ", logger.Any("result", ok),
			logger.String("validator tokens", validator.Tokens))
	}

	tokensAsRat := new(big.Rat)
	tokensAsRat.Mul(rate, sharesAsRat)
	amount, _ := strconv.ParseFloat(tokensAsRat.FloatString(4), 64)

	return utils.Coin{
		Amount: amount,
		Denom:  constant.IrisAttoUnit,
	}
}

func (service *AccountService) QueryUnbondingDelegations(address string) (result []*vo.AccountUnbondingDelegationsVo) {

	unbondingdelegations := lcd.GetUnbondingDelegationsByDelegatorAddr(address)
	var  valaddrlist  []string
	for _,val := range unbondingdelegations {
		valaddrlist = append(valaddrlist, val.ValidatorAddr)
	}
	validatorMap := getValidators(valaddrlist)

	for _, val := range unbondingdelegations {
		data := vo.AccountUnbondingDelegationsVo{
			Address: val.ValidatorAddr,
			EndTime: val.MinTime,
			Height:  val.CreationHeight,
			Amount:  utils.ParseCoin(val.Balance),
		}
		if validatorMap != nil {
			if valdator,ok := validatorMap[val.ValidatorAddr];ok {
				data.Moniker = valdator.Description.Moniker
			}
		}

	}
	return result
}

func (service *AccountService) QueryRewards(address string) (result vo.AccountRewardsVo) {

	commissionrewards, delegationrewards, rewards, err := lcd.GetDistributionRewardsByValidatorAcc(address)
	if err != nil {
		logger.Error("GetDistributionRewardsByValidatorAcc have error", logger.String("err", err.Error()))
		return
	}

	result.CommissionRewards = commissionrewards
	result.TotalRewards = rewards
	var  valaddrlist  []string
	for _,val := range delegationrewards {
		valaddrlist = append(valaddrlist, val.Validator)
	}
	validatorMap := getValidators(valaddrlist)

	for _, val := range delegationrewards {
		data := vo.DelagationsRewards{Address: val.Validator, Amount: val.Reward}
		if validatorMap != nil {
			if valdator,ok := validatorMap[val.Validator];ok {
				data.Moniker = valdator.Description.Moniker
			}
		}
		result.DelagationsRewards = append(result.DelagationsRewards, data)
	}

	return result
}
