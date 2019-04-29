package model

import (
	"time"

	"github.com/irisnet/explorer/backend/orm/document"
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

type Block struct {
	BlockInfo    BlockInfo        `json:"block_info"`
	ValidatorSet ValidatorSet     `json:"validator_set"`
	TokenFlows   TokenFlows       `json:"token_flows"`
	Proposals    []ProposalInfoVo `json:"proposals_info"`
}

type BlockInfo struct {
	BlockHeight          string         `json:"block_height"`
	Timestamp            time.Time      `json:"timestamp"`
	BlockHash            string         `json:"block_hash"`
	Transactions         string         `json:"transactions"`
	PropopserMoniker     string         `json:"propopser_moniker"`
	PropoperAddr         string         `json:"propopser_addr"`
	Rewards              document.Coins `json:"rewards"`
	LastBlock            int64          `json:"last_block"`
	LastBlockHash        string         `json:"last_block_hash"`
	VoteValidatorNum     int            `json:"vote_validator_num"`
	ValidatorNum         int            `json:"validator_num"`
	PrecommitVotingPower int            `json:"precommit_voting_power"`
	TotalVotingPower     int            `json:"total_voting_power"`
	LatestHeight         string         `json:"latest_height"`
}
