package controller

import (
	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/service"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
	"strconv"
	"sync"
	"time"
)

func RegisterHome(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		registerNavigationBar,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func registerNavigationBar(r *mux.Router) error {
	doApi(r, types.UrlRegisterNavigationBar, "GET", func(request model.IrisReq) interface{} {
		var block = lcd.BlockLatest()
		var height, ok = utils.ParseInt(block.BlockMeta.Header.Height)
		if !ok {
			panic(types.CodeNotFound)
		}
		var result = NavigationData{
			BlockHeight: height,
			TotalTxs:    utils.ParseIntWithDefault(block.BlockMeta.Header.TotalTxs, 0),
			BlockTime:   block.BlockMeta.Header.Time,
		}
		var proposer = block.BlockMeta.Header.ProposerAddress
		var funGroup []func()
		var group sync.WaitGroup
		var queryVoteInfo = func() {
			defer func() {
				group.Done()
				if r := recover(); r != nil {
					logger.Error("queryVoteInfo error", logger.Any("err", r))
				}
			}()
			var validatorSet = lcd.ValidatorSet(height - 1)

			var totalVp = int64(0)
			var voteVp = int64(0)
			var voteValNum = 0
			for i, commit := range block.Block.LastCommit.Precommits {
				var vp = utils.ParseIntWithDefault(validatorSet.Validators[i].VotingPower, 0)
				if len(commit.Height) > 0 {
					voteVp += vp
					voteValNum++
				}
				totalVp += vp
			}
			result.VotingRatio = float32(voteVp) / float32(totalVp)
			result.VotingTokens = strconv.FormatInt(voteVp, 10)
			result.VoteValNum = voteValNum
			result.ActiveValNum = len(validatorSet.Validators)

			validator := service.Get(service.Candidate).(*service.CandidateService).QueryValidatorByConAddr(proposer)
			result.Moniker = validator.Description.Moniker
			result.OperatorAddr = validator.Address
		}
		funGroup = append(funGroup, queryVoteInfo)
		group.Add(1)

		var queryBondedInfo = func() {
			defer func() {
				group.Done()
				if r := recover(); r != nil {
					logger.Error("queryBondedInfo error", logger.Any("err", r))
				}
			}()
			var poolStake = lcd.StakePool()
			result.BondedTokens = poolStake.BondedTokens
			result.BondedRatio = poolStake.BondedRatio
			result.TotalSupply = poolStake.TotalSupply
		}
		funGroup = append(funGroup, queryBondedInfo)
		group.Add(1)

		var calculateAvgBlockTime = func() {
			defer func() {
				group.Done()
				if r := recover(); r != nil {
					logger.Error("calculateAvgBlockTime error", logger.Any("err", r))
				}
			}()
			var startHeight = int64(1)
			var period = int64(100)
			if height > 100 {
				startHeight = height - 100
			} else {
				period = height
			}
			var tmpBlock = lcd.Block(startHeight)
			var bwtSeconds = block.BlockMeta.Header.Time.Sub(tmpBlock.BlockMeta.Header.Time).Seconds()
			result.AvgBlockTime = float32(bwtSeconds) / float32(period)
		}
		funGroup = append(funGroup, calculateAvgBlockTime)
		group.Add(1)

		for _, fun := range funGroup {
			go fun()
		}

		group.Wait()
		return result
	})
	return nil
}

type NavigationData struct {
	BlockHeight  int64     `json:"block_height"`
	BlockTime    time.Time `json:"block_time"`
	TotalTxs     int64     `json:"total_txs"`
	Tps          float32   `json:"tps"`
	AvgBlockTime float32   `json:"avg_block_time"`
	VotingRatio  float32   `json:"voting_ratio"`
	VotingTokens string    `json:"voting_tokens"`
	VoteValNum   int       `json:"vote_val_num"`
	ActiveValNum int       `json:"active_val_num"`
	BondedRatio  string    `json:"bonded_ratio"`
	BondedTokens string    `json:"bonded_tokens"`
	TotalSupply  string    `json:"total_supply"`
	Moniker      string    `json:"moniker"`
	OperatorAddr string    `json:"operator_addr"`
}
