package model

import (
	"time"

	"github.com/irisnet/explorer/backend/utils"
)

type Proposal struct {
	Title           string      `json:"title"`
	ProposalId      uint64      `json:"proposal_id"`
	Type            string      `json:"type"`
	Description     string      `json:"description"`
	Status          string      `json:"status"`
	SubmitTime      time.Time   `json:"submit_time"`
	DepositEndTime  time.Time   `json:"deposit_end_time"`
	VotingStartTime time.Time   `json:"voting_start_time"`
	VotingEndTime   time.Time   `json:"voting_end_time"`
	TotalDeposit    utils.Coins `json:"total_deposit"`
	Proposer        string      `json:"proposer"`
	TxHash          string      `json:"tx_hash"`
	Parameters      []Param     `json:"parameters,omitempty"`
	Version         uint64      `json:"version,omitempty"`
	Software        string      `json:"software,omitempty"`
	SwitchHeight    uint64      `json:"switch_height,omitempty"`
	Threshold       string      `json:"threshold,omitempty"`
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

type FinalVotes struct {
	Yes        string `json:"yes,omitempty"`
	No         string `json:"no,omitempty"`
	NoWithVeto string `json:"no_with_veto,omitempty"`
	Abstain    string `json:"abstain,omitempty"`
}

type (
	ProposalNewStyle struct {
		ProposalId       uint64              `json:"proposal_id"`
		Title            string              `json:"title"`
		Type             string              `json:"type"`
		Status           string              `json:"status"`
		Level            Level               `json:"level,omitempty"`
		InitialDeposit   Coin                `json:"intial_deposit,omitempty"`
		TotalDeposit     Coin                `json:"total_deposit,omitempty"`
		Votes            []VoteWithVoterInfo `json:"votes"`
		TotalVotingPower int64               `json:"voting_power_for_height"`
		SubmitTime       time.Time           `json:"submit_time,omitempty"`
		DepositEndTime   time.Time           `json:"deposit_end_time,omitempty"`
		VotingEndTime    time.Time           `json:"voting_end_time,omitempty"`
		FinalVotes       FinalVotes          `json:"final_votes"`
	}

	Level struct {
		Name     string   `json:"name"`
		GovParam GovParam `json:"gov_param,omitempty"`
	}

	GovParam struct {
		Participation string `json:"participation,omitempty"`
		PassThreshold string `json:"pass_threshold,omitempty"`
		VetoThreshold string `json:"veto_threshold,omitempty"`
		MinDeposit    Coin   `json:"min_deposit,omitempty"`
	}

	VoteWithVoterInfo struct {
		Voter        string    `json:"voter"`
		VoterMoniker string    `json:"voter_moniker,omitempty"`
		Option       string    `json:"option"`
		VotingPower  int64     `json:"voting_power"`
		Time         time.Time `json:"time"`
	}
)

type GetVoteTxResponse struct {
	Total int      `json:"total"`
	Items []VoteTx `json:"items"`
}

type VoteTx struct {
	Voter     string    `json:"voter"`
	Moniker   string    `json:"moniker"`
	Option    string    `json:"option"`
	TxHash    string    `json:"tx_hash"`
	Timestamp time.Time `json:"timestamp"`
}
