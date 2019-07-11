package document

import (
	"fmt"

	"github.com/irisnet/explorer/backend/orm"
)

const CollectionNmUptimeChange = "ex_uptime_change"

type UptimeChange struct {
	Address string
	Time    string
	Uptime  float64
}

func (uptime UptimeChange) String() string {
	return fmt.Sprintf(`
		Address :%v
		Time    :%v
		Uptime  :%v
		`, uptime.Address, uptime.Time, uptime.Uptime)
}

func (_ UptimeChange) QueryOne() (UptimeChange, error) {

	db := orm.GetDatabase()
	u := db.C(CollectionNmUptimeChange)

	defer db.Session.Close()

	var uptimeChange UptimeChange
	err := u.Find(nil).Sort("-time").One(&uptimeChange)

	return uptimeChange, err
}

func (u UptimeChange) Insert() error {
	db := orm.GetDatabase()
	uCollection := db.C(CollectionNmUptimeChange)
	defer db.Session.Close()

	return uCollection.Insert(&u)
}
