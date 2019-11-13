package vo

type HtlcInfo struct {
	HashLock           string    `json:"hash_lock"`
	From               string    `json:"from"`
	To                 string    `json:"to"`
	ToMoniker          string    `json:"to_moniker"`
	FromMoniker        string    `json:"from_moniker"`
	Amount             []*CoinVo `json:"amount"`
	TimeLock           int64     `json:"time_lock"`
	Timestamp          int64     `json:"timestamp"`
	CrossChainReceiver string    `json:"cross_chain_receiver"`
	ExpireHeight       int64     `json:"expire_height"`
	State              string    `json:"state"`
}
