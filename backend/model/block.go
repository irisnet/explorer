package model

import (
	"time"
)

type BlockInfoVo struct {
	Height        int64     `json:"height,omitempty"`
	Hash          string    `json:"hash,omitempty"`
	Time          time.Time `json:"time,omitempty"`
	NumTxs        int64     `json:"num_txs,omitempty"`
	Validators    []ValInfo `json:"validators,omitempty"`
	LastCommit    []string  `json:"last_commit,omitempty"`
	TotalTxs      int64     `json:"total_txs,omitempty"`
	LastBlockHash string    `json:"last_block_hash,omitempty"`
}

type ValInfo struct {
	Address     string `json:"address"`
	VotingPower int64  `json:"voting_power"`
}
