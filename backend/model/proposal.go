package model

import (
	"time"

	"github.com/irisnet/explorer/backend/orm/document"
)

type Proposal struct {
	Title           string         `json:"title"`
	ProposalId      uint64         `json:"proposal_id"`
	Type            string         `json:"type"`
	Description     string         `json:"description"`
	Status          string         `json:"status"`
	SubmitTime      time.Time      `json:"submit_time"`
	DepositEndTime  time.Time      `json:"deposit_end_time"`
	VotingStartTime time.Time      `json:"voting_start_time"`
	VotingEndTime   time.Time      `json:"voting_end_time"`
	TotalDeposit    document.Coins `json:"total_deposit"`
	Proposer        string         `json:"proposer"`
	TxHash          string         `json:"tx_hash"`
	Parameters      []Param        `json:"parameters,omitempty"`
	Version         uint64         `json:"version,omitempty"`
	Software        string         `json:"software,omitempty"`
	SwitchHeight    uint64         `json:"switch_height,omitempty"`
	Threshold       string         `json:"threshold,omitempty"`
}

type Vote struct {
	Voter  string    `json:"voter"`
	Option string    `json:"option"`
	Time   time.Time `json:"time"`
}

type ProposalInfoVo struct {
	Proposal Proposal   `json:"proposal"`
	Votes    []Vote     `json:"votes"`
	Result   VoteResult `json:"result"`
}

type VoteResult struct {
	Yes        int
	No         int
	NoWithVeto int
	Abstain    int
}
