package model

import (
	"time"
)

type BlockInfoVo struct {
	Height        int64     `json:"height"`
	Hash          string    `json:"hash"`
	Time          time.Time `json:"time"`
	NumTxs        int64     `json:"num_txs"`
	Validators    []ValInfo `json:"validators"`
	LastCommit    []string  `json:"last_commit"`
	TotalTxs      int64     `json:"total_txs"`
	LastBlockHash string    `json:"last_block_hash"`
}

type ValInfo struct {
	Address     string `json:"address"`
	VotingPower int64  `json:"voting_power"`
}

type BlockValidator struct {
	Moniker          string `json:"moniker"`
	OperatorAddress  string `json:"operator_address"`
	Consensus        string `json:"consensus"`
	ProposerPriority string `json:"proposer_priority"`
	VotingPower      string `json:"voting_power"`
}

type ValidatorSet struct {
	Total int              `json:"total"`
	Items []BlockValidator `json:"items"`
}
