package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/irisnet/explorer/backend/conf"

	"github.com/irisnet/explorer/backend/utils"
)

const (
	//module
	GovModuleAuth     = "auth"
	GovModuleStake    = "stake"
	GovModuleMint     = "mint"
	GovModuleDistr    = "distr"
	GovModuleSlashing = "slashing"
	//auth key
	GovModuleAuthGasPriceThreshold string = "gas_price_threshold"
	GovModuleAuthTxSize            string = "tx_size"
	//stake key
	GovModuleStakeUnbondingTime = "unbonding_time"
	GovModuleStakeMaxValidators = "max_validators"
	//mint key
	GovModuleMintInflation = "inflation"
	//distr key
	GovModuleDistrCommunityTax        = "community_tax"
	GovModuleDistrBaseProposerReward  = "base_proposer_reward"
	GovModuleDistrBonusProposerReward = "bonus_proposer_reward"

	//slashing key
	GovModuleSlashingSlashFractionCensorship = "slash_fraction_censorship"
	GovModuleSlashingMaxEvidenceAge          = "max_evidence_age"
	GovModuleSlashingSignedBlocksWindow      = "signed_blocks_window"
	GovModuleSlashingDoubleSignJailDuration  = "double_sign_jail_duration"
	GovModuleSlashingMinSignedPerWindow      = "min_signed_per_window"
	GovModuleSlashingCensorshipJailDuration  = "censorship_jail_duration"
	GovModuleSlashingSlashFractionDoubleSign = "slash_fraction_double_sign"
	GovModuleSlashingSlashFractionDowntime   = "slash_fraction_downtime"
	GovModuleSlashingDowntimJailDuration     = "downtime_jail_duration"
)

var GovModuleList = []string{GovModuleAuth, GovModuleStake, GovModuleMint, GovModuleDistr, GovModuleSlashing}

type RangeDescription struct {
	Range       string
	Description string
}

func GetAuthKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleAuthGasPriceThreshold] = RangeDescription{Range: "0,1000000000000000000", Description: "the minimum of gas price"}
	result[GovModuleAuthTxSize] = RangeDescription{Range: "500,1500", Description: "transaction size limit"}
	return result
}

func GetStakeKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleStakeUnbondingTime] = RangeDescription{Range: "1209600000000000,", Description: "unbonding time"}
	result[GovModuleStakeMaxValidators] = RangeDescription{Range: "100,200", Description: "the maximum number of validators"}
	return result
}

func GetDistrKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleDistrCommunityTax] = RangeDescription{Range: "0,0.2", Description: "community tax"}
	result[GovModuleDistrBonusProposerReward] = RangeDescription{Range: "0,0.02", Description: "maximum additional ratio bonus ratio"}
	result[GovModuleDistrBaseProposerReward] = RangeDescription{Range: "0,0.08", Description: "base radio of the block reward"}
	return result
}

func GetMintKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleMintInflation] = RangeDescription{Range: "0,0.2", Description: "inflation rate"}
	return result
}

func GetSlashingKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}

	result[GovModuleSlashingSlashFractionCensorship] = RangeDescription{Range: "0,0.1", Description: "slash fraction censorship"}
	result[GovModuleSlashingMaxEvidenceAge] = RangeDescription{Range: "34560,", Description: "max evidence age"}
	result[GovModuleSlashingSignedBlocksWindow] = RangeDescription{Range: "100,140000", Description: "signed blocks window"}
	result[GovModuleSlashingDoubleSignJailDuration] = RangeDescription{Range: "0,1209600000000000", Description: "double sign jail duration"}
	result[GovModuleSlashingMinSignedPerWindow] = RangeDescription{Range: "0.5,0.9", Description: "min signed per window"}
	result[GovModuleSlashingCensorshipJailDuration] = RangeDescription{Range: "0,1209600000000000", Description: "censorship jail duration"}
	result[GovModuleSlashingSlashFractionDoubleSign] = RangeDescription{Range: "0,0.1", Description: "slash fraction double sign"}
	result[GovModuleSlashingSlashFractionDowntime] = RangeDescription{Range: "0,0.1", Description: "slash fraction downtime"}
	result[GovModuleSlashingDowntimJailDuration] = RangeDescription{Range: "0,604800000000000", Description: "downtime jial duration"}
	return result
}
func GetGovModuleParam(module string) ([]byte, error) {
	url := fmt.Sprintf(UrlGovParam, conf.Get().Hub.LcdUrl, module)
	return utils.Get(url)
}

func GetGovModuleParamMap(module string) (map[string]interface{}, error) {
	resAsByte, err := GetGovModuleParam(module)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(resAsByte), &m)
	if err != nil {
		return nil, err
	}

	return m["value"].(map[string]interface{}), nil
}

func GetGovAuthParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleAuth)
}

func GetGovStakeParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleStake)
}

func GetGovMintParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleMint)
}

func GetGovDistrParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleDistr)
}

func GetGovSlashingParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleSlashing)
}

func GetAllGovModuleParam() (map[string]interface{}, error) {

	result := map[string]interface{}{}

	authMap, err := GetGovAuthParam()
	if err != nil {
		return nil, err
	}
	stakeMap, err := GetGovStakeParam()
	if err != nil {
		return nil, err
	}
	mintMap, err := GetGovMintParam()
	if err != nil {
		return nil, err
	}
	distrMap, err := GetGovDistrParam()
	if err != nil {
		return nil, err
	}
	slashingMap, err := GetGovSlashingParam()
	if err != nil {
		return nil, err
	}

	for k, v := range authMap {
		result[k] = v
	}

	for k, v := range stakeMap {
		result[k] = v
	}
	for k, v := range mintMap {
		result[k] = v
	}
	for k, v := range distrMap {
		result[k] = v
	}
	for k, v := range slashingMap {
		result[k] = v
	}
	return result, nil
}
