package model

import (
	"github.com/irisnet/irishub-sync/store"
	"time"
)

type AccountResp struct {
	Address         string      `bson:"address"`
	Amount          store.Coins `bson:"amount"`
	Time            time.Time   `bson:"time"`
	Height          int64       `bson:"height"`
	WithdrawAddress string
}
