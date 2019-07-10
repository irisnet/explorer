package model

type BondedTokensVo struct {
	Moniker         string `json:"moniker,omitempty"`
	Identity        string `json:"identity,omitempty"`
	VotingPower     string `json:"voting_power,omitempty"`
	OperatorAddress string `json:"operator_address,omitempty"`
}
