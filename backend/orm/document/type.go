package document

type Coin struct {
	Denom  string  `json:"denom"`
	Amount float64 `json:"amount"`
}

type Coins []Coin

type Fee struct {
	Amount Coins `json:"amount"`
	Gas    int64 `json:"gas"`
}

type ActualFee struct {
	Denom  string  `json:"denom"`
	Amount float64 `json:"amount"`
}
