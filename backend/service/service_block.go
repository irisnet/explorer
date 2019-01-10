package service

import (
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
)

type BlockService struct {
	BaseService
}

func (service *BlockService) GetModule() Module {
	return Block
}

func (service *BlockService) Query(height int64) (result model.BlockRsp) {
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
		var counter model.Counter
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

func (service *BlockService) QueryPrecommits(address string, page, size int) (result model.Page) {
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

	return queryPage(document.CollectionNmBlock, &data, bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": candidate.PubKeyAddr}}}, sort, page, size)
}
