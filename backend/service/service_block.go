package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var blockSelector = bson.M{"height": 1, "time": 1, "num_txs": 1, "hash": 1, "validators.address": "1", "validators.voting_power": "1", "block.last_commit.precommits.validator_address": "1"}

type BlockService struct {
	BaseService
}

func (service *BlockService) GetModule() Module {
	return Block
}

func (service *BlockService) Query(height int64) model.BlockInfoVo {
	var data = queryRowField(document.CollectionNmBlock, blockSelector, bson.M{document.Block_Field_Height: height})
	var b TmpBlock
	utils.Map2Struct(data, &b)
	return buildBlock(b)
}

func (service *BlockService) QueryList(page, size int) model.PageVo {
	var result []model.BlockInfoVo
	var pageInfo model.PageVo

	sort := desc(document.Block_Field_Height)
	cnt, data := queryRowsField(document.CollectionNmBlock, blockSelector, nil, sort, page, size)

	for _, block := range data {
		var b TmpBlock
		utils.Map2Struct(block, &b)
		result = append(result, buildBlock(b))
	}
	pageInfo.Data = result
	pageInfo.Count = cnt
	return pageInfo
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

func buildBlock(block TmpBlock) (result model.BlockInfoVo) {
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
	return result
}

type TmpBlock struct {
	ID    string `json:"_id"`
	Block struct {
		LastCommit struct {
			Precommits []struct {
				ValidatorAddress string `json:"validator_address"`
			} `json:"precommits"`
		} `json:"last_commit"`
	} `json:"block"`
	Hash       string    `json:"hash"`
	Height     int64     `json:"height"`
	NumTxs     int64     `json:"num_txs"`
	Time       time.Time `json:"time"`
	Validators []struct {
		Address     string `json:"address"`
		VotingPower int64  `json:"voting_power"`
	} `json:"validators"`
}
