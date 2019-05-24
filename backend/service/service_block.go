package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type BlockService struct {
	BaseService
}

func (service *BlockService) GetModule() Module {
	return Block
}

func (service *BlockService) GetValidatorSet(height int64, page, size int) model.ValidatorSet {
	lcdValidators := lcd.ValidatorSet(height)
	if page > 0 {
		page = page - 1
	}

	b := lcd.Block(height)
	if b.Block.Header.Height == "" {
		panic(types.CodeNotFound)
	}

	var validatorsDoc []document.Validator
	var selector = bson.M{"description.moniker": 1, "operator_address": 1, "consensus_pubkey": 1, "proposer_addr": 1}

	err := queryAll(document.CollectionNmValidator, selector, nil, "", 0, &validatorsDoc)
	if err != nil {
		logger.Error("get validator set err", logger.String("error", err.Error()), service.GetTraceLog())
		panic(types.CodeNotFound)
	}

	var proposal string
	for _, v := range validatorsDoc {
		if v.ProposerAddr == b.BlockMeta.Header.ProposerAddress {
			proposal = v.OperatorAddress
		}
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
					tmp.IsProposer = tmp.OperatorAddress == proposal
				}
			}
			items = append(items, tmp)
		}
	}

	return model.ValidatorSet{
		Items: items,
		Total: len(lcdValidators.Validators),
	}
}

func (service *BlockService) QueryBlockInfo(height int64) model.BlockInfo {
	var result model.BlockInfo

	currentBlock := lcd.Block(height)
	if currentBlock.Block.Header.Height == "" {
		panic(types.CodeNotFound)
	}

	var selector = bson.M{"description.moniker": 1, "operator_address": 1}
	var validatorDoc document.Validator
	err := queryOne(document.CollectionNmValidator, selector, bson.M{"proposer_addr": currentBlock.BlockMeta.Header.ProposerAddress}, &validatorDoc)

	if err != nil {
		logger.Error("query validator collection  err", logger.String("error", err.Error()), service.GetTraceLog())
		panic(types.CodeNotFound)
	}

	result.LatestHeight = lcd.BlockLatest().BlockMeta.Header.Height
	currentBlockRes := lcd.BlockResult(height)
	lcdValidators := lcd.ValidatorSet(height)

	result.TotalValidatorNum = len(lcdValidators.Validators)
	nextBlock := lcd.Block(height + 1)
	var totalVotingPower, precommitVotingPower, precommitValidatorNum int
	for k, v := range lcdValidators.Validators {
		powerAsInt, err := strconv.Atoi(v.VotingPower)
		if err != nil {
			logger.Error("strconv VotingPower err", logger.String("error", err.Error()), service.GetTraceLog())
		}
		totalVotingPower += powerAsInt

		if nextBlock.Block.Header.Height != "" {
			for _, precommitValidator := range nextBlock.Block.LastCommit.Precommits {
				if strconv.Itoa(k) == precommitValidator.ValidatorIndex {
					precommitVotingPower += powerAsInt
					precommitValidatorNum++
				}
			}
		}
	}

	if precommitVotingPower != 0 {
		result.PrecommitVotingPower = precommitVotingPower
	}
	result.TotalVotingPower = totalVotingPower

	if precommitValidatorNum != 0 {
		result.PrecommitValidatorNum = precommitValidatorNum
	}
	result.TotalValidatorNum = len(lcdValidators.Validators)

	for _, v := range currentBlockRes.Results.BeginBlock.Tags {
		if v.Key == "mint-coin" {
			result.MintCoin = utils.ParseCoin(v.Value)
		}
	}

	result.PropoperAddr = validatorDoc.OperatorAddress
	result.PropopserMoniker = validatorDoc.Description.Moniker
	result.BlockHash = currentBlock.BlockMeta.BlockID.Hash
	result.BlockHeight = currentBlock.Block.Header.Height
	result.Timestamp = currentBlock.BlockMeta.Header.Time
	result.Transactions = currentBlock.BlockMeta.Header.NumTxs

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

func (service *BlockService) QueryTxsExcludeTxGovByBlock(height int64, page, size int) model.TxPage {

	itemsAsDoc, total, err := service.getTxsByBlock(height, false, page, size)

	if err != nil {
		logger.Error("QueryTxsExcludeProposal err", logger.String("error", err.Error()))
		panic(types.CodeNotFound)
	}

	items := make([]model.Tx, 0, len(itemsAsDoc))
	forwardTxHashs := make([]string, 0, len(itemsAsDoc))

	for _, v := range itemsAsDoc {
		tmp := model.Tx{
			Hash:        v.TxHash,
			TxInitiator: v.From,
			To:          v.To,
			Amount:      v.Amount,
			ActualFee:   v.ActualFee,
			Type:        v.Type,
			Status:      v.Status,
			Timestamp:   v.Time,
		}
		if service.isForwardTxByType(v.Type) {
			forwardTxHashs = append(forwardTxHashs, v.TxHash)
		} else {
			tmp.From = v.From
		}
		items = append(items, tmp)
	}

	if len(forwardTxHashs) == 0 {

		for i := 0; i < len(items); i++ {
			items[i].From, items[i].To = service.ParseCoinFlowFromAndTo(items[i].Type, items[i].From, items[i].To)
		}
		return model.TxPage{
			Total: total,
			Items: items,
		}
	}

	selector := bson.M{
		document.TxMsg_Field_Hash:    1,
		document.TxMsg_Field_Content: 1,
		document.TxMsg_Field_Type:    1,
	}
	condition := bson.M{
		document.TxMsg_Field_Hash: bson.M{"$in": forwardTxHashs},
	}

	txMsgs := make([]document.TxMsg, 0, len(forwardTxHashs))

	if err := queryAll(document.CollectionNmTxMsg, selector, condition, "", 0, &txMsgs); err != nil {
		logger.Error("query tx msg", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	for _, vMsg := range txMsgs {
		for k, vTx := range items {
			if vMsg.Hash == vTx.Hash {
				forwarAddr, err := service.getForwardAddr(vMsg.Type, vMsg.Content)
				if err != nil {
					logger.Error("get forward addr ", logger.String("err", err.Error()))
					continue
				}
				items[k].From = forwarAddr
			}
		}
	}

	for i := 0; i < len(items); i++ {
		items[i].From, items[i].To = service.ParseCoinFlowFromAndTo(items[i].Type, items[i].From, items[i].To)
	}

	return model.TxPage{
		Total: total,
		Items: items,
	}
}

func (service *BlockService) ParseCoinFlowFromAndTo(txType, from, to string) (string, string) {
	switch txType {
	case types.TxTypeTransfer:
		return from, to
	case types.TxTypeBurn:
		return from, ""
	case types.TxTypeStakeEditValidator:
		return "", ""
	case types.TxTypeStakeDelegate:
		return from, to
	case types.TxTypeUnjail:
		return "", ""
	case types.TxTypeSetWithdrawAddress:
		return "", ""
	case types.TxTypeStakeCreateValidator:
		return from, to
		//exchange
	case types.TxTypeBeginRedelegate:
		return to, from
	case types.TxTypeWithdrawDelegatorReward:
		return to, from
	case types.TxTypeWithdrawDelegatorRewardsAll:
		return to, from
	case types.TxTypeWithdrawValidatorRewardsAll:
		return to, utils.Convert(conf.Get().Hub.Prefix.AccAddr, from)
	case types.TxTypeStakeBeginUnbonding:
		return to, from
	default:
		return from, to
	}
}

func (service *BlockService) QueryTxsOnlyTxGovByBlock(height int64, page, size int) model.ProposalPage {

	items := make([]model.ProposalInfo, 0, size)
	itemsAsDoc, total, err := service.getTxsByBlock(height, true, page, size)
	if err != nil {
		logger.Error("QueryTxsExcludeProposal err", logger.String("error", err.Error()))
		panic(types.CodeNotFound)
	}

	submitProposalHashArr := make([]string, 0, size)
	voteDepositIdArr := make([]uint64, 0, size)
	for _, v := range itemsAsDoc {
		if v.Type == types.TxTypeSubmitProposal {
			submitProposalHashArr = append(submitProposalHashArr, v.TxHash)
			continue
		}
		if v.ProposalId != 0 {
			voteDepositIdArr = append(voteDepositIdArr, v.ProposalId)
		}
	}

	voteDepositProposals := make([]document.Proposal, 0, size)
	if len(voteDepositIdArr) != 0 {
		var err error
		voteDepositProposals, err = Get(Proposal).(*ProposalService).QueryTypeAndTitleByIds(voteDepositIdArr)
		if err != nil {
			panic(types.CodeNotFound)
		}
	}

	submitProposalTxMsgs := []document.TxMsg{}
	if len(submitProposalHashArr) != 0 {
		submitProposalTxMsgs = service.QueryTxMsgByHashArr(submitProposalHashArr)
	}
	for _, vTx := range itemsAsDoc {
		tmp := model.ProposalInfo{
			ProposalId:  vTx.ProposalId,
			Hash:        vTx.TxHash,
			ActualFee:   vTx.ActualFee,
			TxInitiator: vTx.From,
			TxType:      vTx.Type,
			Status:      vTx.Status,
		}
		if vTx.Type == types.TxTypeSubmitProposal {
			for _, txMsg := range submitProposalTxMsgs {
				if vTx.TxHash == txMsg.Hash {
					proTitle, err := service.GetValueByKey(txMsg.Content, "title")
					if err != nil {
						logger.Error("query proposal type from txMsg", logger.String("k", document.Proposal_Field_Type), logger.String("err", err.Error()), logger.String("JsonStr", txMsg.Content))
					}
					proType, err := service.GetValueByKey(txMsg.Content, "proposalType")
					if err != nil {
						logger.Error("query proposal title from txMsg", logger.String("k", document.Proposal_Field_Title), logger.String("err", err.Error()), logger.String("JsonStr", txMsg.Content))
					}
					tmp.ProposalType = proType
					tmp.ProposalTitle = proTitle
				}
			}
		} else {
			for _, vProposal := range voteDepositProposals {
				if vTx.ProposalId == vProposal.ProposalId {
					tmp.ProposalType = vProposal.Type
					tmp.ProposalTitle = vProposal.Title
				}
			}
		}
		items = append(items, tmp)
	}
	return model.ProposalPage{
		Total: total,
		Items: items,
	}
}

func (service *BlockService) QueryTxMsgByHashArr(hashArr []string) []document.TxMsg {

	selector := bson.M{
		document.TxMsg_Field_Hash:    1,
		document.TxMsg_Field_Content: 1,
		document.TxMsg_Field_Type:    1,
	}
	condition := bson.M{
		document.TxMsg_Field_Hash: bson.M{"$in": hashArr},
	}

	txMsgs := []document.TxMsg{}
	if err := queryAll(document.CollectionNmTxMsg, selector, condition, "", 0, &txMsgs); err != nil {
		logger.Error("query tx msg", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	return txMsgs
}

func (service *BlockService) getTxsByBlock(height int64, onlyOrExcludeProposal bool, page, size int) ([]document.CommonTx, int, error) {

	itemsAsDoc := make([]document.CommonTx, 0, size)

	selector := bson.M{}
	condition := bson.M{}
	if onlyOrExcludeProposal {
		selector = bson.M{
			document.Tx_Field_Hash:       1,
			document.Tx_Field_ActualFee:  1,
			document.Tx_Field_From:       1,
			document.Tx_Field_Type:       1,
			document.Tx_Field_Status:     1,
			document.Tx_Field_ProposalId: 1,
		}

		condition = bson.M{
			document.Tx_Field_Height: height,
			document.Tx_Field_Type:   bson.M{"$in": types.GovernanceList}}
	} else {
		selector = bson.M{
			document.Tx_Field_Hash:      1,
			document.Tx_Field_From:      1,
			document.Tx_Field_To:        1,
			document.Tx_Field_Amount:    1,
			document.Tx_Field_ActualFee: 1,
			document.Tx_Field_Type:      1,
			document.Tx_Field_Status:    1,
			document.Tx_Field_Time:      1,
		}

		condition = bson.M{
			document.Tx_Field_Height: height,
			document.Tx_Field_Type:   bson.M{"$in": types.TxTypeExcludeGov}}
	}

	total, err := pageQuery(document.CollectionNmCommonTx, selector, condition, "", page, size, &itemsAsDoc)
	if err != nil {
		return nil, total, err
	}

	return itemsAsDoc, total, nil
}

func (service *BlockService) GetValueByKey(content, k string) (string, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(content), &m)
	if err != nil {
		return "", err
	}

	if v, ok := m[k].(string); ok {
		return v, nil
	}

	return "", errors.New(fmt.Sprintf("assert type err, expect string but actual: %T   value: %v", m[k], m[k]))
}

func (service *BlockService) getForwardAddr(txType, content string) (string, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(content), &m)
	if err != nil {
		return "", err
	}

	switch txType {
	case types.TxTypeBeginRedelegate:
		if v, ok := m["validator_src_addr"].(string); ok {
			return v, nil
		}
	case types.TxTypeWithdrawDelegatorReward:
		if v, ok := m["validator_addr"].(string); ok {
			return v, nil
		}

	case types.TxTypeWithdrawValidatorRewardsAll:
		if v, ok := m["validator_addr"].(string); ok {
			return v, nil
		}
	}
	return "", nil
}

func (service *BlockService) isForwardTxByType(t string) bool {
	for _, v := range types.ForwardList {
		if v == t {
			return true
		}
	}
	return false
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
