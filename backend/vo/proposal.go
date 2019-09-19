package vo

import (
	"time"

	"github.com/irisnet/explorer/backend/utils"
)

type ValidatorVote struct {
	Title      string `json:"title"`
	ProposalId uint64 `json:"proposal_id"`
	Status     string `json:"status"`
	Voted      string `json:"voted"`
	TxHash     string `json:"tx_hash"`
}

type ValidatorVotePage struct {
	Total int             `json:"total"`
	Items []ValidatorVote `json:"items"`
}

type ValidatorDepositTxPage struct {
	Total int                  `json:"total"`
	Items []ValidatorDepositTx `json:"items"`
}

type ValidatorDepositTx struct {
	ProposalId      uint64      `json:"proposal_id"`
	Proposer        string      `json:"proposer"`
	DepositedAmount utils.Coins `json:"deposited_amount"`
	Submited        bool        `json:"submited"`
	TxHash          string      `json:"tx_hash"`
	Moniker         string      `json:"moniker"`
}

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
	Parameters      []Param     `json:"parameters"`
	Version         uint64      `json:"version"`
	Software        string      `json:"software"`
	SwitchHeight    uint64      `json:"switch_height"`
	Threshold       string      `json:"threshold"`
	Level           string      `json:"level"`
	YesThreshold    string      `json:"yes_threshold"`
	VetoThreshold   string      `json:"veto_threshold"`
	Participation   string      `json:"participation"`
	Penalty         string      `json:"penalty"`
	Usage           string      `json:"usage"`
	BurnPercent     float32     `json:"burn_percent"`
}

type Vote struct {
	Voter  string    `json:"voter"`
	Option string    `json:"option"`
	Time   time.Time `json:"time"`
}

type ProposalInfoVo struct {
	Proposal Proposal `json:"proposal"`
	Votes    []Vote   `json:"votes"`
}

type FinalVotes struct {
	Yes               string `json:"yes,omitempty"`
	No                string `json:"no,omitempty"`
	NoWithVeto        string `json:"no_with_veto,omitempty"`
	Abstain           string `json:"abstain,omitempty"`
	SystemVotingPower string `json:"system_voting_power,omitempty"`
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
		TotalVotingPower float64             `json:"voting_power_for_height"`
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
		Voter          string    `json:"voter"`
		VoterMoniker   string    `json:"voter_moniker,omitempty"`
		Option         string    `json:"option"`
		VotingPower    float64   `json:"voting_power"`
		Time           time.Time `json:"time"`
		DelVotingPower float64   `json:"del_voting_power"`
		ValVotingPower float64   `json:"val_voting_power"`
	}

	DelegatorGovInfo struct {
		Address        string    `json:"address"`
		Option         string    `json:"option"`
		Moniker        string    `json:"moniker"`
		DelVotingPower float64   `json:"del_voting_power"`
		ValVotingPower float64   `json:"val_voting_power"`
		IsValidator    bool      `json:"is_validator"`
		ValAddr        string    `json:"val_addr"`
		Time           time.Time `json:"time"`
	}

	ValidatorGovInfo struct {
		Address            string  `json:"address"`
		Tokens             float64 `json:"token"`
		DelShares          float64 `json:"del_shares"`
		DelDeductionShares float64 `json:"del_deduction_shares"`
	}
)

type GetVoteTxResponse struct {
	Total int       `json:"total"`
	Items []VoteTx  `json:"items"`
	Stats VoteStats `json:"stats"`
}

type VoteTx struct {
	Voter     string    `json:"voter"`
	Moniker   string    `json:"moniker"`
	Option    string    `json:"option"`
	TxHash    string    `json:"tx_hash"`
	Timestamp time.Time `json:"timestamp"`
	Height    int64     `json:"height"`
}

type LookupIcons struct {
	Status struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"status"`
	Them []Them `json:"them"`
}

type Them struct {
	Id       string `json:"id"`
	Pictures struct {
		Primary struct {
			Url    string `json:"url"`
			Source string `json:"source"`
		} `json:"primary"`
	} `json:"pictures"`
}

type VoteStats struct {
	All        int64 `json:"all"`
	Validator  int64 `json:"validator"`
	Delegator  int64 `json:"delegator"`
	Yes        int64 `json:"yes"`
	No         int64 `json:"no"`
	NoWithVeto int64 `json:"no_with_veto"`
	Abstain    int64 `json:"abstain"`
}
