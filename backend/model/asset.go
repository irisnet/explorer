package model

import "fmt"

type AssetTokens struct {
	Symbol  string `json:"symbol"`
	Decimal int    `json:"decimal"`
}

func (b AssetTokens) String() string {

	return fmt.Sprintf(`
		Symbol          :%v
		Decimal         :%v
		`, b.Symbol, b.Decimal)
}
