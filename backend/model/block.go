package model

import (
	"time"
)

type BlockInfoVo struct {
	Height     int64     `json:"height"`
	Hash       string    `json:"hash"`
	Time       time.Time `json:"time"`
	NumTxs     int64     `json:"num_txs"`
	Validators []ValInfo `json:"validators"`
	LastCommit []string  `json:"last_commit"`
}

type ValInfo struct {
	Address     string `json:"address"`
	VotingPower int64  `json:"voting_power"`
}
