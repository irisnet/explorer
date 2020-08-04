package lcd

import (
	"fmt"
	"github.com/irisnet/explorer/backend/orm/document"
	"errors"
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
	GovModule         = "gov"
	GovModuleNft      = "nft"
	GovModuleCrisis   = "crisis"
	GovModuleGuardian = "guardian"
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
	GovModuleSlashingSignedBlocksWindow    = "signed_blocks_window"
	GovModuleSlashingMinSignedPerWindow    = "min_signed_per_window"
	GovModuleSlashingSlashFractionDowntime = "slash_fraction_downtime"

	//token key
	GovModuleAssetAssetTaxRate      = "token_tax_rate"
	GovModuleAssetIssueTokenBaseFee = "issue_token_base_fee"
	GovModuleAssetMintTokenFeeRatio = "mint_token_fee_ratio"
)

var GovModuleList = []string{GovModuleAuth, GovModuleStaking, GovModuleMint, GovModuleDistr, GovModuleSlashing, GovModuleAsset, GovModule}
//GovModuleIBC, GovModuleHtlc, GovModuleService, GovModuleCoinSwap, GovModuleCrisis, GovModuleNft, GovModuleOrcale}

type RangeDescription struct {
	Range        string
	Description  string
	InitialValue string
}

type Voterinfo struct {
	Voter      string `json:"voter"`
	ProposalId uint64 `json:"proposal_id"`
	Option     string `json:"option"`
}

var (
	NormalThreshold     string
	NormalMinDeposit    document.Coin
	NormalParticipation string
	NormalVeto          string
)

const (
	//GovModule (iris app module)
	//GovModule = "gov"
	//gov module params key
	NormalMinDepositKey = "min_deposit"
	NormalThresholdKey  = "threshold"
	NormalQuorumKey     = "quorum"
	NormalVetoKey       = "veto"

	Critical  = "Critical"
	Important = "Important"
	Normal    = "Normal"
	// Critical：SoftwareUpgrade
	// Important：ParamChange, CommunityPoolSpend
	// Normal：TxTaxUsage
	ProposalTypeSoftwareUpgrade    = "SoftwareUpgrade"
	ProposalTypeParamChange        = "ParamChange"
	ProposalTypeCommunityPoolSpend = "CommunityPoolSpend"
	ProposalTypePlainText          = "Text"
)

func errorMsg(proposalType string) error {
	return errors.New(fmt.Sprintf("expect proposal type: %v %v %v %v ,but actual: %v",
		ProposalTypeSoftwareUpgrade, ProposalTypeCommunityPoolSpend, ProposalTypeParamChange, ProposalTypePlainText, proposalType))
}

func GetProposalLevelByType(proposalType string) (string, error) {
	switch proposalType {
	case ProposalTypeSoftwareUpgrade:
		return Critical, nil
	case ProposalTypeCommunityPoolSpend, ProposalTypeParamChange:
		return Important, nil
	case ProposalTypePlainText:
		return Normal, nil
	default:
		return "", errorMsg(proposalType)
	}
}

func GetProposalBurnPercentByResult(result string, isRejectVote bool) (float32, error) {
	switch result {
	case document.ProposalStatusPassed:
		return 0.2, nil
	case document.ProposalStatusRejected:
		if isRejectVote {
			return 1, nil
		}
		return 0.2, nil
	default:
		return 0, errors.New(fmt.Sprintf("expect proposal result status: %v %v ,but actual: %v",
			document.ProposalStatusPassed, document.ProposalStatusRejected, result))
	}
}

func GetMinDepositByProposalType(proposalType string) (document.Coin, error) {
	switch proposalType {
	case ProposalTypeSoftwareUpgrade, ProposalTypeCommunityPoolSpend, ProposalTypePlainText, ProposalTypeParamChange:
		return NormalMinDeposit, nil

	default:
		return document.Coin{}, errorMsg(proposalType)
	}

}

func GetPassVetoThresholdAndParticipationMinDeposit(proposalType string) (string, string, string, error) {

	switch proposalType {
	case ProposalTypeSoftwareUpgrade, ProposalTypeCommunityPoolSpend, ProposalTypePlainText, ProposalTypeParamChange:
		return NormalThreshold, NormalVeto, NormalParticipation, nil

	default:
		return "", "", "", errorMsg(proposalType)
	}
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
	//result[GovModuleSlashingSlashFractionDoubleSign] = RangeDescription{Range: "0,0.1", Description: "Slash ratio of DoubleSign"}
	result[GovModuleSlashingSlashFractionDowntime] = RangeDescription{Range: "0,0.1", Description: "Slash ratio of  Downtime"}
	//result[GovModuleSlashingDowntimJailDuration] = RangeDescription{Range: "0,604800000000000", Description: "Jail duration of Downtime"}
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

func GetGovModuleParam(module string) (Params, error) {
	datas, err := client.Params().QueryParams(module)
	if err != nil {
		return nil, err
	}
	var resp Params
	for _, val := range datas {
		resp = append(resp, Param{
			Type:  val.Type,
			Value: val.Value,
		})
	}
	return resp, nil
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

func GetGovGuardianParam() (map[string]interface{}, error) {
	return GetGovModuleParamMap(GovModuleGuardian)
}

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
	votes, err := client.Gov().QueryVotes(proposalid)
	if err != nil {
		return result, err
	}

	for _, val := range votes {
		result = append(result, Voterinfo{Voter: val.Voter, Option: val.Option, ProposalId: val.ProposalID})
	}
	return result, nil
}

func QueryBaseDenom() (string, error) {
	var baseDenom string
	if stakeMap, err := GetGovAssetParam(); err == nil {
		if data, ok := stakeMap["params"]; ok {
			params := data.(map[string]interface{})
			if basefee, ok := params["issue_token_base_fee"]; ok {
				obj := basefee.(map[string]interface{})
				if basedenom, ok := obj["denom"]; ok {
					baseDenom = basedenom.(string)
				}
			}
		}
	} else {
		return baseDenom, err
	}
	return baseDenom, nil
}
