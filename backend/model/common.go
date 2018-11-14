package model

import "time"

type ResultVo struct {
	Type string      `json:"Type,omitempty"`
	Data interface{} `json:"Data,omitempty"`
}

type SimpleBlock struct {
	Height    int64
	Timestamp time.Time
	Hash      string
}

type SimpleProposal struct {
	ProposalId int64
	Title      string
	Type       string
	Status     string
	SubmitTime time.Time
}
