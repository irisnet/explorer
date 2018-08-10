package utils

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"encoding/json"
	"github.com/irisnet/explorer/server/types"
	"time"
)

var database string

func init() {
	mongoUrl := GetEnv("DB_URL", "192.168.150.7:30000")
	database := GetEnv("DB_DATABASE", "sync-iris")
	user := GetEnv("DB_USER", "iris")
	passwd := GetEnv("DB_PASSWORD", "irispassword")
	log.Println("Connecting " + mongoUrl + "/" + database)
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{mongoUrl},
		Database:  database,
		Username:  user,
		Password:  passwd,
		Direct:    false,
		Timeout:   time.Second * 10,
		PoolLimit: 100, // Session.SetPoolLimit
	}

	var err error
	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln(err)
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(300)
}

var session *mgo.Session

func GetDatabase() *mgo.Database {
	// max session num is 4096
	return session.Copy().DB(database)
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
