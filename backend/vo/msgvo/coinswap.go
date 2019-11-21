package msgvo

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/utils"
)

type TxMsgAddLiquidity struct {
	MaxToken     utils.Coin `json:"max_token"`      // coin to be deposited as liquidity with an upper bound for its amount
	ExactIrisAmt string     `json:"exact_iris_amt"` // exact amount of native asset being add to the liquidity pool
	MinLiquidity string     `json:"min_liquidity"`  // lower bound UNI sender is willing to accept for deposited coins
	Deadline     int64      `json:"deadline"`
	Sender       string     `json:"sender"`
}

type TxMsgRemoveLiquidity struct {
	MinToken          string     `json:"min_token"`          // coin to be withdrawn with a lower bound for its amount
	WithdrawLiquidity utils.Coin `json:"withdraw_liquidity"` // amount of UNI to be burned to withdraw liquidity from a reserve pool
	MinIrisAmt        string     `json:"min_iris_amt"`       // minimum amount of the native asset the sender is willing to accept
	Deadline          int64      `json:"deadline"`
	Sender            string     `json:"sender"`
}

type TxMsgSwapOrder struct {
	Input      Input  `json:"input"`        // the amount the sender is trading
	Output     Output `json:"output"`       // the amount the sender is receiving
	Deadline   int64  `json:"deadline"`     // deadline for the transaction to still be considered valid
	IsBuyOrder bool   `json:"is_buy_order"` // boolean indicating whether the order should be treated as a buy or sell
}

type Input struct {
	Address string     `json:"address"`
	Coin    utils.Coin `json:"coin"`
}

type Output struct {
	Address string     `json:"address"`
	Coin    utils.Coin `json:"coin"`
}

func (vo *TxMsgAddLiquidity) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgRemoveLiquidity) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

func (vo *TxMsgSwapOrder) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
