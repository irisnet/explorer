package model

import (
	"github.com/irisnet/explorer/backend/orm/document"
)

type AccountVo struct {
	Amount          document.Coins `json:"amount"`
	WithdrawAddress string         `json:"withdrawAddress"`
	Address         string         `json:"address"`
	Deposits        document.Coin  `json:"deposits"`
	IsProfiler      bool           `json:"isProfiler"`
}

type ValAccVo struct {
	AccountVo
	ValProfile
}

type AccountInfo struct {
	Address string          `json:"address"`
	Balance []document.Coin `json:"balance"`
	Percent float64         `json:"percent"`
}
