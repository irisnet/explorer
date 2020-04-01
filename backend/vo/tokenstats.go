package vo

type TokenStatsVo struct {
	DelegatedTokens   *CoinVo `json:"delegated_tokens"`
	BurnedTokens      *CoinVo `json:"burned_tokens"`
	CirculationTokens *CoinVo `json:"circulation_tokens"`
	TotalsupplyTokens *CoinVo `json:"totalsupply_tokens"`
	//InitsupplyTokens  *CoinVo `json:"initsupply_tokens"`
	CommunityTax     *CoinVo `json:"community_tax"`
	FoundationBonded *CoinVo `json:"foundation_bonded"`
}

type CoinVo struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type TokenStatsSegment struct {
	TotalAmount *CoinVo `json:"total_amount"`
	Percent     float64 `json:"percent"`
}
