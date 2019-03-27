package service

import (
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
)

type DelegatorService struct {
	BaseService
}

func (service *DelegatorService) GetModule() Module {
	return Delegator
}

func (service *DelegatorService) QueryDelegation(valAddr string) (info ValInfo) {
	// query delegation info
	var accAddr = utils.Convert(conf.Get().Hub.Prefix.AccAddr, valAddr)

	var query = orm.NewQuery()
	defer query.Release()

	var validator document.Candidate
	query.SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(bson.M{document.Candidate_Field_Address: valAddr}).
		SetResult(&validator)

	//query validator info
	if err := query.Exec(); err != nil {
		logger.Error("validator not found", logger.Any("err", err.Error()))
		return
	}

	var delegations []document.Delegator
	query.Reset().
		SetCollection(document.CollectionNmStakeRoleDelegator).
		SetCondition(bson.M{document.Delegator_Field_ValidatorAddr: valAddr}).
		SetResult(&delegations)

	if err := query.Exec(); err != nil {
		logger.Warn("validator not exist", logger.String("valAddr", valAddr))
		return
	}

	var selfBondShares float64
	var delegatedShares float64
	for _, d := range delegations {
		if d.Address == accAddr {
			selfBondShares = d.Shares
		} else {
			delegatedShares += d.Shares
		}
	}

	rate := validator.Tokens / validator.DelegatorShares

	selfBond := document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: selfBondShares * rate,
	}

	delegated := document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: delegatedShares * rate,
	}

	query.Reset().
		SetCollection(document.CollectionNmBlock).
		SetCondition(bson.M{"block.last_commit.precommits": bson.M{"$elemMatch": bson.M{"validator_address": validator.PubKeyAddr}}})

	var preCommitCount int
	if cnt, err := query.Count(); err == nil {
		preCommitCount = cnt
	}

	//query uptime
	var upTime = getValUpTime(query)[validator.PubKeyAddr]

	return ValInfo{
		selfBond:  selfBond,
		delegated: delegated,
		ValProfile: model.ValProfile{
			PubKey:         validator.PubKey,
			Owner:          accAddr,
			BondHeight:     validator.BondHeight,
			VotingPower:    getVotingPowerFromToken(validator.Tokens),
			CommitBlockNum: int64(preCommitCount),
			UpTime:         upTime,
			Description: model.Description{
				Moniker:  validator.Description.Moniker,
				Identity: validator.Description.Identity,
				Website:  validator.Description.Website,
				Details:  validator.Description.Details,
			},
		},
	}

}

func (service *DelegatorService) GetDeposits(delAddr string) (coin document.Coin) {

	var query = orm.NewQuery()
	defer query.Release()

	var delegations []document.Delegator
	query.SetCollection(document.CollectionNmStakeRoleDelegator).
		SetCondition(bson.M{document.Delegator_Field_Addres: delAddr}).
		SetResult(&delegations)

	if query.Exec() != nil {
		logger.Warn("delegator address not exist", logger.String("delAddr", delAddr))
		return
	}

	var delegationMap = make(map[string]document.Delegator, len(delegations))
	var valAddrs []string
	for _, d := range delegations {
		delegationMap[d.ValidatorAddr] = d
		valAddrs = append(valAddrs, d.ValidatorAddr)
	}

	var condition = bson.M{}
	condition[document.Candidate_Field_Address] = bson.M{
		"$in": valAddrs,
	}
	var validators []document.Candidate
	query.Reset().SetCollection(document.CollectionNmStakeRoleCandidate).
		SetCondition(condition).
		SetResult(&validators)

	if query.Exec() != nil {
		logger.Error("validator not found", logger.Any("valAddrs", valAddrs))
		return
	}

	var totalAmt float64
	for _, v := range validators {
		delegation := delegationMap[v.Address]
		if v.Tokens > 0 && v.DelegatorShares > 0 {
			rate := v.Tokens / v.DelegatorShares
			totalAmt += delegation.Shares * rate
		}
	}
	return document.Coin{
		Denom:  utils.CoinTypeAtto,
		Amount: totalAmt,
	}
}

type ValInfo struct {
	model.ValProfile
	selfBond  document.Coin
	delegated document.Coin
}
