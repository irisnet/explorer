package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/irisnet/explorer/backend/lcd"
	"github.com/irisnet/explorer/backend/orm/document"
	"github.com/irisnet/explorer/backend/utils"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

var genesis GenesisDoc
var exportHeight = int64(0)
var vMap = make(map[string]float64)

func TestCountFromGenesisFile(t *testing.T) {
	var file = "/Users/zhangzhiqiang/Downloads/genesis.json"
	buf, _ := readFile(file)

	if err := json.Unmarshal(buf, &genesis); err != nil {
		panic(err)
	}

	buildRate(genesis.AppState.StakeData.Validators)

	var totalDelegationAmt = document.Coin{Denom: "iris-atto", Amount: 0}
	var totalUnbondingAmt = document.Coin{Denom: "iris-atto", Amount: 0}
	for _, d := range genesis.AppState.StakeData.Bonds {
		rate := vMap[d.ValidatorAddr]

		shares, err := strconv.ParseFloat(d.Shares, 0)
		if err != nil {
			panic(err)
		}
		amount := rate * shares
		totalDelegationAmt = totalDelegationAmt.Add(
			document.Coin{
				Amount: amount,
				Denom:  "iris-atto",
			})
	}

	for _, d := range genesis.AppState.StakeData.UnbondingDelegations {
		var amt, err = strconv.ParseFloat(d.Balance.Amount, 0)
		if err != nil {
			panic(err)
		}
		totalUnbondingAmt = totalUnbondingAmt.Add(document.Coin{
			Denom:  d.Balance.Denom,
			Amount: amt,
		})
	}

	fmt.Println(fmt.Sprintf("total Delegate:%v,total unbonding:%v", totalDelegationAmt, totalUnbondingAmt))

}

func buildRate(validators []lcd.ValidatorVo) {
	for _, v := range genesis.AppState.StakeData.Validators {
		token, err := strconv.ParseFloat(v.Tokens, 0)
		if err != nil {
			panic(err)
		}

		shares, err := strconv.ParseFloat(v.DelegatorShares, 0)
		if err != nil {
			panic(err)
		}

		vMap[v.OperatorAddress] = token / shares
	}
}

func TestInitAccountFromGenesisFile(t *testing.T) {
	var file = "/Users/zhangzhiqiang/Downloads/genesis.json"
	buf, _ := readFile(file)

	if err := json.Unmarshal(buf, &genesis); err != nil {
		panic(err)
	}

	buildRate(genesis.AppState.StakeData.Validators)

	var ops []txn.Op
	for _, acc := range genesis.AppState.Accounts {
		coin := utils.ParseCoin(acc.Coins[0])
		accountNumber, _ := utils.ParseUint(acc.AccountNumber)

		delegateAmt, delegateHeight := getDelegateAmt(acc.Address)
		if delegateHeight == 0 {
			delegateHeight = exportHeight
		}
		unbondingAmt, undHeight := getUnbondingAmt(acc.Address)
		if undHeight == 0 {
			undHeight = exportHeight
		}
		totalAmt := coin.Add(delegateAmt).Add(unbondingAmt)

		ops = append(ops, txn.Op{
			C:  document.CollectionNmAccount,
			Id: bson.NewObjectId(),
			Insert: document.Account{
				Address:           acc.Address,
				Total:             totalAmt,
				TotalUpdateHeight: exportHeight,
				TotalUpdateAt:     genesis.GenesisTime.Unix(),

				CoinIris:             coin,
				CoinIrisUpdateHeight: exportHeight,
				CoinIrisUpdateAt:     genesis.GenesisTime.Unix(),

				Delegation:             delegateAmt,
				DelegationUpdateHeight: delegateHeight,
				DelegationUpdateAt:     genesis.GenesisTime.Unix(),

				UnbondingDelegation:             unbondingAmt,
				UnbondingDelegationUpdateHeight: undHeight,
				UnbondingDelegationUpdateAt:     genesis.GenesisTime.Unix(),

				AccountNumber: accountNumber,
			},
		})
	}

	err := document.Account{}.Batch(ops)

	if err != nil {
		fmt.Println(err)
	}
}

func getDelegateAmt(address string) (document.Coin, int64) {
	var total = document.Coin{
		Amount: 0,
		Denom:  "iris-atto",
	}
	var updateHeight = int64(0)
	for _, d := range genesis.AppState.StakeData.Bonds {
		if d.DelegatorAddr == address {
			rate := vMap[d.ValidatorAddr]

			shares, err := strconv.ParseFloat(d.Shares, 0)
			if err != nil {
				panic(err)
			}
			amount := rate * shares
			total = total.Add(
				document.Coin{
					Amount: amount,
					Denom:  "iris-atto",
				})

			height, _ := strconv.ParseInt(d.Height, 10, 0)
			if height > updateHeight {
				updateHeight = height
			}
		}
	}
	return total, updateHeight
}

func getUnbondingAmt(address string) (document.Coin, int64) {
	var total = document.Coin{
		Amount: 0,
		Denom:  "iris-atto",
	}
	var updateHeight = int64(0)
	for _, d := range genesis.AppState.StakeData.UnbondingDelegations {
		if d.DelegatorAddr == address {
			var amt, err = strconv.ParseFloat(d.Balance.Amount, 0)
			if err != nil {
				panic(err)
			}
			total = total.Add(document.Coin{
				Denom:  d.Balance.Denom,
				Amount: amt,
			})

			height, _ := strconv.ParseInt(d.CreationHeight, 10, 0)
			if height > updateHeight {
				updateHeight = height
			}
		}
	}
	return total, updateHeight
}

type GenesisDoc struct {
	GenesisTime time.Time `json:"genesis_time"`
	ChainID     string    `json:"chain_id"`
	AppState    struct {
		Accounts  []AccountStake `json:"accounts"`
		StakeData StakeState     `json:"stake"`
	} `json:"app_state"`
}

type AccountStake struct {
	Address       string   `json:"address"`
	AccountNumber string   `json:"account_number"`
	Coins         []string `json:"coins"`
}

type StakeState struct {
	Validators           []lcd.ValidatorVo     `json:"validators"`
	Bonds                []Delegation          `json:"bonds"`
	UnbondingDelegations []UnbondingDelegation `json:"unbonding_delegations"`
}

type Delegation struct {
	DelegatorAddr string `json:"delegator_addr"`
	ValidatorAddr string `json:"validator_addr"`
	Shares        string `json:"shares"`
	Height        string `json:"height"`
}

type UnbondingDelegation struct {
	DelegatorAddr  string `json:"delegator_addr"` // delegator
	ValidatorAddr  string `json:"validator_addr"` // validator unbonding from operator addr
	Balance        Coin   `json:"balance"`        // atoms to receive at completion
	CreationHeight string `json:"creation_height"`
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

func readFile(filename string) (content []byte, err error) {
	fp, err := os.Open(filename) // 获取文件指针
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	fileInfo, err := fp.Stat()
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, fileInfo.Size())
	_, err = fp.Read(buffer) // 文件内容读取到buffer中
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
