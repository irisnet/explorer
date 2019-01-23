package model

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"time"
)

type AccountVo struct {
	Address         string         `json:"address"`
	Amount          document.Coins `json:"amount"`
	Time            time.Time      `json:"time"`
	Height          int64          `json:"height"`
	WithdrawAddress string         `json:"withdrawAddress"`
	Deposits        document.Coin  `json:"deposits"`
}
