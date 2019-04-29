package service

import (
	"strconv"

	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
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

func (service *BlockService) Query(height int64) model.Block {
	result := model.Block{}
	result.BlockInfo = service.QueryBlockInfo(height)
	result.ValidatorSet = service.GetValidatorSet(height, DefaultPageNum, DefaultPageSize)
	result.TokenFlows = Get(Tx).(*TxService).QueryTokenFlow(height, DefaultPageNum, DefaultPageSize)
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
	var cnt, err = pageQuery(document.CollectionNmBlock, selector, nil, sort, page, size, &blocks)
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
	var selector = bson.M{"height": 1, "time": 1, "num_txs": 1, "hash": 1, "validators.address": 1, "validators.voting_power": 1, "block.last_commit.precommits.validator_address": 1, "meta.header.total_txs": 1}

	sort := desc(document.Block_Field_Height)
	err := queryAll(document.CollectionNmBlock, selector, nil, sort, 10, &blocks)
	if err != nil {
		panic(types.CodeNotFound)
	}
	for _, block := range blocks {
		result = append(result, buildBlock(block))
	}
	return result
}

//TODO
func (service *BlockService) QueryPrecommits(address string, page, size int) (result model.PageVo) {
	var candidate document.Candidate
	var selector = bson.M{"pub_key_addr": 1}
	var query = orm.NewQuery().
		SetCollection(document.CollectionNmStakeRoleCandidate).
		SetSelector(selector).
		SetCondition(bson.M{document.Candidate_Field_Address: address}).
		SetResult(&candidate)
	defer query.Release()
	err := query.Exec()
	if err != nil {
		panic(types.CodeNotFound)
	}

	var data []document.Block
	//var selector = bson.M{"pub_key_addr": 1}
	query.Reset().
		SetCollection(document.CollectionNmBlock).
		SetResult(&data).
		SetCondition(bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": candidate.PubKeyAddr}}}).
		SetSort(desc(document.Block_Field_Height)).
		SetPage(page).
		SetSize(size)

	cnt, err := query.ExecPage()
	if err != nil {
		panic(types.CodeNotFound)
	}
	return model.PageVo{
		Count: cnt,
		Data:  data, //TODO
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
