package model

import (
	"github.com/irisnet/irishub-sync/store/document"
	"time"
)

type CandidateStatus struct {
	Uptime         int
	TotalBlock     int
	PrecommitCount float64
}

type CandidatesTopN struct {
	PowerAll   float64
	Candidates []CandidateAll
}

type Candidates struct {
	Count      int
	PowerAll   float64
	Candidates []CandidateAll
}

type CandidateAll struct {
	document.Candidate
	CandidateStatus
}

type CandidateWithPower struct {
	PowerAll float64
	document.Candidate
}

type ChainStatus struct {
	ValidatorsCount int
	TxCount         int
	VotingPower     float64
	Tps             float64
}

type UptimeChange struct {
	Address string
	Time    string
	Uptime  float64
}

type CandidateUptime struct {
	Time   string `bson:"_id,omitempty"`
	Uptime float64
}

type PowerChangeOrder struct {
	Address string `bson:"_id,omitempty"`
	Power   int64
	Time    time.Time
	Change  string
}

type PowerChange struct {
	Height  int64
	Address string
	Power   int64
	Time    time.Time
	Change  string
}
