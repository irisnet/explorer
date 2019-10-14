package vo

type BondedTokensVo struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	VotingPower     int64  `json:"voting_power,string"`
	OperatorAddress string `json:"operator_address"`
	OwnerAddress    string `json:"owner_address"`
	Icons           string `json:"icons"`
}

type BondedTokensRespond []*BondedTokensVo
