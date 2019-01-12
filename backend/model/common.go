package model

import "time"

type ResultVo struct {
	Type string      `json:"Type,omitempty"`
	Data interface{} `json:"Data,omitempty"`
}

type SimpleBlockVo struct {
	Height    int64
	Timestamp time.Time
	Hash      string
}

type SimpleProposalVo struct {
	ProposalId uint64
	Title      string
	Type       string
	Status     string
	SubmitTime time.Time
}
