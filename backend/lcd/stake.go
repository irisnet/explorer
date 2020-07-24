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
	url := fmt.Sprintf(UrlValidator, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	return result, nil
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

func GetDelegationsByDelAddr(delAddr string) (delegations []DelegationVo) {
	if !strings.HasPrefix(delAddr, conf.Get().Hub.Prefix.AccAddr) {
		return
	}
	url := fmt.Sprintf(UrlDelegationsByDelegator, conf.Get().Hub.LcdUrl, delAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}

	if err := json.Unmarshal(resAsBytes, &delegations); err != nil {
		logger.Error("Unmarshal DelegationsByDelAddr error", logger.String("err", err.Error()), logger.String("URL", url))
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

func GetDelegationsByValidatorAddr(valAddr string) (delegations []DelegationVo) {

	url := fmt.Sprintf(UrlDelegationsByValidator, conf.Get().Hub.LcdUrl, valAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}

	if err := json.Unmarshal(resAsBytes, &delegations); err != nil {
		logger.Error("Unmarshal DelegationsByValidatorAddr error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return
}

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

func GetJailedUntilAndMissedBlocksCountByConsensusPublicKey(publicKey string) (string, string, string, error) {

	url := fmt.Sprintf(UrlValidatorsSigningInfoByConsensuPublicKey, conf.Get().Hub.LcdUrl, publicKey)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return "", "", "", err
	}

	var valSign ValidatorSigningInfo

	err = json.Unmarshal(resAsBytes, &valSign)
	if err != nil {
		return "", "", "", err
	}

	return valSign.JailedUntil, valSign.MissedBlocksCount, valSign.StartHeight, nil
}

func GetRedelegationsByValidatorAddr(valAddr string) (redelegations []ReDelegations) {

	url := fmt.Sprintf(UrlRedelegationsByValidator, conf.Get().Hub.LcdUrl, valAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}

	if err := json.Unmarshal(resAsBytes, &redelegations); err != nil {
		logger.Error("Unmarshal Redelegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return
}

func GetUnbondingDelegationsByValidatorAddr(valAddr string) (unbondingDelegations []UnbondingDelegations) {

	url := fmt.Sprintf(UrlUnbondingDelegationByValidator, conf.Get().Hub.LcdUrl, valAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}

	if err := json.Unmarshal(resAsBytes, &unbondingDelegations); err != nil {
		logger.Error("Unmarshal UnbondingDelegationsByValidatorAddr error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return
}

func GetUnbondingDelegationsByDelegatorAddr(delAddr string) (unbondingDelegations []UnbondingDelegations) {
	if !strings.HasPrefix(delAddr, conf.Get().Hub.Prefix.AccAddr) {
		return
	}
	url := fmt.Sprintf(UrlUnbondingDelegationByDelegator, conf.Get().Hub.LcdUrl, delAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}

	if err := json.Unmarshal(resAsBytes, &unbondingDelegations); err != nil {
		logger.Error("Unmarshal UnbondingDelegationsByDelegatorAddr error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return
}

func DelegationByValidator(address string) (result []DelegationVo) {
	url := fmt.Sprintf(UrlDelegationByVal, conf.Get().Hub.LcdUrl, address)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal Delegation error", logger.String("err", err.Error()))
		return result
	}
	return
}

func StakePool() (result StakePoolVo) {
	stakepool, err := client.Staking().QueryPool()
	if err != nil {
		logger.Error("RPC Query Pool error", logger.String("err", err.Error()))
		return result
	}
	//data,_ := json.Marshal(stakepool)
	//fmt.Println(data)
	result = StakePoolVo{
		LooseTokens:  stakepool.LooseTokens,
		BondedTokens: stakepool.BondedTokens,
		//TotalSupply:,
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
