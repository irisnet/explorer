package model

import (
	"github.com/irisnet/explorer/backend/utils"
)

type AccountVo struct {
	Amount          utils.Coins `json:"amount"`
	WithdrawAddress string      `json:"withdrawAddress"`
	Address         string      `json:"address"`
	Deposits        utils.Coin  `json:"deposits"`
	IsProfiler      bool        `json:"isProfiler"`
}

type ValAccVo struct {
	AccountVo
	ValProfile
}

type AccountInfo struct {
	Rank     int          `json:"rank"`
	Address  string       `json:"address"`
	Balance  []utils.Coin `json:"balance"`
	Percent  float64      `json:"percent"`
	UpdateAt int64        `json:"update_at"`
}
