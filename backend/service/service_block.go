package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"gopkg.in/mgo.v2/bson"
)

type BlockService struct {
	BaseService
}

func (service *BlockService) GetModule() Module {
	return Block
}

func (service *BlockService) Query(height int64) model.BlockInfoVo {
	var selector = bson.M{"height": 1, "time": 1, "num_txs": 1, "hash": 1, "validators.address": 1, "validators.voting_power": 1, "block.last_commit.precommits.validator_address": 1, "block.last_commit.block_id.hash": 1, "meta.header.total_txs": 1}
	var result document.Block

	var err = one(document.CollectionNmBlock, selector, bson.M{document.Block_Field_Height: height}, &result)
	if err != nil {
		panic(types.CodeNotFound)
	}
	return buildBlock(result)
}

func (service *BlockService) QueryList(page, size int) model.PageVo {
	var result []model.BlockInfoVo
	var pageInfo model.PageVo

	var selector = bson.M{"height": 1, "time": 1, "num_txs": 1, "hash": 1, "validators.address": 1, "validators.voting_power": 1, "block.last_commit.precommits.validator_address": 1, "meta.header.total_txs": 1}

	var blocks []document.Block

	sort := desc(document.Block_Field_Height)
	var cnt, err = all(document.CollectionNmBlock, selector, nil, sort, page, size, &blocks)
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
	err := limitQuery(document.CollectionNmBlock, selector, nil, sort, 10, &blocks)
	if err != nil {
		panic(types.CodeNotFound)
	}
	for _, block := range blocks {
		result = append(result, buildBlock(block))
	}
	return result
}

func (service *BlockService) QueryPrecommits(address string, page, size int) (result model.PageVo) {
	c := getDb().C(document.CollectionNmStakeRoleCandidate)
	defer c.Database.Session.Close()

	sort := desc(document.Candidate_Field_BondHeight)
	var candidate document.Candidate
	err := c.Find(bson.M{document.Candidate_Field_Address: address}).Sort(sort).One(&candidate)
	if err != nil {
		panic(types.CodeNotFound)
		return
	}

	var data []document.Block
	sort = desc(document.Block_Field_Height)

	return queryRows(document.CollectionNmBlock, &data, bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": candidate.PubKeyAddr}}}, sort, page, size)
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
