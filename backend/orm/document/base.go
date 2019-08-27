package document

import (
	"fmt"

	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2"
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

func querylistByOffsetAndSize(collection string, selector, condition bson.M, sort string, offset, size int, result interface{}) error {

	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(collection).
		SetCondition(condition).
		SetSelector(selector).
		SetSort(sort).
		SetOffset(offset).
		SetSize(size).
		SetResult(result)

	err := query.Exec()
	if err != nil {
		logger.Error("QuerylistByOffsetAndSize error", logger.Any("sort", sort), logger.Any("offset", offset), logger.Any("size", size), logger.Any("query", condition), logger.String("err", err.Error()))
	}
	return err
}

func pageQuery(collation string, selector, condition bson.M, sort string, page, size int, total bool, result interface{}) (int, error) {
	var query = orm.NewQuery()
	defer query.Release()
	query.SetCollection(collation).
		SetCondition(condition).
		SetSelector(selector).
		SetSort(sort).
		SetPage(page).
		SetSize(size).
		SetResult(result)

	cnt, err := query.ExecPage(total)
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
