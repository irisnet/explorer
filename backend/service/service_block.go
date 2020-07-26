package service

import (
	"strconv"

	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/explorer/backend/vo"
	"gopkg.in/mgo.v2"
	"fmt"
)

type BlockService struct {
	BaseService
}

func (service *BlockService) GetValidatorSet(height int64, page, size int) vo.ValidatorSet {
	lcdValidators := lcd.ValidatorSet(height)
	if page > 0 {
		page = page - 1
	}

	b := lcd.Block(height)
	if b.Block.Header.Height == 0 {
		panic(types.CodeNotFound)
	}

	validatorArr, err := document.Validator{}.GetValidatorList()

	if err != nil {
		logger.Error("get validator set err", logger.String("error", err.Error()), service.GetTraceLog())
		panic(types.CodeNotFound)
	}

	var proposal string
	for _, v := range validatorArr {
		if v.ProposerAddr == b.BlockMeta.Header.ProposerAddress {
			proposal = v.OperatorAddress
		}
	}

	items := []vo.BlockValidator{}
	for k, v := range lcdValidators.Validators {
		if k >= page*size && k < (page+1)*size {
			var tmp vo.BlockValidator
			tmp.Consensus = v.Address
			tmp.VotingPower = fmt.Sprint(v.VotingPower)
			tmp.ProposerPriority = fmt.Sprint(v.ProposerPriority)
			for _, validator := range validatorArr {
				if validator.ConsensusPubkey == v.PubKey {
					tmp.OperatorAddress = validator.OperatorAddress
					tmp.Moniker = validator.Description.Moniker
					tmp.IsProposer = tmp.OperatorAddress == proposal
				}
			}
			items = append(items, tmp)
		}
	}

	return vo.ValidatorSet{
		Items: items,
		Total: len(lcdValidators.Validators),
	}
}

func (service *BlockService) QueryBlockInfo(height int64) vo.BlockInfo {
	var result vo.BlockInfo

	currentBlock := lcd.Block(height)
	if currentBlock.Block.Header.Height == 0 {
		panic(types.CodeNotFound)
	}

	proposerHexAddr := currentBlock.BlockMeta.Header.ProposerAddress
	validatorDoc, err := document.Validator{}.GetValidatorByProposerAddr(proposerHexAddr)

	result.PropopserMoniker = proposerHexAddr
	if err != nil && err != mgo.ErrNotFound {
		logger.Error("query validator collection  err", logger.String("error", err.Error()), service.GetTraceLog())

	} else {
		result.PropoperAddr = validatorDoc.OperatorAddress
		result.PropopserMoniker = validatorDoc.Description.Moniker
	}

	result.LatestHeight = fmt.Sprint(lcd.BlockLatest().BlockMeta.Header.Height)
	currentBlockRes := lcd.BlockResult(height)
	lcdValidators := lcd.ValidatorSet(height)

	result.TotalValidatorNum = len(lcdValidators.Validators)
	nextBlock := lcd.Block(height + 1)
	var totalVotingPower, precommitVotingPower, precommitValidatorNum int
	for k, v := range lcdValidators.Validators {
		powerAsInt := int(v.VotingPower)

		totalVotingPower += powerAsInt

		if nextBlock.Block.Header.Height != 0 {
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

	result.BlockHash = currentBlock.BlockMeta.BlockID.Hash
	result.BlockHeight = fmt.Sprint(currentBlock.Block.Header.Height)
	result.Timestamp = currentBlock.BlockMeta.Header.Time.UTC()
	result.Transactions = currentBlock.BlockMeta.Header.NumTxs

	return result
}

func (service *BlockService) QueryList(page, size int) vo.BlockForListRespond {

	offset := 0

	if page > 1 {
		offset = (page-1)*size - 1
		size = size + 1
	}

	blocks, err := document.Block{}.GetBlockListByOffsetAndSize(offset, size)

	if page == 1 && len(blocks) > 1 {
		nextHeight := blocks[0].Height + 1
		nextBlock := lcd.Block(nextHeight)
		nextBlockAsDoc := document.Block{}

		votes := []document.Vote{}
		for _, v := range nextBlock.Block.LastCommit.Precommits {
			if len(v.ValidatorAddress) > 0 {
				vote := document.Vote{}
				vote.ValidatorAddress = v.ValidatorAddress
				votes = append(votes, vote)
			}
		}
		nextBlockAsDoc.Height = nextHeight
		nextBlockAsDoc.Block.LastCommit.Precommits = votes
		blocks = append([]document.Block{nextBlockAsDoc}, blocks...)
	}

	if err != nil {
		logger.Error("GetBlockListByOffsetAndSize", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}

	proposerAsHashAddr := make([]string, 0, len(blocks))

	for _, v := range blocks {
		proposerAsHashAddr = append(proposerAsHashAddr, v.ProposalAddr)
	}

	proposerAsHashAddr = utils.RemoveDuplicationStrArr(proposerAsHashAddr)

	validators, err := document.Validator{}.QueryValidatorMonikerOpAddrByHashAddr(proposerAsHashAddr)

	if err != nil {
		logger.Error("QueryValidatorMonikerOpAddrConsensusPubkeyByAddr", logger.String("err", err.Error()))
	}

	validatorMapByHashAddr := map[string]document.Validator{}
	blockMapByHeight := map[int64]document.Block{}

	for _, v := range validators {
		validatorMapByHashAddr[v.ProposerAddr] = v
	}

	for _, v := range blocks {
		blockMapByHeight[v.Height] = v
	}

	blocksAsModel := make([]vo.BlockForList, 0, len(blocks))

	for _, block := range blocks {

		var proposerMoniker, proposerValidatorAddr string
		if v, ok := validatorMapByHashAddr[block.ProposalAddr]; ok {
			proposerMoniker = v.Description.Moniker
			proposerValidatorAddr = v.OperatorAddress
		}

		votingPower := int64(0)
		precommitVotingPower := int64(0)
		precomitValidatorNum := 0

		for _, v := range block.Validators {
			votingPower += v.VotingPower
		}

		if v, ok := blockMapByHeight[block.Height+1]; ok {
			precomitValidatorNum = len(v.Block.LastCommit.Precommits)
			for _, preValidator := range v.Block.LastCommit.Precommits {
				for _, validator := range block.Validators {
					if preValidator.ValidatorAddress == validator.Address {
						precommitVotingPower += validator.VotingPower
						continue
					}
				}
			}
		}

		tmp := vo.BlockForList{
			Height:                  block.Height,
			ProposerMoniker:         proposerMoniker,
			ProposerAsValidatorAddr: proposerValidatorAddr,
			TxnNum:                  block.NumTxs,
			PrecommitValidatorNum:   precomitValidatorNum,
			ValidatorNumForHeight:   len(block.Validators),
			PrecommitVotingPower:    precommitVotingPower,
			VotingPowerForHeight:    votingPower,
			Timestamp:               block.Time.UTC(),
		}

		blocksAsModel = append(blocksAsModel, tmp)
	}

	if len(blocksAsModel) > 1 {
		return blocksAsModel[1:]
	}

	return []vo.BlockForList{}
}

func (service *BlockService) QueryRecent() vo.BlockInfoVoRespond {
	var result []vo.BlockInfoVo

	blockList, err := document.Block{}.GetRecentBlockList()

	if err != nil {
		logger.Error("GetRecentBlockList have error", logger.String("err", err.Error()))
		panic(types.CodeNotFound)
	}
	for _, block := range blockList {
		result = append(result, vo.BlockInfoVo{
			Time:   block.Time.UTC(),
			Height: block.Height,
			NumTxs: block.NumTxs,
		})
	}
	return result
}

func (service *BlockService) QueryLatestHeight() (result vo.LatestHeightRespond) {
	var block = lcd.BlockLatest()
	var height = block.BlockMeta.Header.Height

	blockdb, err := document.Block{}.QueryLatestBlockFromDB()
	if err != nil {
		logger.Error("QueryLatestBlockFromDB have error", logger.String("err", err.Error()))
	}

	result.BlockHeightLcd = height
	result.BlockHeightDB = blockdb.Height
	return result
}
