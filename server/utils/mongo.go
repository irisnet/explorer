package utils

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/irisnet/explorer/server/types"
)

func init() {
	mongoUrl := GetEnv("MONGO_URL", "116.62.62.39:27117")
	url := fmt.Sprintf("mongodb://%s", mongoUrl)

	var err error
	session, err = mgo.Dial(url)
	if err != nil {
		log.Fatalln(err)
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(300)
}

var session *mgo.Session

func GetDatabase() *mgo.Database {
	// max session num is 4096
	return session.Copy().DB("sync-iris-dev")
}

func QueryList(collation string, data interface{}, m map[string]interface{}, sort string, r *http.Request) []byte {
	page, size := TransferPage(r)
	c := GetDatabase().C(collation)
	defer c.Database.Session.Close()
	count, err := c.Find(m).Count()
	if err != nil {
		resultByte, _ := json.Marshal(types.Page{Count: 0, Data: nil})
		return resultByte
	}
	err = c.Find(m).Skip((page - 1) * size).Limit(size).Sort(sort).All(data)
	if err != nil {
		resultByte, _ := json.Marshal(types.Page{Count: count, Data: nil})
		return resultByte
	} else {
		resultByte, _ := json.Marshal(types.Page{Count: count, Data: data})
		return resultByte
	}
}
