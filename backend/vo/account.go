package vo

import (
	"fmt"

	"github.com/irisnet/explorer/backend/utils"
)

type AccountVo struct {
	Amount          utils.Coins `json:"amount"`
	WithdrawAddress string      `json:"withdrawAddress"`
	Address         string      `json:"address"`
	Deposits        utils.Coin  `json:"deposits"`
	IsProfiler      bool        `json:"isProfiler"`
}

func (a AccountVo) String() string {
	return fmt.Sprintf(`
		Amount          :%v
		WithdrawAddress :%v
		Address         :%v
		Deposits        :%v
		IsProfiler      :%v
		`, a.Amount, a.WithdrawAddress, a.Address, a.Deposits, a.IsProfiler)
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

func (a AccountInfo) String() string {
	return fmt.Sprintf(`
		Rank     :%v
		Address  :%v
		Balance  :%v
		Percent  :%v
		UpdateAt :%v
		`, a.Rank, a.Address, a.Balance, a.Percent, a.UpdateAt)
}
