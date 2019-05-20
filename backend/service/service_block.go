package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

const (
	TxTypeTransfer                    = "Transfer"
	TxTypeStakeCreateValidator        = "CreateValidator"
	TxTypeStakeEditValidator          = "EditValidator"
	TxTypeStakeDelegate               = "Delegate"
	TxTypeStakeBeginUnbonding         = "BeginUnbonding"
	TxTypeBeginRedelegate             = "BeginRedelegate"
	TxTypeUnjail                      = "Unjail"
	TxTypeSetWithdrawAddress          = "SetWithdrawAddress"
	TxTypeWithdrawDelegatorReward     = "WithdrawDelegatorReward"
	TxTypeWithdrawDelegatorRewardsAll = "WithdrawDelegatorRewardsAll"
	TxTypeWithdrawValidatorRewardsAll = "WithdrawValidatorRewardsAll"
	TxTypeSubmitProposal              = "SubmitProposal"
	TxTypeDeposit                     = "Deposit"
	TxTypeVote                        = "Vote"
)

type BlockService struct {
	BaseService
}

func (service *BlockService) GetModule() Module {
	return Block
}

func (service *BlockService) Query(height int64) model.Block {
	result := model.Block{}
	result.BlockInfo = service.QueryBlockInfo(height)
	result.ValidatorSet = service.GetValidatorSet(height, DefaultPageNum, DefaultPageSize)
	//	result.TokenFlows = service.QueryTokenFlow(height, DefaultPageNum, DefaultPageSize)
	result.Proposals = Get(Proposal).(*ProposalService).QueryProposalsByHeight(height)

	return result
}

func (service *BlockService) GetValidatorSet(height int64, page, size int) model.ValidatorSet {
	result := model.ValidatorSet{}
	lcdValidators := lcd.ValidatorSet(height)
	if page > 0 {
		page = page - 1
	}

	var validatorsDoc []document.Validator
	var selector = bson.M{"description.moniker": 1, "operator_address": 1, "consensus_pubkey": 1}

	err := queryAll(document.CollectionNmValidator, selector, nil, "", 0, &validatorsDoc)
	if err != nil {
		logger.Error("get validator set err", logger.String("error", err.Error()), service.GetTraceLog())
		return result
	}

	items := []model.BlockValidator{}
	for k, v := range lcdValidators.Validators {
		if k >= page*size && k < (page+1)*size {
			var tmp model.BlockValidator
			tmp.Consensus = v.PubKey
			tmp.VotingPower = v.VotingPower
			tmp.ProposerPriority = v.ProposerPriority
			for _, validator := range validatorsDoc {
				if validator.ConsensusPubkey == v.PubKey {
					tmp.OperatorAddress = validator.OperatorAddress
					tmp.Moniker = validator.Description.Moniker
				}
			}
			items = append(items, tmp)
		}
	}
	result.Items = items
	result.Total = len(lcdValidators.Validators)
	return result
}

func (service *BlockService) QueryBlockInfo(height int64) model.BlockInfo {
	var result model.BlockInfo

	b := lcd.Block(height)
	if b.Block.Header.Height == "" {
		panic(types.CodeNotFound)
	}

	var selector = bson.M{"description.moniker": 1, "operator_address": 1}
	var validatorDoc document.Validator
	err := queryOne(document.CollectionNmValidator, selector, bson.M{"proposer_addr": b.BlockMeta.Header.ProposerAddress}, &validatorDoc)

	if err != nil {
		logger.Error("query validator collection  err", logger.String("error", err.Error()), service.GetTraceLog())
		return result
	}

	result.LatestHeight = lcd.BlockLatest().BlockMeta.Header.Height

	blockRes := lcd.BlockResult(height)

	if height == 1 {
		result.VoteValidatorNum = InitVoteValidatorNum
		result.ValidatorNum = InitValidatorNum
		result.PrecommitVotingPower = InitPrecommitVotingPower
		result.TotalVotingPower = InitTotalVotingPower
	} else {
		lcdValidators := lcd.ValidatorSet(height - 1)
		result.VoteValidatorNum = len(b.Block.LastCommit.Precommits)
		result.ValidatorNum = len(lcdValidators.Validators)
		var totalVotingPower, precommitVotingPower int
		for k, v := range lcdValidators.Validators {
			powerAsInt, err := strconv.Atoi(v.VotingPower)
			if err != nil {
				logger.Error("strconv VotingPower err", logger.String("error", err.Error()), service.GetTraceLog())
			}
			totalVotingPower += powerAsInt
			for _, precommitValidator := range b.Block.LastCommit.Precommits {
				if strconv.Itoa(k) == precommitValidator.ValidatorIndex {
					precommitVotingPower += powerAsInt
				}
			}
		}
		result.PrecommitVotingPower = precommitVotingPower
		result.TotalVotingPower = totalVotingPower
	}

	for _, v := range blockRes.Results.BeginBlock.Tags {
		if v.Key == "mint-coin" {
			result.Rewards = utils.ParseCoins(v.Value)
		}
	}

	result.PropoperAddr = validatorDoc.OperatorAddress
	result.PropopserMoniker = validatorDoc.Description.Moniker
	result.LastBlock = height - 1
	result.BlockHash = b.BlockMeta.BlockID.Hash
	result.BlockHeight = b.Block.Header.Height
	result.LastBlockHash = b.Block.Header.LastBlockID.Hash
	result.Timestamp = b.BlockMeta.Header.Time
	result.Transactions = b.BlockMeta.Header.NumTxs

	return result
}

func (service *BlockService) QueryList(page, size int) model.PageVo {
	var result []model.BlockInfoVo
	var pageInfo model.PageVo

	var selector = bson.M{"height": 1, "time": 1, "num_txs": 1, "hash": 1, "validators.address": 1, "validators.voting_power": 1, "block.last_commit.precommits.validator_address": 1, "meta.header.total_txs": 1}

	var blocks []document.Block

	sort := desc(document.Block_Field_Height)
	var cnt, err = pageQuery(document.CollectionNmBlock, selector, bson.M{"height": bson.M{"$gt": 0}}, sort, page, size, &blocks)
	if err != nil {
		panic(types.CodeNotFound)
	}
	for _, block := range blocks {
		result = append(result, buildBlock(block))
	}
	pageInfo.Data = result
	pageInfo.Count = cnt
	return pageInfo
}

func (service *BlockService) QueryRecent() []model.BlockInfoVo {
	var result []model.BlockInfoVo
	var blocks []document.Block
	var selector = bson.M{"height": 1, "time": 1, "num_txs": 1}

	sort := desc(document.Block_Field_Height)
	err := queryAll(document.CollectionNmBlock, selector, nil, sort, 10, &blocks)
	if err != nil {
		panic(types.CodeNotFound)
	}
	for _, block := range blocks {
		result = append(result, model.BlockInfoVo{
			Time:   block.Time,
			Height: block.Height,
			NumTxs: block.NumTxs,
		})
	}
	return result
}

func (service *BlockService) QueryTxsExcludeProposalByBlock(height int64, page, size int) model.TxsWithPage {

	items, total, err := service.GetTxsByBlock(height, false, page, size)

	if err != nil {
		logger.Error("QueryTxsExcludeProposal err", logger.String("error", err.Error()))
	}

	return model.TxsWithPage{
		Total: total,
		Items: items,
	}
}

func (service *BlockService) QueryTxsOnlyProposalByBlock(height int64, page, size int) model.TxsWithPage {

	items, total, err := service.GetTxsByBlock(height, true, page, size)

	if err != nil {
		logger.Error("QueryTxsExcludeProposal err", logger.String("error", err.Error()))
	}
	return model.TxsWithPage{
		Total: total,
		Items: items,
	}
}

func (service *BlockService) GetTxsByBlock(height int64, onlyOrExcludeProposal bool, page, size int) ([]model.Tx, int, error) {

	itemsAsDoc := make([]document.CommonTx, 0, size)

	selector := bson.M{
		document.Tx_Field_Hash:   1,
		document.Tx_Field_From:   1,
		document.Tx_Field_To:     1,
		document.Tx_Field_Amount: 1,
		document.Tx_Field_Fee:    1,
		document.Tx_Field_Type:   1,
		document.Tx_Field_Status: 1,
		document.Tx_Field_Time:   1,
	}
	condition := bson.M{}
	if onlyOrExcludeProposal {
		condition = bson.M{
			document.Tx_Field_Height: height,
			document.Tx_Field_Type:   bson.M{"$in": []string{"GovDeposit", "GovDepositBurn", "GovDepositRefund"}}}
	} else {
		condition = bson.M{
			document.Tx_Field_Height: height,
			document.Tx_Field_Type:   bson.M{"$nin": []string{"GovDeposit", "GovDepositBurn", "GovDepositRefund"}}}
	}

	total, err := pageQuery(document.CollectionNmCommonTx, selector, condition, "", page, size, &itemsAsDoc)
	if err != nil {
		return nil, total, err
	}

	items := make([]model.Tx, 0, len(itemsAsDoc))

	// A sign,  transfer coin from B to C
	transferTxHashs := make([]string, 0, len(itemsAsDoc))
	for _, v := range itemsAsDoc {
		tmp := model.Tx{
			Hash:        v.TxHash,
			TxInitiator: v.From,
			To:          v.To,
			Amount:      v.Amount,
			Fee:         v.Fee,
			Type:        v.Type,
			Status:      v.Status,
			Timestamp:   v.Time,
		}
		if service.isTransferTxByType(v.Type) {
			transferTxHashs = append(transferTxHashs, v.TxHash)
		} else {
			tmp.From = v.From
		}
		items = append(items, tmp)
	}

	if len(transferTxHashs) == 0 {
		return items, total, nil
	}

	selectorAsTxMsg := bson.M{
		document.TxMsg_Field_Hash:    1,
		document.TxMsg_Field_Content: 1,
		document.TxMsg_Field_Type:    1,
	}
	conditionAsTxMsg := bson.M{
		document.TxMsg_Field_Hash: bson.M{"$in": transferTxHashs},
	}

	txMsgs := make([]document.TxMsg, 0, len(transferTxHashs))

	if err := queryAll(document.CollectionNmTxMsg, selectorAsTxMsg, conditionAsTxMsg, "", 0, &txMsgs); err != nil {
		return nil, 0, err
	}

	for _, vMsg := range txMsgs {
		for k, vTx := range items {
			if vMsg.Hash == vTx.Hash {
				transferAddr, err := service.GetTransferAddr(vMsg.Type, vMsg.Content)
				if err != nil {
					logger.Error("get transfer addr ", logger.String("err", err.Error()))
					continue
				}
				items[k].From = transferAddr
			}
		}
	}

	return items, total, nil
}

func (service *BlockService) GetTransferAddr(txType interface{}, content string) (string, error) {

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(content), &m)
	if err != nil {
		return "", err
	}
	switch txType {
	case TxTypeBeginRedelegate:
		if v, ok := m["delegator_addr"].(string); ok {
			return v, nil
		}
	case TxTypeWithdrawDelegatorReward:
		if v, ok := m["delegator_addr"].(string); ok {
			return v, nil
		}
	case TxTypeWithdrawDelegatorRewardsAll:
		if v, ok := m["delegator_addr"].(string); ok {
			return v, nil
		}
	default:
		return "", nil
	}
	return "", errors.New(fmt.Sprintf("assert type err, expect string but actual: %T   value: %v", m["degelator"], m["degelator"]))
}

//  MsgBeginRedelegate MsgWithdrawDelegatorReward MsgWithdrawDelegatorRewardsAll
func (service *BlockService) isTransferTxByType(t string) bool {

	switch t {
	case TxTypeBeginRedelegate:
		return true
	case TxTypeWithdrawDelegatorReward:
		return true
	case TxTypeWithdrawDelegatorRewardsAll:
		return true
	case TxTypeWithdrawValidatorRewardsAll:
		return false
	case TxTypeTransfer:
		return false
	case TxTypeStakeCreateValidator:
		return false
	case TxTypeStakeEditValidator:
		return false
	case TxTypeStakeDelegate:
		return false
	case TxTypeStakeBeginUnbonding:
		return false
	case TxTypeUnjail:
		return false
	case TxTypeSetWithdrawAddress:
		return false
	case TxTypeSubmitProposal:
		return false
	case TxTypeDeposit:
		return false
	case TxTypeVote:
		return false
	default:
		return false
	}

}

func buildBlock(block document.Block) (result model.BlockInfoVo) {
	result.Height = block.Height
	result.Hash = block.Hash
	result.Time = block.Time
	result.NumTxs = block.NumTxs
	var validators []model.ValInfo
	for _, v := range block.Validators {
		validators = append(validators, model.ValInfo{
			Address:     v.Address,
			VotingPower: v.VotingPower,
		})
	}
	result.Validators = validators

	var lastCommit []string
	for _, v := range block.Block.LastCommit.Precommits {
		lastCommit = append(lastCommit, v.ValidatorAddress)
	}
	result.LastCommit = lastCommit
	result.TotalTxs = block.Meta.Header.TotalTxs
	result.LastBlockHash = block.Block.LastCommit.BlockID.Hash
	return result
}
