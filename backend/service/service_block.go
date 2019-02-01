package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var blockSelector = bson.M{"height": 1, "time": 1, "num_txs": 1, "hash": 1, "validators.address": 1, "validators.voting_power": 1, "block.last_commit.precommits.validator_address": 1, "meta.header.total_txs": 1}

type BlockService struct {
	BaseService
}

func (service *BlockService) GetModule() Module {
	return Block
}

func (service *BlockService) Query(height int64) (result model.BlockVo) {
	dbm := getDb()
	defer dbm.Session.Close()

	var block document.Block
	err := dbm.C(document.CollectionNmBlock).Find(bson.M{document.Block_Field_Height: height}).Sort("-time").One(&block)

	if err != nil {
		panic(types.CodeNotFound)
	}

	var pres []string
	for _, pre := range block.Block.LastCommit.Precommits {
		pres = append(pres, pre.ValidatorAddress)
	}
	if len(pres) > 0 {
		var candidates []document.Candidate
		err = dbm.C(document.CollectionNmStakeRoleCandidate).Find(bson.M{document.Candidate_Field_PubKeyAddr: bson.M{"$in": pres}}).All(&candidates)
		candidateMap := make(map[string]string)
		for _, candidate := range candidates {
			candidateMap[candidate.PubKeyAddr] = candidate.Address
		}
		result.CandidateMap = candidateMap
	}
	result.Block = block

	//query txs from block height
	var txs []document.CommonTx
	err = dbm.C(document.CollectionNmCommonTx).Find(bson.M{document.Block_Field_Height: height}).All(&txs)
	if err == nil {
		var counter model.CounterVo
		for _, tx := range txs {
			if types.IsGovernanceType(tx.Type) {
				counter.PropoCnt = counter.PropoCnt + 1
			} else {
				counter.TxCnt = counter.TxCnt + 1
			}

		}
		result.Counter = counter
	}
	return
}

func (service *BlockService) QueryList(page, size int) model.Page {
	var data []document.Block
	sort := desc(document.Block_Field_Height)
	return queryPage(document.CollectionNmBlock, &data, nil, sort, page, size)
}

func (service *BlockService) QueryRecent() []model.BlockInfoVo {
	var result []model.BlockInfoVo

	sort := desc(document.Block_Field_Height)
	data, err := LimitQuery(document.CollectionNmBlock, blockSelector, nil, sort, 10)
	if err != nil {
		panic(types.CodeNotFound)
	}
	for _, block := range data {
		var b TmpBlock
		utils.Map2Struct(block, &b)
		result = append(result, buildBlock(b))
	}
	return result
}

func (service *BlockService) QueryPrecommits(address string, page, size int) (result model.Page) {
	c := getDb().C(document.CollectionNmStakeRoleCandidate)
	defer c.Database.Session.Close()

	sort := desc(document.Candidate_Field_BondHeight)
	var candidate document.Candidate
	err := c.Find(bson.M{document.Candidate_Field_Address: address}).Sort(sort).One(&candidate)
	if err != nil {
		panic(types.CodeNotFound)
	}

	var data []document.Block
	sort = desc(document.Block_Field_Height)

	return queryPage(document.CollectionNmBlock, &data, bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": candidate.PubKeyAddr}}}, sort, page, size)
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
	result.TotalTxs = block.Meta.Header.TotalTxs
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
	Meta struct {
		Header struct {
			TotalTxs int64 `json:"total_txs"`
		} `json:"header"`
	} `json:"meta"`
	Hash       string    `json:"hash"`
	Height     int64     `json:"height"`
	NumTxs     int64     `json:"num_txs"`
	Time       time.Time `json:"time"`
	Validators []struct {
		Address     string `json:"address"`
		VotingPower int64  `json:"voting_power"`
	} `json:"validators"`
}
