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
