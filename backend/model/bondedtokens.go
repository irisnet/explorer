package model

type BondedTokensVo struct {
	Moniker         string `json:"moniker,omitempty"`
	Identity        string `json:"identity,omitempty"`
	VotingPower     int64  `json:"voting_power,string,omitempty"`
	OperatorAddress string `json:"operator_address,omitempty"`
}
