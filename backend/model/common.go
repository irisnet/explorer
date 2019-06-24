package model

import "time"

type ResultVo struct {
	Block    SimpleBlockVo    `json:"block"`
	Proposal SimpleProposalVo `json:"proposal"`
}

type SimpleBlockVo struct {
	Height    int64     `json:"height"`
	Timestamp time.Time `json:"timestamp"`
	Hash      string    `json:"hash"`
}

type SimpleProposalVo struct {
	ProposalId uint64    `json:"proposal_id"`
	Title      string    `json:"title"`
	Type       string    `json:"type"`
	Status     string    `json:"status"`
	SubmitTime time.Time `json:"submit_time"`
}
