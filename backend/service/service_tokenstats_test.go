package service

import (
	"testing"
	"encoding/json"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
)

func TestTokenStatsService_QueryTokensAccountTotal(t *testing.T) {
	res, err := (&TokenStatsService{}).QueryTokensAccountTotal()
	if err != nil {
		t.Fatal(res)
	}
	ret, err := json.Marshal(res)
	t.Log(string(ret))
}

func TestTokenStatsService_QueryTokenStats(t *testing.T) {
	res, err := (&TokenStatsService{}).QueryTokenStats()
	if err != nil {
		t.Fatal(res)
	}
	ret, err := json.Marshal(res)
	t.Log(string(ret))
}

func TestCountNTotalAmount(t *testing.T) {
	type args struct {
		start   int
		end     int
		total   float64
		account []document.Account
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test CountNTotalAmount",
			args: args{
				start: 0,
				end:   4,
				total: 100,
				account:[]document.Account{
					{Total:utils.Coin{Amount:11.45,Denom:"iris-atto"}},
					{Total:utils.Coin{Amount:11.55,Denom:"iris-atto"}},
					{Total:utils.Coin{Amount:11.45,Denom:"iris-atto"}},
					{Total:utils.Coin{Amount:11.55,Denom:"iris-atto"}},
					{Total:utils.Coin{Amount:11.00,Denom:"iris-atto"}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CountNTotalAmount(tt.args.start, tt.args.end, tt.args.total, tt.args.account)
			resBytes, _ := json.MarshalIndent(res, "", "\t")
			t.Log(string(resBytes))
		})
	}
}
