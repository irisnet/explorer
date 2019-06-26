package lcd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
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
	url := fmt.Sprintf(UrlValidators, conf.Get().Hub.LcdUrl, page, size)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get Validators error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal Validators error", logger.String("err", err.Error()))
		return result
	}
	return result
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
	url := fmt.Sprintf(UrlDelegationsByDelegator, conf.Get().Hub.LcdUrl, delAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}

	if err := json.Unmarshal(resAsBytes, &delegations); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return
}

func GetDelegationsByValidatorAddr(valAddr string) (delegations []DelegationVo) {

	url := fmt.Sprintf(UrlDelegationsByValidator, conf.Get().Hub.LcdUrl, valAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}

	if err := json.Unmarshal(resAsBytes, &delegations); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return
}

func GetRedelegationsByValidatorAddr(valAddr string) (redelegations []ReDelegations) {

	url := fmt.Sprintf(UrlRedelegationsByValidator, conf.Get().Hub.LcdUrl, valAddr)
	resAsBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("get delegations by delegator adr from lcd error", logger.String("err", err.Error()), logger.String("URL", url))
		return
	}

	if err := json.Unmarshal(resAsBytes, &redelegations); err != nil {
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
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
		logger.Error("Unmarshal Delegations error", logger.String("err", err.Error()), logger.String("URL", url))
	}
	return
}

func GetRedelegationsByvalidatorAddr(valAddr string) {

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
	url := fmt.Sprintf(UrlStakePool, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal StakePool error", logger.String("err", err.Error()))
		return result
	}
	return
}

func SignInfo(consensusPubkey string) (result SignInfoVo) {
	url := fmt.Sprintf(UrlSignInfo, conf.Get().Hub.LcdUrl, consensusPubkey)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal SignInfoVo error", logger.String("err", err.Error()))
		return result
	}
	return
}
