package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/irisnet/explorer/server/utils"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Account struct {
	Address string    `bson:"address"`
	Amount  []Coins   `bson:"amount"`
	Time    time.Time `bson:"time"`
	Height  int64     `bson:"height"`
}

type Coins struct {
	Denom  string `bson:"denom"`
	Amount int64  `bson:"amount"`
}

// mux.Router registrars
func RegisterAccount(r *mux.Router) error {
	funs := []func(*mux.Router) error{
		RegisterQueryAccount,
		RegisterQueryAllAccount,
	}

	for _, fn := range funs {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

func queryAccount(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	address := args["address"]

	c := utils.GetDatabase().C("account")
	defer c.Database.Session.Close()
	var result Account
	err := c.Find(bson.M{"address": address}).Sort("-amount.amount").One(&result)
	if err == nil {
		resultByte, _ := json.Marshal(result)
		w.Write(resultByte)
	}
}

func queryAccounts(w http.ResponseWriter, r *http.Request) {
	var data []Account
	w.Write(utils.QueryList("account", &data, r))
}

func RegisterQueryAccount(r *mux.Router) error {
	r.HandleFunc("/account/{address}", queryAccount).Methods("GET")
	return nil
}

func RegisterQueryAllAccount(r *mux.Router) error {
	r.HandleFunc("/accounts/{page}/{size}", queryAccounts).Methods("GET")
	return nil
}
