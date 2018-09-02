package types

import "time"

var (
	Change  = "powerChange"
	Slash   = "slash"
	Recover = "recover"
)

type PowerChangeOrder struct {
	Address string `bson:"_id,omitempty"`
	Power   int64
	Time    time.Time
	Change  string
}

type PowerChange struct {
	Address string
	Power   int64
	Time    time.Time
	Change  string
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
