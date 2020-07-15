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
	"strings"
	"sort"
	"math"
	"sync"
)

type AccountService struct {
	BaseService
}

func (service *AccountService) GetModule() Module {
	return Account
}

func (service *AccountService) Query(address string) (result vo.AccountVo) {
	var group sync.WaitGroup
	prefix, _, _ := utils.DecodeAndConvert(address)
	if prefix == conf.Get().Hub.Prefix.ValAddr {
		self, delegated := delegatorService.QueryDelegation(address)
		result.Amount = utils.Coins{self}
		result.Deposits = delegated

	} else {
		group.Add(1)
		go func() {
			defer group.Done()
			res, err := lcd.Account(address)
			if err == nil {
				decimalMap := make(map[string]int, len(res.Coins))
				var unit []string
				var amount utils.Coins
				for _, coinStr := range res.Coins {
					coin := utils.ParseCoin(coinStr)
					unit = append(unit, strings.Split(coin.Denom, types.AssetMinDenom)[0])
					amount = append(amount, coin)
				}

				assetkokens, err := document.AssetToken{}.GetAssetTokenDetailByTokenids(unit)
				if err == nil {
					for _, val := range assetkokens {
						decimalMap[val.TokenId] = val.Decimal
					}
				}

				for i, val := range amount {
					denom := strings.Split(val.Denom, types.AssetMinDenom)[0]
					if dem, ok := decimalMap[denom]; ok && dem > 0 {
						amount[i].Denom = denom
						amount[i].Amount = amount[i].Amount / float64(math.Pow10(dem))
					}

				}
				result.Amount = amount
			}
		}()

		//result.Deposits = delegatorService.GetDeposits(address)
		valaddress := utils.Convert(conf.Get().Hub.Prefix.ValAddr, address)
		validator, err := document.Validator{}.QueryValidatorDetailByOperatorAddr(valaddress)
		if err == nil {
			result.Status = getValidatorStatus(validator)
			result.Moniker = validator.Description.Moniker
			result.OperatorAddress = valaddress
		}
	}
	group.Add(1)

	go func() {
		defer group.Done()
		result.WithdrawAddress = lcd.QueryWithdrawAddr(address)
	}()

	group.Wait()


	result.IsProfiler = isProfiler(address)
	result.Address = address
	return result
}

func getValidatorStatus(validator document.Validator) string {
	if validator.Jailed == false {
		if validator.Status == types.Bonded {
			return "Active"
		} else {
			return "Candidate"
		}
	}
	return "Jailed"
}

func (service *AccountService) QueryRichList() (vo.AccountsInfoRespond) {

	result, err := document.Account{}.GetAccountList()

	if err != nil {
		logger.Error("GetAccountList have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	var accList []vo.AccountInfo
	var totalAmt = float64(0)
	//AccountAmtMap := make(map[string]float64,len(result))

	for _, acc := range result {
		//AccountAmtMap[acc.Address] = acc.Total.Amount
		//_,_,rewards,err := lcd.GetDistributionRewardsByValidatorAcc(acc.Address)
		//if err == nil && len(rewards) > 0 {
		//	if rewards[0].Denom == types.IRISAttoUint {
		//		rewardsAmt,_ := strconv.ParseFloat(rewards[0].Amount,64)
		//		AccountAmtMap[acc.Address] += rewardsAmt
		//		totalAmt += rewardsAmt
		//	}
		//
		//}
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
			UpdateAt: acc.UpdateAt,
		})
	}
	return accList
}

func isProfiler(address string) bool {
	//genesis := commonService.GetGenesis()
	//for _, profiler := range genesis.Result.Genesis.AppState.Guardian.Profilers {
	//	if profiler.Address == address {
	//		return true
	//	}
	//}
	if _, ok := types.ProfilerAddrList[address]; ok {
		return true
	}
	return false
}

func (service *AccountService) QueryDelegations(address string) (result vo.AccountDelegationsRespond) {
	delegations := lcd.GetDelegationsByDelAddr(address)
	var valaddrlist []string
	for _, val := range delegations {
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
			if valdator, ok := validatorMap[val.ValidatorAddr]; ok {
				data.Moniker = valdator.Description.Moniker
				data.Amount = computeVotingPower(valdator, val.Shares)
			}
		}
		result = append(result, &data)
	}
	sort.Sort(result)

	return result
}

func getValidators(valaddrlist []string) (validatorMap map[string]document.Validator) {

	valdators, err := document.Validator{}.QueryValidatorListByAddrList(valaddrlist)
	if err == nil {
		validatorMap = make(map[string]document.Validator, len(valdators))
		for _, val := range valdators {
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
		Denom:  types.IRISUint,
	}
}

func (service *AccountService) QueryUnbondingDelegations(address string) (result vo.AccountUnbondingDelegationsRespond) {

	unbondingdelegations := lcd.GetUnbondingDelegationsByDelegatorAddr(address)
	var valaddrlist []string
	for _, val := range unbondingdelegations {
		valaddrlist = append(valaddrlist, val.ValidatorAddr)
	}
	validatorMap := getValidators(valaddrlist)
	result = make(vo.AccountUnbondingDelegationsRespond, 0, len(unbondingdelegations))

	for _, val := range unbondingdelegations {
		data := vo.AccountUnbondingDelegationsVo{
			Address: val.ValidatorAddr,
			EndTime: val.MinTime,
			Height:  val.CreationHeight,
			Amount:  utils.ParseCoin(val.Balance),
		}
		if validatorMap != nil {
			if valdator, ok := validatorMap[val.ValidatorAddr]; ok {
				data.Moniker = valdator.Description.Moniker
			}
		}
		result = append(result, &data)

	}
	sort.Sort(result)
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
	var valaddrlist []string
	for _, val := range delegationrewards {
		valaddrlist = append(valaddrlist, val.Validator)
	}
	validatorMap := getValidators(valaddrlist)

	for _, val := range delegationrewards {
		data := vo.DelagationsRewards{Address: val.Validator, Amount: val.Reward}
		if validatorMap != nil {
			if valdator, ok := validatorMap[val.Validator]; ok {
				data.Moniker = valdator.Description.Moniker
			}
		}
		result.DelagationsRewards = append(result.DelagationsRewards, data)
	}

	return result
}
