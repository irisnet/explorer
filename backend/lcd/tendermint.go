package lcd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"strings"
)

var (
	CriticalThreshold     string
	CriticalMinDeposit    document.Coin
	CriticalParticipation string
	CriticalVeto          string
	CriticalPenalty       string

	ImportantThreshold     string
	ImportantMinDeposit    document.Coin
	ImportantParticipation string
	ImportantVeto          string
	ImportantPenalty       string

	NormalThreshold     string
	NormalMinDeposit    document.Coin
	NormalParticipation string
	NormalVeto          string
	NormalPenalty       string
)

const (
	//GovModule (iris app module)
	GovModule = "gov"
	//gov module params key
	CriticalThresholdKey     = "critical_threshold"
	CriticalMinDepositKey    = "critical_min_deposit"
	CriticalParticipationKey = "critical_participation"
	CriticalVetoKey          = "critical_veto"
	CriticalPenaltyKey       = "critical_penalty"

	ImportantThresholdKey     = "important_threshold"
	ImportantParticipationKey = "important_participation"
	ImportantMinDepositKey    = "important_min_deposit"
	ImportantVetoKey          = "important_veto"
	ImportantPenaltyKey       = "important_penalty"

	NormalMinDepositKey    = "normal_min_deposit"
	NormalThresholdKey     = "normal_threshold"
	NormalParticipationKey = "normal_participation"
	NormalVetoKey          = "normal_veto"
	NormalPenaltyKey       = "normal_penalty"

	Critical  = "Critical"
	Important = "Important"
	Normal    = "Normal"
	// Critical：SoftwareUpgrade, SystemHalt
	// Important：ParameterChange
	// Normal：TxTaxUsage
	ProposalTypeSoftwareUpgrade   = "SoftwareUpgrade"
	ProposalTypeSystemHalt        = "SystemHalt"
	ProposalTypeParameter         = "Parameter"
	ProposalTypeTokenAddition     = "TokenAddition"
	ProposalTypeTxTaxUsage        = "TxTaxUsage"
	ProposalTypeCommunityTaxUsage = "CommunityTaxUsage"
	ProposalTypePlainText         = "PlainText"
)

func GetProposalLevelByType(proposalType string) (string, error) {
	switch proposalType {
	case ProposalTypeSoftwareUpgrade, ProposalTypeSystemHalt:
		return Critical, nil
	case ProposalTypeParameter, ProposalTypeTokenAddition, ProposalTypeCommunityTaxUsage:
		return Important, nil
	case ProposalTypeTxTaxUsage, ProposalTypePlainText:
		return Normal, nil
	default:
		return "", errors.New(fmt.Sprintf("expect proposal type: %v %v %v %v %v %v %v ,but actual: %v",
			ProposalTypeSoftwareUpgrade, ProposalTypeSystemHalt, ProposalTypeParameter, ProposalTypeTokenAddition, ProposalTypeTxTaxUsage, ProposalTypeCommunityTaxUsage, ProposalTypePlainText, proposalType))
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
	case ProposalTypeSoftwareUpgrade, ProposalTypeSystemHalt:
		return CriticalMinDeposit, nil
	case ProposalTypeParameter, ProposalTypeTokenAddition, ProposalTypeCommunityTaxUsage:
		return ImportantMinDeposit, nil
	case ProposalTypeTxTaxUsage, ProposalTypePlainText:
		return NormalMinDeposit, nil

	default:
		return document.Coin{}, errors.New(fmt.Sprintf("expect proposal type:  %v %v %v %v %v %v %v ,but actual: %v",
			ProposalTypeSoftwareUpgrade, ProposalTypeSystemHalt, ProposalTypeParameter, ProposalTypeTokenAddition, ProposalTypeTxTaxUsage, ProposalTypeCommunityTaxUsage, ProposalTypePlainText, proposalType))
	}

}

func GetPassVetoThresholdAndParticipationMinDeposit(proposalType string) (string, string, string, string, error) {

	switch proposalType {
	case ProposalTypeSoftwareUpgrade, ProposalTypeSystemHalt:
		return CriticalThreshold, CriticalVeto, CriticalParticipation, CriticalPenalty, nil

	case ProposalTypeParameter, ProposalTypeTokenAddition, ProposalTypeCommunityTaxUsage:
		return ImportantThreshold, ImportantVeto, ImportantParticipation, ImportantPenalty, nil
	case ProposalTypeTxTaxUsage, ProposalTypePlainText:
		return NormalThreshold, NormalVeto, NormalParticipation, NormalPenalty, nil

	default:
		return "", "", "", "", errors.New(fmt.Sprintf("expect proposal type: %v %v %v %v %v %v %v ,but actual: %v",
			ProposalTypeSoftwareUpgrade, ProposalTypeSystemHalt, ProposalTypeParameter, ProposalTypeTokenAddition, ProposalTypeTxTaxUsage, ProposalTypeCommunityTaxUsage, ProposalTypePlainText, proposalType))
	}
}

func init() {

	govParamMap, err := GetGovModuleParamMap(GovModule)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if v, ok := govParamMap[CriticalVetoKey].(string); ok {
		CriticalVeto = v
	}

	if v, ok := govParamMap[ImportantVetoKey].(string); ok {
		ImportantVeto = v
	}

	if v, ok := govParamMap[NormalVetoKey].(string); ok {
		NormalVeto = v
	}

	if v, ok := govParamMap[CriticalThresholdKey].(string); ok {
		CriticalThreshold = v
	}

	if v, ok := govParamMap[CriticalMinDepositKey].([]interface{}); ok {
		if len(v) > 0 {
			first := v[0].(map[string]interface{})

			coinAsUtils := utils.ParseCoin(fmt.Sprintf("%v%v", first["amount"], first["denom"]))
			CriticalMinDeposit.Amount = coinAsUtils.Amount
			CriticalMinDeposit.Denom = coinAsUtils.Denom
		}
	}
	if v, ok := govParamMap[CriticalParticipationKey].(string); ok {
		CriticalParticipation = v
	}

	if v, ok := govParamMap[ImportantThresholdKey].(string); ok {
		ImportantThreshold = v
	}

	if v, ok := govParamMap[ImportantMinDepositKey].([]interface{}); ok {
		if len(v) > 0 {
			first := v[0].(map[string]interface{})
			coinAsUtils := utils.ParseCoin(fmt.Sprintf("%v%v", first["amount"], first["denom"]))
			ImportantMinDeposit.Amount = coinAsUtils.Amount
			ImportantMinDeposit.Denom = coinAsUtils.Denom
		}
	}

	if v, ok := govParamMap[ImportantParticipationKey].(string); ok {
		ImportantParticipation = v
	}

	if v, ok := govParamMap[NormalThresholdKey].(string); ok {
		NormalThreshold = v
	}

	if v, ok := govParamMap[NormalMinDepositKey].([]interface{}); ok {
		if len(v) > 0 {
			first := v[0].(map[string]interface{})
			coinAsUtils := utils.ParseCoin(fmt.Sprintf("%v%v", first["amount"], first["denom"]))
			NormalMinDeposit.Amount = coinAsUtils.Amount
			NormalMinDeposit.Denom = coinAsUtils.Denom
		}
	}

	if v, ok := govParamMap[NormalParticipationKey].(string); ok {
		NormalParticipation = v
	}

	if v, ok := govParamMap[CriticalPenaltyKey].(string); ok {
		CriticalPenalty = v
	}

	if v, ok := govParamMap[ImportantPenaltyKey].(string); ok {
		ImportantPenalty = v
	}

	if v, ok := govParamMap[NormalPenaltyKey].(string); ok {
		NormalPenalty = v
	}

}

func NodeInfo() (result NodeInfoVo, err error) {

	nodeinfo, err := client.Tendermint().QueryNodeInfo()
	if err != nil {
		return result, err
	}
	data, _ := json.Marshal(nodeinfo)
	fmt.Println(string(data))
	result = NodeInfoVo{
		Version: nodeinfo.Version,
		Network: nodeinfo.Network,
		Moniker: nodeinfo.Moniker,
	}
	return result, nil
}

func NodeVersion() (result string, err error) {

	nodeversion, err := client.Tendermint().QueryNodeVersion()
	if err != nil {
		return result, err
	}

	result = strings.Split(nodeversion, "-")[0]
	return result, nil
}

func GetGenesisAppState() (map[string]interface{}, error) {

	genesisdoc, err := client.Tendermint().QueryGenesis()
	if err != nil {
		return nil, err
	}
	resBytes, _ := json.Marshal(genesisdoc)

	var m map[string]interface{}
	if err := json.Unmarshal(resBytes, &m); err != nil {
		return nil, err
	}

	appStateMap := m["app_state"].(map[string]interface{})

	return appStateMap, nil
}

func GetGenesisAppStateGovParam() (map[string]interface{}, error) {
	appStateMap, err := GetGenesisAppState()
	if err != nil {
		return nil, err
	}

	govMap := appStateMap[GovModule].(map[string]interface{})
	govParamMap := govMap["params"].(map[string]interface{})

	return govParamMap, nil
}

func GetGenesisGovModuleParamMap() (map[string]interface{}, error) {

	appStateMap, err := GetGenesisAppState()
	if err != nil {
		return nil, err
	}

	authMap := appStateMap[GovModuleAuth].(map[string]interface{})
	authParamMap := authMap["params"].(map[string]interface{})

	stakeMap := appStateMap[GovModuleStake].(map[string]interface{})
	stakeParamMap := stakeMap["params"].(map[string]interface{})

	mintMap := appStateMap[GovModuleMint].(map[string]interface{})
	mintParamMap := mintMap["params"].(map[string]interface{})

	distrMap := appStateMap[GovModuleDistr].(map[string]interface{})
	distrParamMap := distrMap["params"].(map[string]interface{})

	slashingMap := appStateMap[GovModuleSlashing].(map[string]interface{})
	slashingParamMap := slashingMap["params"].(map[string]interface{})

	for k, v := range distrParamMap {
		slashingParamMap[k] = v
	}

	for k, v := range mintParamMap {
		slashingParamMap[k] = v
	}

	for k, v := range stakeParamMap {
		slashingParamMap[k] = v
	}
	for k, v := range authParamMap {
		slashingParamMap[k] = v
	}

	return slashingParamMap, nil
}

func Block(height int64) (result BlockVo) {

	block, err := client.Tendermint().QueryBlock(height)
	if err != nil {
		return result
	}
	//data,_ := json.Marshal(block)
	//fmt.Println(string(data))

	result = BlockVo{
		BlockMeta: BlockMeta{
			BlockID: BlockID{
				Hash: block.Header.LastBlockID.Hash.String(),
			},
			Header: BlockHeader{
				Height:          block.Header.Height,
				ProposerAddress: block.Header.ProposerAddress.String(),
				Time:            block.Header.Time,
				NumTxs:          fmt.Sprint(len(block.Data.Txs)),
			},
		},
		Block: BlockI{
			Header: Header{
				Height: block.LastCommit.Height,
			},
			LastCommit: LastCommit{
				Precommits: Precommits{
				},
			},
		},
	}
	return result
}

func BlockLatest() (result BlockVo) {

	block, err := client.Tendermint().QueryBlockLatest()
	if err != nil {
		return result
	}
	//data,_ := json.Marshal(block)
	//fmt.Println(string(data))

	result = BlockVo{
		BlockMeta: BlockMeta{
			BlockID: BlockID{
				Hash: block.Header.LastBlockID.Hash.String(),
			},
			Header: BlockHeader{
				Height:          block.Header.Height,
				ProposerAddress: block.Header.ProposerAddress.String(),
				Time:            block.Header.Time,
				NumTxs:          fmt.Sprint(len(block.Data.Txs)),
			},
		},
		Block: BlockI{
			Header: Header{
				Height: block.LastCommit.Height,
			},
			LastCommit: LastCommit{
				Precommits: Precommits{

				},
			},
		},
	}
	return result
}

func ValidatorSet(height int64) (result ValidatorSetVo) {
	validatorset, err := client.Tendermint().QueryValidators(height)
	if err != nil {
		return result
	}
	result = ValidatorSetVo{
		BlockHeight: fmt.Sprint(validatorset.BlockHeight),
	}

	for _, val := range validatorset.Validators {
		result.Validators = append(result.Validators, StakeValidatorVo{
			Address:          val.Bech32Address,
			PubKey:           val.Bech32PubKey,
			ProposerPriority: val.ProposerPriority,
			VotingPower:      val.VotingPower,
		})
	}
	return result
}

func LatestValidatorSet() (result ValidatorSetVo) {
	url := fmt.Sprintf(UrlValidatorSetLatest, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		return result
	}
	return result
}

func BlockResult(height int64) (result BlockResultVo) {

	//url := fmt.Sprintf(UrlBlocksResult, conf.Get().Hub.LcdUrl, height)
	//resBytes, err := utils.Get(url)
	//if err != nil {
	//	return result
	//}
	//
	//if err := json.Unmarshal(resBytes, &result); err != nil {
	//	return result
	//}

	blockresult, err := client.Tendermint().QueryBlockResult(height)
	if err != nil {
		return result
	}
	fmt.Println(blockresult)

	result = BlockResultVo{

	}

	return result

}

func BlockCoinFlow(txhash string) (result BlockCoinFlowVo) {
	url := fmt.Sprintf(UrlTxsTxHash, conf.Get().Hub.LcdUrl, txhash)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		return result
	}
	return result
}
