package document

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionName = "validator_up_time"

	ValidatorUpTime_Field_ValAddress = "val_address"
	ValidatorUpTime_Field_UpTime     = "up_time"
)

type ValidatorUpTime struct {
	ValAddress string  `bson:"val_address"`
	UpTime     float64 `bson:"up_time"`
}

func (d ValidatorUpTime) Name() string {
	return CollectionName
}

func (d ValidatorUpTime) PkKvPair() map[string]interface{} {
	return bson.M{ValidatorUpTime_Field_ValAddress: d.ValAddress}
}
