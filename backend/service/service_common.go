package service

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"sync"
	"time"
)

type CommonService struct {
	BaseService
	genesis lcd.GenesisVo
}

func (service *CommonService) GetModule() Module {
	return Common
}

func (service CommonService) QueryText(text string) []model.ResultVo {
	db := getDb()
	defer db.Session.Close()

	var result []model.ResultVo
	i, isUint := utils.ParseUint(text)
	if isUint {
		//查询block信息
		var block document.Block
		blockStore := db.C(document.CollectionNmBlock)
		err := blockStore.Find(bson.M{document.Block_Field_Height: i}).One(&block)
		if err == nil {
			vo := model.ResultVo{
				Type: "block",
				Data: model.SimpleBlockVo{
					Height:    block.Height,
					Timestamp: block.Time,
					Hash:      block.Hash,
				},
			}
			result = append(result, vo)
		}

		//查询proposal信息
		var proposal document.Proposal
		proposalStore := db.C(document.CollectionNmProposal)
		err = proposalStore.Find(bson.M{document.Proposal_Field_ProposalId: i}).One(&proposal)
		if err == nil {
			vo := model.ResultVo{
				Type: "proposal",
				Data: model.SimpleProposalVo{
					ProposalId: proposal.ProposalId,
					Title:      proposal.Title,
					Type:       proposal.Type,
					Status:     proposal.Status,
					SubmitTime: proposal.SubmitTime,
				},
			}
			result = append(result, vo)
		}
	}
	return result
}

func (service CommonService) GetGenesis() lcd.GenesisVo {
	if len(service.genesis.Result.Genesis.ChainID) == 0 {
		result, err := lcd.Genesis()
		if err != nil {
			panic(err)
		}
		service.genesis = result
	}
	return service.genesis
}

func (service CommonService) GetConfig() []document.Config {
	dbm := getDb()
	defer dbm.Session.Close()

	var configs []document.Config
	configStore := dbm.C(document.CollectionNmConfig)
	if err := configStore.Find(nil).All(&configs); err != nil {
		panic(types.ErrForEmpty("config document is not set"))
	}
	return configs
}

func (service CommonService) QueryBjValInfo() interface{} {
	var validator Validator
	var group sync.WaitGroup
	var valInfo ValDetail

	//query validator by address
	var queryValidator = func() {
		url := fmt.Sprintf("%s/staking/validators/%s", conf.Get().Hub.CosmosLcd, conf.Get().Hub.CosmosValAddr)
		resBytes, err := utils.Get(url)

		if err != nil {
			logger.Error("validator not found", service.GetTraceLog(), logger.String("url", url))
			panic(err)
		}

		if err := json.Unmarshal(resBytes, &validator); err != nil {
			logger.Error("validator unmarshal error", service.GetTraceLog(), logger.String("err", err.Error()))
			panic(err)
		}
		var reduction = new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
		var token = utils.ParseBigInt(validator.Tokens)
		valInfo.OperatorAddress = validator.OperatorAddress
		valInfo.Moniker = validator.Description.Moniker
		valInfo.Tokens = new(big.Int).Quo(token, reduction).String()
		valInfo.Rate = validator.Commission.Rate
	}

	//query all bond tokens
	var queryBondToken = func() {
		defer group.Done()
		queryValidator()
		url := fmt.Sprintf("%s/staking/pool", conf.Get().Hub.CosmosLcd)
		resBytes, err := utils.Get(url)

		if err != nil {
			logger.Error("staking/pool not found", service.GetTraceLog(), logger.String("url", url))
			return
		}

		var pool struct {
			NotBondedTokens string `json:"not_bonded_tokens"`
			BondedTokens    string `json:"bonded_tokens"`
		}
		if err := json.Unmarshal(resBytes, &pool); err != nil {
			logger.Error("staking/pool unmarshal error", service.GetTraceLog(), logger.String("err", err.Error()))
			return
		}
		var totalBondTokens = utils.ParseBigFloat(pool.BondedTokens, 4)
		var toketFloat = utils.ParseBigFloat(validator.Tokens, 4)
		valInfo.VotingPowerPer = (new(big.Float).Quo(toketFloat, totalBondTokens)).String()
	}
	//query slash info
	var querySlashInfo = func() {
		defer group.Done()
		url := fmt.Sprintf("%s/slashing/validators/%s/signing_info", conf.Get().Hub.CosmosLcd, conf.Get().Hub.CosmosValPub)
		resBytes, err := utils.Get(url)

		if err != nil {
			logger.Error("slashing/validators not found", service.GetTraceLog(), logger.String("url", url))
			return
		}

		var signInfo struct {
			StartHeight         string `json:"start_height"`
			IndexOffset         string `json:"index_offset"`
			JailedUntil         string `json:"jailed_until"`
			MissedBlocksCounter string `json:"missed_blocks_counter"`
		}
		if err := json.Unmarshal(resBytes, &signInfo); err != nil {
			logger.Error("slashing/validators unmarshal error", service.GetTraceLog(), logger.String("err", err.Error()))
			return
		}
		valInfo.MissedBlocksCnt = int(utils.ParseIntWithDefault(signInfo.MissedBlocksCounter, 0))
	}

	//query slash parameters
	var querySlashParams = func() {
		defer group.Done()
		url := fmt.Sprintf("%s/slashing/parameters", conf.Get().Hub.CosmosLcd)
		resBytes, err := utils.Get(url)

		if err != nil {
			logger.Error("slashing/parameters not found", service.GetTraceLog(), logger.String("url", url))
			return
		}

		var slashParams struct {
			MaxEvidenceAge          string `json:"max_evidence_age"`
			SignedBlocksWindow      string `json:"signed_blocks_window"`
			MinSignedPerWindow      string `json:"min_signed_per_window"`
			DowntimeJailDuration    string `json:"downtime_jail_duration"`
			SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
			SlashFractionDowntime   string `json:"slash_fraction_downtime"`
		}
		if err := json.Unmarshal(resBytes, &slashParams); err != nil {
			logger.Error("slashing/validators unmarshal error", service.GetTraceLog(), logger.String("err", err.Error()))
			return
		}
		valInfo.SignedBlocksWindow = int(utils.ParseIntWithDefault(slashParams.SignedBlocksWindow, 0))

	}

	var goroutinesArr = []func(){queryBondToken, querySlashInfo, querySlashParams}
	for _, r := range goroutinesArr {
		group.Add(1)
		go r()
	}
	group.Wait()
	return valInfo
}

type ValDetail struct {
	OperatorAddress    string `json:"operator_address"`
	Moniker            string `json:"moniker"`
	Tokens             string `json:"tokens"`
	Rate               string `json:"rate"`
	MissedBlocksCnt    int    `json:"missed_blocks_cnt"`
	SignedBlocksWindow int    `json:"signed_blocks_window"`
	VotingPowerPer     string `json:"voting_power_per"`
}
type Validator struct {
	OperatorAddress string `json:"operator_address"`
	ConsensusPubkey string `json:"consensus_pubkey"`
	Jailed          bool   `json:"jailed"`
	Status          int    `json:"status"`
	Tokens          string `json:"tokens"`
	DelegatorShares string `json:"delegator_shares"`
	Description     struct {
		Moniker  string `json:"moniker"`
		Identity string `json:"identity"`
		Website  string `json:"website"`
		Details  string `json:"details"`
	} `json:"description"`
	UnbondingHeight string    `json:"unbonding_height"`
	UnbondingTime   time.Time `json:"unbonding_time"`
	Commission      struct {
		Rate          string    `json:"rate"`
		MaxRate       string    `json:"max_rate"`
		MaxChangeRate string    `json:"max_change_rate"`
		UpdateTime    time.Time `json:"update_time"`
	} `json:"commission"`
	MinSelfDelegation string `json:"min_self_delegation"`
}
