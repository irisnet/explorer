package document

import (
	"fmt"

	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/irishub-sync/logger"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func desc(field string) string {
	return fmt.Sprintf("-%s", field)
}

func queryAll(collation string, selector, condition bson.M, sort string, size int, result interface{}) error {
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(collation).
		SetCondition(condition).
		SetSelector(selector).
		SetSort(sort).
		SetSize(size).
		SetResult(result)

	err := query.Exec()
	if err != nil {
		logger.Error("queryAll error", logger.Any("query", condition), logger.String("err", err.Error()))
	}
	return err
}

func queryOne(collation string, selector, condition bson.M, result interface{}) error {
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(collation).
		SetCondition(condition).
		SetSelector(selector).
		SetResult(result)

	err := query.Exec()
	if err != nil {
		logger.Error("queryOne", logger.Any("query", condition), logger.String("err", err.Error()))
	}
	return err
}

func pageQuery(collation string, selector, condition bson.M, sort string, page, size int, result interface{}) (int, error) {
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(collation).
		SetCondition(condition).
		SetSelector(selector).
		SetSort(sort).
		SetPage(page).
		SetSize(size).
		SetResult(result)

	cnt, err := query.ExecPage()
	if err != nil {
		logger.Error("pageQuery", logger.Any("query", condition), logger.String("err", err.Error()))
	}

	return cnt, err
}

func getDb() *mgo.Database {
	return orm.GetDatabase()
}

func asc(field string) string {
	return fmt.Sprintf("%s", field)
}
