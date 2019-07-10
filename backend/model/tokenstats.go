package model

type TokenStatsVo struct {
	DelegatedTokens   *Coin `json:"delegated_tokens"`
	BurnedTokens      *Coin `json:"burned_tokens"`
	CirculationTokens *Coin `json:"circulation_tokens"`
	TotalsupplyTokens *Coin `json:"totalsupply_tokens"`
	InitsupplyTokens  *Coin `json:"initsupply_tokens"`
}

