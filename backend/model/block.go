package model

import (
	"time"

	"github.com/irisnet/explorer/backend/orm/document"
)

type BlockInfoVo struct {
	Height        int64     `json:"height"`
	Hash          string    `json:"hash,omitempty"`
	Time          time.Time `json:"time"`
	NumTxs        int64     `json:"num_txs"`
	Validators    []ValInfo `json:"validators,omitempty"`
	LastCommit    []string  `json:"last_commit,omitempty"`
	TotalTxs      int64     `json:"total_txs,omitempty"`
	LastBlockHash string    `json:"last_block_hash,omitempty"`
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
	IsProposer       bool   `json:"is_proposer"`
}

type ValidatorSet struct {
	Total int              `json:"total"`
	Items []BlockValidator `json:"items"`
}

type Block struct {
	BlockInfo    BlockInfo    `json:"block_info"`
	ValidatorSet ValidatorSet `json:"validator_set"`
	TxPage       TxPage       `json:"tx_page"`
	ProposalPage ProposalPage `json:"proposals_page"`
}

type BlockInfo struct {
	BlockHeight           string        `json:"block_height"`
	Timestamp             time.Time     `json:"timestamp"`
	BlockHash             string        `json:"block_hash"`
	Transactions          string        `json:"transactions"`
	PropopserMoniker      string        `json:"propopser_moniker"`
	PropoperAddr          string        `json:"propopser_addr"`
	PrecommitValidatorNum interface{}   `json:"precommit_validator_num"`
	TotalValidatorNum     int           `json:"total_validator_num"`
	PrecommitVotingPower  interface{}   `json:"precommit_voting_power"`
	TotalVotingPower      int           `json:"total_voting_power"`
	LatestHeight          string        `json:"latest_height"`
	MintCoin              document.Coin `json:"mint_coin"`
}
