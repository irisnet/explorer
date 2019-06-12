package task

import (
	"time"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/model"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/types"
	"github.com/irisnet/explorer/backend/utils"
)

var validatorMap = make(map[string]model.ValVotingPowerChangeVo)

type ResValidatorPreCommits struct {
	Id            ID    `bson:"_id"`
	PreCommitsNum int64 `bson:"num"`
}

type ID struct {
	Address string `bson:"address"`
	Hour    string `bson:"hour"`
}

type UpTimeChangeTask struct{}

func (task UpTimeChangeTask) Name() string {
	return "uptime_change_task"
}

func (task UpTimeChangeTask) Start() {
	task.init()
	utils.RunTimer(10, utils.Min, func() {
		logger.Info("UptimeChangeVo task start")
		var uptimeChange model.UptimeChangeVo
		var blocks []document.Block
		var firstBlock document.Block
		var lastBlock document.Block
		var err error

		uptimeChangeAsDoc, err := document.UptimeChange{}.QueryOne()
		uptimeChange.Address = uptimeChangeAsDoc.Address
		uptimeChange.Time = uptimeChangeAsDoc.Time
		uptimeChange.Uptime = uptimeChangeAsDoc.Uptime

		if err != nil {
			logger.Error("uptime change query one ", logger.String("err", err.Error()))
		}

		var startTime time.Time
		d, _ := time.ParseDuration("1h") //1 hour later
		if uptimeChange.Time == "" {
			firstBlock, err = document.Block{}.QueryOneBlockOrderByHeightAsc()
			if err != nil {
				logger.Error("query one blcok ", logger.String("err", err.Error()))
			}

			startTime = firstBlock.Time
		} else {
			startTime, _ = time.ParseInLocation("2006-01-02 15", uptimeChange.Time, time.UTC)
			startTime = startTime.Add(d)
		}

		lastBlock, err = document.Block{}.QueryOneBlockOrderByHeightDesc()
		if err != nil {
			logger.Error("can't find any block")
			return
		}

		lastTime := lastBlock.Time
		endTime := startTime.Add(d)

		startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), 0, 0, 0, startTime.Location())
		endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour(), 0, 0, 0, endTime.Location())

		logger.Info("UptimeChangeVo", logger.String("startTime", startTime.UTC().Format("2006-01-02 15")), logger.String("endTime", endTime.UTC().Format("2006-01-02 15")))
		if !endTime.Before(lastTime) {
			logger.Info("UptimeChangeVo end", logger.String("startTime", startTime.Format("2006-01-02 15")))
			return
		}

		blocks, err = document.Block{}.QueryBlocksByDurationWithHeightAsc(startTime, endTime)

		if err != nil {
			logger.Error("query blocks by duration with height asc", logger.String("errr", err.Error()))
		}

		for len(blocks) == 0 {
			//往前推进一个小时
			startTime = startTime.Add(d)
			endTime = endTime.Add(d)
			if !endTime.Before(lastTime) {
				logger.Info("UptimeChangeVo end", logger.String("startTime", startTime.Format("2006-01-02 15")))
				return
			}
			blocks, err = document.Block{}.QueryBlocksByDurationWithHeightAsc(startTime, endTime)
			if err != nil {
				logger.Error("query blocks by duration with height asc", logger.String("errr", err.Error()))
			}
		}

		if err != nil {
			logger.Error("can't find any block")
			return
		}

		//init validatorMap if validatorMap length is 0
		if len(validatorMap) == 0 && len(blocks) > 0 {
			for _, validator := range blocks[0].Validators {
				powerChange := document.PowerChange{
					Address: validator.Address,
					Power:   validator.VotingPower,
					Time:    blocks[0].Time,
					Change:  types.Change,
				}
				validatorMap[validator.Address] = model.ValVotingPowerChangeVo{
					Address: validator.Address,
					Power:   validator.VotingPower,
					Time:    blocks[0].Time,
					Change:  types.Change,
				}
				if err := powerChange.Insert(); err != nil {
					logger.Error("power change insert", logger.String("err", err.Error()))
				}
			}

			uptimeMap := make(map[string]int)
			for _, block := range blocks {

				// power change handle
				validatorMapNow := make(map[string]int64)
				for _, validator := range block.Validators {
					validatorMapNow[validator.Address] = validator.VotingPower

					// validator add or update
					if _, ok := validatorMap[validator.Address]; !ok {
						powerChange := document.PowerChange{
							Address: validator.Address,
							Power:   validator.VotingPower,
							Time:    block.Time,
							Height:  block.Height,
							Change:  types.Change,
						}
						validatorMap[validator.Address] = model.ValVotingPowerChangeVo{
							Address: validator.Address,
							Power:   validator.VotingPower,
							Time:    block.Time,
							Height:  block.Height,
							Change:  types.Change,
						}
						if err := powerChange.Insert(); err != nil {
							logger.Error("power change insert", logger.String("err", err.Error()))
						}

					} else {
						if validatorMap[validator.Address].Power != validator.VotingPower {
							powerChange := document.PowerChange{
								Address: validator.Address,
								Power:   validator.VotingPower,
								Time:    block.Time,
								Height:  block.Height,
								Change:  types.Change,
							}
							validatorMap[validator.Address] = model.ValVotingPowerChangeVo{
								Address: validator.Address,
								Power:   validator.VotingPower,
								Time:    block.Time,
								Height:  block.Height,
								Change:  types.Change,
							}
							if err := powerChange.Insert(); err != nil {
								logger.Error("power change insert", logger.String("err", err.Error()))
							}
						}
					}
				}

				// validator slash
				for k, v := range validatorMap {
					if v.Address == "" {
						continue
					}
					if _, ok := validatorMapNow[k]; !ok {
						powerChange := document.PowerChange{
							Address: v.Address,
							Power:   v.Power,
							Time:    block.Time,
							Height:  block.Height,
							Change:  types.Slash,
						}

						delete(validatorMap, k)
						if err := powerChange.Insert(); err != nil {
							logger.Error("power change insert", logger.String("err", err.Error()))
						}
					}
				}

				// uptime handle
				for _, validator := range block.Validators {
					if _, ok := uptimeMap[validator.Address]; !ok {
						uptimeMap[validator.Address] = 0
					}
				}
				for _, commit := range block.Block.LastCommit.Precommits {
					uptimeMap[commit.ValidatorAddress]++
				}

				doneTime := startTime.UTC().Format("2006-01-02 15")
				for k, v := range uptimeMap {
					err := document.UptimeChange{Address: k, Uptime: float64(100*v) / float64(len(blocks)), Time: doneTime}.Insert()
					if err != nil {
						logger.Error("uptimeChange insert", logger.String("err", err.Error()))
					}
				}

				logger.Info("UptimeChangeVo task end")
			}
		}
	})
}

func (task UpTimeChangeTask) init() {

	powerChanges, err := document.PowerChange{}.QueryPowerChangeList()

	if err != nil {
		logger.Error("query power change list", logger.String("err", err.Error()))
	}

	for _, powerChange := range powerChanges {
		if powerChange.Change != types.Slash {
			validatorMap[powerChange.Address] = model.ValVotingPowerChangeVo{Address: powerChange.Address, Time: powerChange.Time, Power: powerChange.Power, Change: powerChange.Change}
		}
	}
}
