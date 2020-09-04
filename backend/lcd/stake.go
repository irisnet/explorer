package lcd

import (
	"fmt"
	"strings"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
	ctypes "github.com/irisnet/explorer/backend/types"
	"time"
	"math/big"
)

func Validator(address string) (result ValidatorVo, err error) {
	validator, err := client.Staking().QueryValidator(address)
	if err != nil {
		return result, err
	}
	//data,_ := json.Marshal(validator)
	//fmt.Println(string(data))
	if validator.OperatorAddress == "" {
		return result, fmt.Errorf("not found this validator %v", address)
	}

	uptime, _ := time.Parse(ctypes.TimeLayout1, validator.Commission.UpdateTime)
	unbondtime, _ := time.Parse(ctypes.TimeLayout1, validator.UnbondingTime)
	result = ValidatorVo{
		OperatorAddress: validator.OperatorAddress,
		ConsensusPubkey: validator.ConsensusPubkey,
		Jailed:          validator.Jailed,
		Status:          BondStatusToInt(validator.Status),
		Tokens:          validator.Tokens,
		DelegatorShares: validator.DelegatorShares,
		Description: Description{
			Moniker:  validator.Description.Moniker,
			Identity: validator.Description.Identity,
			Details:  validator.Description.Details,
			Website:  validator.Description.Website,
		},
		UnbondingHeight: fmt.Sprint(validator.UnbondingHeight),
		UnbondingTime:   unbondtime,
		Commission: Commission{
			Rate:          validator.Commission.Rate,
			MaxRate:       validator.Commission.MaxRate,
			MaxChangeRate: validator.Commission.MaxChangeRate,
			UpdateTime:    uptime,
		},
	}

	return
}

func Validators(page, size int) (result []ValidatorVo) {

	validators, err := client.Staking().QueryValidators(page, size)
	if err != nil {
		logger.Error("RPC Query Validators failed", logger.String("err", err.Error()))
		return result
	}

	for _, val := range validators {
		if tokenRat, ok := new(big.Rat).SetString(val.Tokens); ok {
			val.Tokens = utils.ConverToDisplayUint(ctypes.NtScale, tokenRat).FloatString(4)
		}
		uptime, _ := time.Parse(ctypes.TimeLayout1, val.Commission.UpdateTime)
		unbondtime, _ := time.Parse(ctypes.TimeLayout1, val.UnbondingTime)
		result = append(result, ValidatorVo{
			OperatorAddress: val.OperatorAddress,
			ConsensusPubkey: val.ConsensusPubkey,
			Jailed:          val.Jailed,
			Status:          BondStatusToInt(val.Status),
			Tokens:          val.Tokens,
			DelegatorShares: val.DelegatorShares,
			Description: Description{
				Moniker:  val.Description.Moniker,
				Identity: val.Description.Identity,
				Details:  val.Description.Details,
				Website:  val.Description.Website,
			},
			UnbondingHeight: fmt.Sprint(val.UnbondingHeight),
			UnbondingTime:   unbondtime,
			Commission: Commission{
				Rate:          val.Commission.Rate,
				MaxRate:       val.Commission.MaxRate,
				MaxChangeRate: val.Commission.MaxChangeRate,
				UpdateTime:    uptime,
			},
		})
	}
	return result
}

func BondStatusToInt(b string) int {
	switch b {
	case ctypes.TypeValStatusUnbonded:
		return 0
	case ctypes.TypeValStatusUnbonding:
		return 1
	case ctypes.TypeValStatusBonded:
		return 2
	default:
		return -1
	}
}

//func QueryWithdrawAddr(address string) (result string) {
//	url := fmt.Sprintf(UrlWithdrawAddress, conf.Get().Hub.LcdUrl, address)
//	resBytes, err := utils.Get(url)
//	if err != nil {
//		return result
//	}
//
//	result = strings.Trim(string(resBytes), "\"")
//	return
//}

func GetDelegationsByDelAddr(delAddr string) (result []DelegationVo) {
	if !strings.HasPrefix(delAddr, conf.Get().Hub.Prefix.AccAddr) {
		return
	}
	delegations, err := client.Staking().QueryDelegations(delAddr)
	if err != nil {
		logger.Error("get delegations by delegator adr from rpc error", logger.String("err", err.Error()))
		return
	}

	for _, val := range delegations {
		result = append(result, DelegationVo{
			DelegatorAddr: val.Delegation.DelegatorAddress,
			ValidatorAddr: val.Delegation.ValidatorAddress,
			Shares:        val.Delegation.Shares,
		})
	}
	return
}

func GetWithdrawAddressByAddress(validatorAcc string) (string, error) {

	withdrawAddr, err := client.Distr().QueryWithdrawAddr(validatorAcc)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()))
		return "", err
	}
	if withdrawAddr != "" {
		withdrawAddr = string([]byte(withdrawAddr)[1 : len(withdrawAddr)-1])
	}

	return withdrawAddr, nil
}

func GetDistributionRewardsByValidatorAcc(validatorAcc string) ([]RewardsFromDelegations, utils.CoinsAsStr, error) {

	rewards, err := client.Distr().QueryRewards(validatorAcc)
	if err != nil {
		logger.Error("get delegations by delegator adr from rpc error", logger.String("err", err.Error()))
		return nil, nil, err
	}

	var (
		total       utils.CoinsAsStr
		delegations []RewardsFromDelegations
	)

	for _, val := range rewards.Total {
		total = append(total, utils.CoinAsStr{Denom: val.Denom, Amount: val.Amount.String()})
	}
	for _, val := range rewards.Rewards {
		item := RewardsFromDelegations{Validator: val.Validator}
		for _, one := range val.Reward {
			item.Reward = append(item.Reward, utils.CoinAsStr{Denom: one.Denom, Amount: one.Amount.String()})
		}
		delegations = append(delegations, item)
	}

	return delegations, total, nil
}

func GetDistributionCommissionRewardsByAddress(validatorAcc string) (utils.CoinsAsStr, error) {

	var (
		commission utils.CoinsAsStr
	)
	if !strings.HasPrefix(validatorAcc, conf.Get().Hub.Prefix.ValAddr) {
		return nil, nil
	}

	commissionData, err := client.Distr().QueryCommission(validatorAcc)
	if err != nil {
		logger.Error("get commisssions by delegator adr from rpc error", logger.String("err", err.Error()))
		return nil, err
	}
	for _, val := range commissionData.Commission {
		commission = append(commission, utils.CoinAsStr{Denom: val.Denom, Amount: val.Amount.String()})
	}

	return commission, nil
}
func GetJailedUntilAndMissedBlocksCountByConsensusPublicKey(publicKey string) (string, string, int64, error) {

	valSign := SignInfo(publicKey)

	return valSign.JailedUntil.String(), fmt.Sprint(valSign.MissedBlocksCounter), valSign.StartHeight, nil
}

func GetRedelegationsByValidatorAddr(valAddr string) (result []ReDelegation) {

	reDelegations, err := client.Staking().QueryRedelegationsFrom(valAddr)
	if err != nil {
		logger.Error("get redelegations by validator adr from rpc error", logger.String("err", err.Error()))
		return
	}
	for _, val := range reDelegations {
		reDelegation := ReDelegation{
			DelegatorAddr:    val.Redelegation.DelegatorAddress,
			ValidatorDstAddr: val.Redelegation.ValidatorDstAddress,
			ValidatorSrcAddr: val.Redelegation.ValidatorSrcAddress,
		}
		if len(val.Entries) > 0 {
			reDelegation.Balance = val.Entries[0].Balance.String()
			reDelegation.InitialBalance = val.Entries[0].RedelegationEntry.InitialBalance.String()
			reDelegation.CreationHeight = fmt.Sprint(val.Entries[0].RedelegationEntry.CreationHeight)
			reDelegation.MinTime = val.Entries[0].RedelegationEntry.CompletionTime.Unix()
			reDelegation.SharesDst = val.Entries[0].RedelegationEntry.SharesDst.String()

		}
		result = append(result, reDelegation)
	}
	return
}

func GetUnbondingDelegationsByValidatorAddr(valAddr string) (result []UnbondingDelegation) {

	unbondingDelegations, err := client.Staking().QueryUnbondingDelegationsFrom(valAddr)
	if err != nil {
		logger.Error("get unbonding delegations by delegator adr from rpc error", logger.String("err", err.Error()))
		return
	}
	for _, val := range unbondingDelegations {
		data := UnbondingDelegation{
			DelegatorAddr: val.DelegatorAddress,
			ValidatorAddr: val.ValidatorAddress,
		}
		if len(val.Entries) > 0 {
			data.MinTime = val.Entries[0].CompletionTime.String()
			data.InitialBalance = val.Entries[0].InitialBalance.String()
			data.CreationHeight = val.Entries[0].CreationHeight
			data.Balance = val.Entries[0].Balance.String()
		}
		result = append(result, data)
	}
	return
}

func GetUnbondingDelegationsByDelegatorAddr(delAddr string) (result []UnbondingDelegation) {
	if !strings.HasPrefix(delAddr, conf.Get().Hub.Prefix.AccAddr) {
		return
	}
	unbondingDelegations, err := client.Staking().QueryUnbondingDelegations(delAddr)
	if err != nil {
		logger.Error("get delegations by delegator adr from rpc error", logger.String("err", err.Error()))
		return
	}

	for _, val := range unbondingDelegations {
		data := UnbondingDelegation{
			DelegatorAddr: val.DelegatorAddress,
			ValidatorAddr: val.ValidatorAddress,
		}
		for _, item := range val.Entries {
			data.MinTime = item.CompletionTime.String()
			data.InitialBalance = item.InitialBalance.String()
			data.CreationHeight = item.CreationHeight
			data.Balance = item.Balance.String()
			result = append(result, data)
		}
		//if len(val.Entries) > 0 {
		//	data.MinTime = val.Entries[0].CompletionTime.String()
		//	data.InitialBalance = val.Entries[0].InitialBalance.String()
		//	data.CreationHeight = val.Entries[0].CreationHeight
		//	data.Balance = val.Entries[0].Balance.String()
		//}

	}
	return
}

func GetDelegationByValidator(address string) (result []DelegationVo) {
	delegations, err := client.Staking().QueryDelegationsTo(address)
	if err != nil {
		logger.Error("Query Delegation Rpc error", logger.String("err", err.Error()))
		return
	}
	for _, val := range delegations {
		result = append(result, DelegationVo{
			DelegatorAddr: val.Delegation.DelegatorAddress,
			ValidatorAddr: val.Delegation.ValidatorAddress,
			Shares:        val.Delegation.Shares,
		})
	}
	return
}

func StakePool() (result StakePoolVo) {
	stakepool, err := client.Staking().QueryPool()
	if err != nil {
		logger.Error("RPC Query Pool error", logger.String("err", err.Error()))
		return result
	}
	result = StakePoolVo{
		LooseTokens:  stakepool.NotBondedTokens,
		BondedTokens: stakepool.BondedTokens,
	}
	return
}

func SignInfo(consensusPubkey string) (result SignInfoVo) {

	signinfo, err := client.Slashing().QueryValidatorSigningInfo(consensusPubkey)
	if err != nil {
		logger.Error("RPC Query Validator Signing Info error", logger.String("err", err.Error()))
		return result
	}
	result = SignInfoVo{
		StartHeight:         signinfo.StartHeight,
		IndexOffset:         signinfo.IndexOffset,
		JailedUntil:         signinfo.JailedUntil,
		MissedBlocksCounter: signinfo.MissedBlocksCounter,
	}
	return
}
