package model

import (
	"github.com/irisnet/explorer/backend/orm/document"
	"time"
)

type CandidateStatus struct {
	Uptime         int
	TotalBlock     int
	PrecommitCount float64
}

type CandidatesTopNVo struct {
	PowerAll   float64
	Candidates []CandidateAll
}

type CandidatesVo struct {
	Count      int
	PowerAll   float64
	Candidates []CandidateAll
}

type CandidateAll struct {
	document.Candidate
	CandidateStatus
}

type CandidatesInfoVo struct {
	PowerAll float64
	document.Candidate
}

type ChainStatusVo struct {
	ValidatorsCount int
	TxCount         int
	VotingPower     float64
	Tps             float64
}

type UptimeChangeVo struct {
	Address string
	Time    string
	Uptime  float64
}

type CandidateUpTimeVo struct {
	Time   string `bson:"_id,omitempty"`
	Uptime float64
}

type ValVotingPowerChangeInfo struct {
	Address string `bson:"_id,omitempty"`
	Power   int64
	Time    time.Time
	Change  string
}

type ValVotingPowerChangeVo struct {
	Height  int64
	Address string
	Power   int64
	Time    time.Time
	Change  string
}
