package document

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNmAccount = "account"

	Account_Field_Addres = "address"
	Account_Field_Amount = "amount"
	Account_Field_Time   = "time"
	Account_Field_Height = "height"
)

type Account struct {
	Address string    `bson:"address"`
	Amount  Coins     `bson:"amount"`
	Time    time.Time `bson:"time"`
	Height  int64     `bson:"height"`
}

func (a Account) Name() string {
	return CollectionNmAccount
}

func (a Account) PkKvPair() map[string]interface{} {
	return bson.M{Account_Field_Addres: a.Address}
}
