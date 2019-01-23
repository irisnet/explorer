package model

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"time"
)

type AccountVo struct {
	Address         string         `bson:"address"`
	Amount          document.Coins `bson:"amount"`
	Time            time.Time      `bson:"time"`
	Height          int64          `bson:"height"`
	WithdrawAddress string         `bson:"withdrawAddress"`
	Deposits        float64        `bson:"deposits"`
}
