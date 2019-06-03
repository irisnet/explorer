package service

import (
	"strconv"

	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
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

	items := []model.BlockValidator{}
	for k, v := range lcdValidators.Validators {
		if k >= page*size && k < (page+1)*size {
			var tmp model.BlockValidator
			tmp.Consensus = v.Address
			tmp.VotingPower = v.VotingPower
			tmp.ProposerPriority = v.ProposerPriority
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

	proposerHexAddr := currentBlock.BlockMeta.Header.ProposerAddress
	validatorDoc, err := document.Validator{}.GetValidatorByProposerAddr(proposerHexAddr)

	if err != nil {
		logger.Error("query validator collection  err", logger.String("error", err.Error()), service.GetTraceLog())
		result.PropopserMoniker = proposerHexAddr
	} else {
		result.PropoperAddr = validatorDoc.OperatorAddress
		result.PropopserMoniker = validatorDoc.Description.Moniker
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

	result.BlockHash = currentBlock.BlockMeta.BlockID.Hash
	result.BlockHeight = currentBlock.Block.Header.Height
	result.Timestamp = currentBlock.BlockMeta.Header.Time
	result.Transactions = currentBlock.BlockMeta.Header.NumTxs

	return result
}

func (service *BlockService) QueryList(page, size int) model.PageVo {
	var result []model.BlockInfoVo
	var pageInfo model.PageVo

	total, blockList, err := document.Block{}.GetBlockListByPage(page, size)

	if err != nil {
		panic(types.CodeNotFound)
	}
	for _, block := range blockList {
		result = append(result, buildBlock(block))
	}
	pageInfo.Data = result
	pageInfo.Count = total
	return pageInfo
}

func (service *BlockService) QueryRecent() []model.BlockInfoVo {
	var result []model.BlockInfoVo

	blockList, err := document.Block{}.GetRecentBlockList()

	if err != nil {
		panic(types.CodeNotFound)
	}
	for _, block := range blockList {
		result = append(result, model.BlockInfoVo{
			Time:   block.Time,
			Height: block.Height,
			NumTxs: block.NumTxs,
		})
	}
	return result
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
