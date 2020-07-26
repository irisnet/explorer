package lcd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
	ctypes "github.com/irisnet/explorer/backend/types"
	"time"
)

func Validator(address string) (result ValidatorVo, err error) {
	validator, err := client.Staking().QueryValidator(address)
	if err != nil {
		return result, err
	}

	uptime, _ := time.Parse(time.RFC3339, validator.Commission.UpdateTime)
	unbondtime, _ := time.Parse(time.RFC3339, validator.UnbondingTime)
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
		uptime, _ := time.Parse(time.RFC3339, val.Commission.UpdateTime)
		unbondtime, _ := time.Parse(time.RFC3339, val.UnbondingTime)
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
		panic("improper use of BondStatusToString")
	}
}

func QueryWithdrawAddr(address string) (result string) {
	url := fmt.Sprintf(UrlWithdrawAddress, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}

	result = strings.Trim(string(resBytes), "\"")
	return
}

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

//func GetDelegationsFromValAddrByDelAddr(delAddr, valAddr string) (delegation DelegationFromVal) {
//	url := fmt.Sprintf(UrlDelegationsFromValidatorByDelegator, conf.Get().Hub.LcdUrl, delAddr, valAddr)
//	resAsBytes, err := utils.Get(url)
//	if err != nil {
//		logger.Error("get delegations from validator by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
//		return
//	}
//
//	if err := json.Unmarshal(resAsBytes, &delegation); err != nil {
//		logger.Error("Unmarshal DelegationsByDelAddr error", logger.String("err", err.Error()), logger.String("URL", url))
//	}
//	return
//}

func GetWithdrawAddressByValidatorAcc(validatorAcc string) (string, error) {

	url := fmt.Sprintf(UrlDistributionWithdrawAddressByValidatorAcc, conf.Get().Hub.LcdUrl, validatorAcc)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return "", err
	}

	var withdrawAddr string
	if err := json.Unmarshal(resAsBytes, &withdrawAddr); err != nil {
		logger.Error("Unmarshal WithdrawAddress error", logger.String("err", err.Error()), logger.String("URL", url))
	}

	return withdrawAddr, nil
}

func GetDistributionRewardsByValidatorAcc(validatorAcc string) (utils.CoinsAsStr, []RewardsFromDelegations, utils.CoinsAsStr, error) {

	if !strings.HasPrefix(validatorAcc, conf.Get().Hub.Prefix.AccAddr) {
		return nil, nil, nil, fmt.Errorf("address prefix is should %v", conf.Get().Hub.Prefix.AccAddr)
	}
	url := fmt.Sprintf(UrlDistributionRewardsByValidatorAcc, conf.Get().Hub.LcdUrl, validatorAcc)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return nil, nil, nil, err
	}

	var rewards DistributionRewards

	err = json.Unmarshal(resAsBytes, &rewards)
	if err != nil {
		return nil, nil, nil, err
	}

	return rewards.Commission, rewards.Delegations, rewards.Total, nil
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
		LooseTokens:  stakepool.LooseTokens,
		BondedTokens: stakepool.BondedTokens,
		//TotalSupply: ,
		//BondedRatio:,
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
