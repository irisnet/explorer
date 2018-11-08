package rest

import (
	"encoding/json"
	"github.com/irisnet/explorer/backend/types"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irisnet/explorer/backend/utils"
	"github.com/irisnet/irishub-sync/store/document"
	"gopkg.in/mgo.v2/bson"
)

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
	var result document.Account
	err := c.Find(bson.M{"address": address}).Sort("-amount.amount").One(&result)
	if err != nil {
		w.Write([]byte("not find account"))
		return
	}

	balance := utils.GetBalance(result.Address)
	if len(balance) >= 0 {
		result.Amount = balance
	}
	resultByte, _ := json.Marshal(result)
	w.Write(resultByte)
}

func queryAccounts(w http.ResponseWriter, r *http.Request) {
	var data []document.Account
	w.Write(utils.QueryList("account", &data, nil, "-time", r))
}

func RegisterQueryAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryAccount, queryAccount).Methods("GET")
	return nil
}

func RegisterQueryAllAccount(r *mux.Router) error {
	r.HandleFunc(types.UrlRegisterQueryAllAccount, queryAccounts).Methods("GET")
	return nil
}
