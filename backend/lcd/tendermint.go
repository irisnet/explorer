package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/irishub-sdk-go/types"
	"strings"
	"github.com/tendermint/tendermint/crypto"
)



func NodeInfo() (result NodeInfoVo, err error) {

	nodeinfo, err := client.Tendermint().QueryNodeInfo()
	if err != nil {
		return result, err
	}
	result = NodeInfoVo{
		Version: nodeinfo.NodeInfo.Version,
		Network: nodeinfo.NodeInfo.Network,
		Moniker: nodeinfo.NodeInfo.Moniker,
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

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(genesisdoc.AppState), &m); err != nil {
		return nil, err
	}

	//appStateMap := m["app_state"].(map[string]interface{})

	return m, nil
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

	stakeMap := appStateMap[GovModuleStaking].(map[string]interface{})
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
	var precommits []Precommit
	for idx := range block.LastCommit.Signatures {
		vote := block.LastCommit.GetVote(idx)
		precommits = append(precommits, Precommit{
			Height:           fmt.Sprint(vote.Height),
			ValidatorAddress: vote.ValidatorAddress.String(),
			ValidatorIndex:   fmt.Sprint(vote.ValidatorIndex),
		})
	}

	result = BlockVo{
		BlockMeta: BlockMeta{
			BlockID: BlockID{
				Hash: block.Header.LastBlockID.Hash.String(),
			},
			Header: BlockHeader{
				Height:          block.Height,
				ProposerAddress: block.Header.ProposerAddress.String(),
				Time:            block.Header.Time,
				NumTxs:          fmt.Sprint(len(block.Txs)),
			},
		},
		Block: BlockI{
			Header: Header{
				Height: block.Height,
			},
			LastCommit: LastCommit{
				Precommits: precommits,
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

	var precommits []Precommit
	for idx := range block.LastCommit.Signatures {
		vote := block.LastCommit.GetVote(idx)
		precommits = append(precommits, Precommit{
			Height:           fmt.Sprint(vote.Height),
			ValidatorAddress: vote.ValidatorAddress.String(),
			ValidatorIndex:   fmt.Sprint(vote.ValidatorIndex),
		})
	}

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
				Height: block.Height,
			},
			LastCommit: LastCommit{
				Precommits: precommits,
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
		var pubKey crypto.PubKey
		if bz, err := client.Cdc.MarshalJSON(val.PubKey); err == nil {
			_ = client.Cdc.UnmarshalJSON(bz, &pubKey)
		}
		bech32Addr, _ := types.ConsAddressFromHex(val.Address)
		bech32PubKey, _ := types.Bech32ifyConsPub(pubKey)
		result.Validators = append(result.Validators, StakeValidatorVo{
			Address:          bech32Addr.String(),
			PubKey:           bech32PubKey,
			ProposerPriority: val.ProposerPriority,
			VotingPower:      val.VotingPower,
		})
	}
	return result
}

func LatestValidatorSet() (result ValidatorSetVo) {

	validator, err := client.Tendermint().QueryValidatorsLatest()
	if err != nil {
		return result
	}
	result = ValidatorSetVo{
		BlockHeight:fmt.Sprint(validator.BlockHeight),

	}
	for _, val := range validator.Validators {
		var pubKey crypto.PubKey
		if bz, err := client.Cdc.MarshalJSON(val.PubKey); err == nil {
			_ = client.Cdc.UnmarshalJSON(bz, &pubKey)
		}
		bech32Addr, _ := types.ConsAddressFromHex(val.Address)
		bech32PubKey, _ := types.Bech32ifyConsPub(pubKey)
		result.Validators = append(result.Validators, StakeValidatorVo{
			Address:          bech32Addr.String(),
			PubKey:           bech32PubKey,
			ProposerPriority: val.ProposerPriority,
			VotingPower:      val.VotingPower,
		})
	}
	return result
}

func BlockResult(height int64) (result BlockResultVo, txsnum int) {

	blockresult, err := client.Tendermint().QueryBlockResult(height)
	if err != nil {
		logger.Error("Query Block Result error",
			logger.String("err", err.Error()))
		return
	}

	var events BeginBlockEvents
	for _, val := range blockresult.BeginBlockEvents {
		event := BlockEvent{
			Type: val.Type,
		}
		event.Attributes = make(map[string]string, len(val.Attributes))

		for _, one := range val.Attributes {
			event.Attributes[one.Key] = one.Value
		}

		events = append(events, event)
	}

	result = BlockResultVo{
		Results: Results{
			BeginBlock: events,
		},
	}
	txsnum = len(blockresult.TxsResults)

	return

}


