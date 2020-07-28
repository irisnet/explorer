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
	GovModuleStaking  = "staking"
	GovModuleMint     = "mint"
	GovModuleDistr    = "distribution"
	GovModuleSlashing = "slashing"
	GovModuleAsset    = "token"
	GovModuleIBC      = "ibc"
	GovModuleHtlc     = "htlc"
	GovModuleService  = "service"
	GovModuleRand     = "rand"
	GovModuleOrcale   = "oracle"
	GovModuleCoinSwap = "coinswap"
	GovModuleGov      = "gov"
	GovModuleNft      = "nft"
	GovModuleCrisis   = "crisis"
	//auth key
	GovModuleAuthGasPriceThreshold = "gas_price_threshold"
	GovModuleAuthTxSize            = "tx_size"
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
	//GovModuleSlashingSlashFractionCensorship = "slash_fraction_censorship"
	//GovModuleSlashingMaxEvidenceAge          = "max_evidence_age"
	GovModuleSlashingSignedBlocksWindow = "signed_blocks_window"
	//GovModuleSlashingDoubleSignJailDuration  = "double_sign_jail_duration"
	GovModuleSlashingMinSignedPerWindow = "min_signed_per_window"
	//GovModuleSlashingCensorshipJailDuration  = "censorship_jail_duration"
	GovModuleSlashingSlashFractionDoubleSign = "slash_fraction_double_sign"
	GovModuleSlashingSlashFractionDowntime   = "slash_fraction_downtime"
	GovModuleSlashingDowntimJailDuration     = "downtime_jail_duration"

	//asset key
	GovModuleAssetAssetTaxRate      = "token_tax_rate"
	GovModuleAssetIssueTokenBaseFee = "issue_token_base_fee"
	GovModuleAssetMintTokenFeeRatio = "mint_token_fee_ratio"
	//GovModuleAssetCreateGatewayBaseFee = "create_gateway_base_fee"
	//GovModuleAssetGatewayAssetFeeRatio = "gateway_asset_fee_ratio"
)

var GovModuleList = []string{GovModuleAuth, GovModuleStaking, GovModuleMint, GovModuleDistr, GovModuleSlashing, GovModuleAsset, GovModuleGov}
//GovModuleIBC, GovModuleHtlc, GovModuleService, GovModuleCoinSwap, GovModuleCrisis, GovModuleNft, GovModuleOrcale, GovModuleGov}

type RangeDescription struct {
	Range        string
	Description  string
	InitialValue string
}

type Voterinfo struct {
	Voter      string `json:"voter"`
	ProposalId string `json:"proposal_id"`
	Option     string `json:"option"`
}

func GetAuthKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleAuthGasPriceThreshold] = RangeDescription{Range: "0,1000000000000000000", Description: "Minimum of gas price"}
	result[GovModuleAuthTxSize] = RangeDescription{Range: "500,1500", Description: "Transaction size limit"}
	return result
}

func GetStakeKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleStakeUnbondingTime] = RangeDescription{Range: "1209600000000000,", Description: "Wait time for undelegation"}
	result[GovModuleStakeMaxValidators] = RangeDescription{Range: "100,200", Description: "Maximum number of validators allowed"}
	return result
}

func GetDistrKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleDistrCommunityTax] = RangeDescription{Range: "0,0.2", Description: "Ratio of contributions to community funds"}
	result[GovModuleDistrBonusProposerReward] = RangeDescription{Range: "0,0.08", Description: "Maximum additional bonus ratio"}
	result[GovModuleDistrBaseProposerReward] = RangeDescription{Range: "0,0.02", Description: "Standard ratio of the block reward"}
	return result
}

func GetMintKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleMintInflation] = RangeDescription{Range: "0,0.2", Description: "Inflation rate"}
	return result
}

func GetSlashingKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}

	//result[GovModuleSlashingSlashFractionCensorship] = RangeDescription{Range: "0,0.1", Description: "Slash ratio of Censorship"}
	//result[GovModuleSlashingMaxEvidenceAge] = RangeDescription{Range: "34560,", Description: "Acceptable earliest time of the evidence"}
	result[GovModuleSlashingSignedBlocksWindow] = RangeDescription{Range: "100,140000", Description: "Number of blocks for slash statistics window"}
	//result[GovModuleSlashingDoubleSignJailDuration] = RangeDescription{Range: "0,1209600000000000", Description: "Jail duration of DoubleSign"}
	result[GovModuleSlashingMinSignedPerWindow] = RangeDescription{Range: "0.5,0.9", Description: "Minimum voting ratio in the slash window"}
	//result[GovModuleSlashingCensorshipJailDuration] = RangeDescription{Range: "0,1209600000000000", Description: "Jail duration of Censorship"}
	result[GovModuleSlashingSlashFractionDoubleSign] = RangeDescription{Range: "0,0.1", Description: "Slash ratio of DoubleSign"}
	result[GovModuleSlashingSlashFractionDowntime] = RangeDescription{Range: "0,0.1", Description: "Slash ratio of  Downtime"}
	result[GovModuleSlashingDowntimJailDuration] = RangeDescription{Range: "0,604800000000000", Description: "Jail duration of Downtime"}
	return result
}

func GetAssetKeyWithRangeMap() map[string]RangeDescription {
	result := map[string]RangeDescription{}
	result[GovModuleAssetAssetTaxRate] = RangeDescription{InitialValue: "0.4", Range: "0,1", Description: "Community Tax rate for issuing or minting assets"}
	result[GovModuleAssetIssueTokenBaseFee] = RangeDescription{InitialValue: "60000", Range: "0,", Description: "Benchmark fees for issuing Fungible Token"}
	result[GovModuleAssetMintTokenFeeRatio] = RangeDescription{InitialValue: "0.1", Range: "0,1", Description: "Fungible Token mint rate (relative to the issue fee)"}
	//result[GovModuleAssetCreateGatewayBaseFee] = RangeDescription{InitialValue: "120000", Range: "0,", Description: "Benchmark fees for creating Gateways"}
	//result[GovModuleAssetGatewayAssetFeeRatio] = RangeDescription{InitialValue: "0.1", Range: "0,1", Description: "Rate of issuing gateway tokens (relative to the issue fee of native tokens)"}
	return result
}

func GetGovModuleParam(module string) ([]byte, error) {
	url := fmt.Sprintf(UrlGovParam, conf.Get().Hub.LcdUrl, module)
	return utils.Get(url)
}

func GetGovModuleParamMap(module string) (map[string]interface{}, error) {
	appstate, err := GetGenesisAppState()
	if err != nil {
		return nil, err
	}
	data, ok := appstate[module]
	if !ok {
		return nil, fmt.Errorf("not find module:%v", module)
	}

	return data.(map[string]interface{}), nil
}

func GetGovAuthParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleAuth)
}

func GetGovStakeParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleStaking)
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

func GetGovAssetParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleAsset)
}

//func GetAllGovModuleParam() (map[string]interface{}, error) {
//
//	result := map[string]interface{}{}
//
//	authMap, err := GetGovAuthParam()
//	if err != nil {
//		return nil, err
//	}
//	stakeMap, err := GetGovStakeParam()
//	if err != nil {
//		return nil, err
//	}
//	mintMap, err := GetGovMintParam()
//	if err != nil {
//		return nil, err
//	}
//	distrMap, err := GetGovDistrParam()
//	if err != nil {
//		return nil, err
//	}
//	slashingMap, err := GetGovSlashingParam()
//	if err != nil {
//		return nil, err
//	}
//	assetMap, err := GetGovAssetParam()
//	if err != nil {
//		return nil, err
//	}
//
//	for k, v := range authMap {
//		result[k] = v
//	}
//
//	for k, v := range stakeMap {
//		result[k] = v
//	}
//	for k, v := range mintMap {
//		result[k] = v
//	}
//	for k, v := range distrMap {
//		result[k] = v
//	}
//	for k, v := range slashingMap {
//		result[k] = v
//	}
//	for k, v := range assetMap {
//		result[k] = v
//	}
//	return result, nil
//}

func GetAllGovModuleParam() map[string]interface{} {
	result := map[string]interface{}{}
	for _, module := range GovModuleList {
		if value, err := GetGovModuleParamMap(module); err == nil {
			if data, ok := value["params"]; ok {
				for k, v := range data.(map[string]interface{}) {
					result[k] = v
				}
			} else {
				for k, v := range value {
					result[k] = v
				}
			}
		}
	}
	return result
}

func GetProposalVoters(proposalid uint64) (result []Voterinfo, err error) {
	url := fmt.Sprintf(UrlProposalVoters, conf.Get().Hub.LcdUrl, proposalid)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		return result, err
	}
	return result, nil
}
