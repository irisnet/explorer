package utils

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"log"
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
	return session.Copy().DB("sync_iris")
}
