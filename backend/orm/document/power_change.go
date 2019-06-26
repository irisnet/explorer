package document

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
)

const CollectionNmPowerChange = "ex_power_change"

type PowerChange struct {
	Height  int64
	Address string
	Power   int64
	Time    time.Time
	Change  string
}

func (power PowerChange) String() string {

	return fmt.Sprintf(`
		Height  :%v
		Address :%v
		Power   :%v
		Time    :%v
		Change  :%v
		`, power.Height, power.Address, power.Power, power.Time, power.Change)
}

func (p PowerChange) Insert() error {

	db := orm.GetDatabase()
	defer db.Session.Close()
	powerChangeCollection := db.C(CollectionNmPowerChange)
	return powerChangeCollection.Insert(&p)
}

func (p PowerChange) QueryPowerChangeList() ([]PowerChange, error) {

	powerChanges := []PowerChange{}
	db := orm.GetDatabase()
	pCollection := db.C("power_change")
	pipe := pCollection.Pipe(
		[]bson.M{
			{"$group": bson.M{
				"_id":    "$address",
				"power":  bson.M{"$last": "$power"},
				"date":   bson.M{"$last": "$date"},
				"change": bson.M{"$last": "$change"},
			}},
		},
	)
	err := pipe.All(&powerChanges)
	return powerChanges, err
}
