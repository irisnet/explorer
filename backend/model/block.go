package model

import (
	"github.com/irisnet/irishub-sync/store/document"
	"time"
)

type BlockInfoVo struct {
	Height     int64     `json:"height"`
	Hash       string    `json:"hash"`
	Time       time.Time `json:"time"`
	NumTxs     int64     `json:"num_txs"`
	Validators []ValInfo `json:"validators"`
	LastCommit []string  `json:"last_commit"`
	TotalTxs   int64     `json:"total_txs"`
}

type ValInfo struct {
	Address     string `json:"address"`
	VotingPower int64  `json:"voting_power"`
}

type BlockVo struct {
	CandidateMap map[string]string
	Counter      CounterVo
	document.Block
}

type CounterVo struct {
	TxCnt    int
	PropoCnt int
}
