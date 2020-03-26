package document

import (
	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNameTaskControl = "ex_task_control"
	TCFieldTaskName           = "task_name"
	TCFieldTimeInterval       = "time_interval"
	TCFieldIsInProcess        = "is_in_process"
	TCFieldLatestExecTime     = "latest_exec_time"
	TCFieldCreateAt           = "create_at"
	TCFieldUpdateAt           = "update_at"
)

type (
	TaskControl struct {
		TaskName       string `bson:"task_name"`
		TimeInterval   int64  `bson:"time_interval"`
		IsInProcess    bool   `bson:"is_in_process"`
		LatestExecTime int64  `bson:"latest_exec_time"`
		CreateAt       int64  `bson:"create_at"`
	}
)

func (d TaskControl) Name() string {
	return CollectionNameTaskControl
}

func (d TaskControl) PkKvPair() map[string]interface{} {
	return bson.M{
		TCFieldTaskName: d.TaskName,
	}
}

func (d TaskControl) EnsureIndexes() []mgo.Index {
	indexes := []mgo.Index{
		{
			Key:        []string{"-task_name"},
			Unique:     true,
			Background: true,
		},
	}

	return indexes
}

func (d TaskControl) findOne(selector, condition bson.M) (TaskControl, error) {
	var res TaskControl
	q := orm.NewQuery()
	defer q.Release()

	if err := queryOne(d.Name(), selector, condition, &res); err != nil {
		return res, err
	} else {
		return res, nil
	}
}

func (d TaskControl) QueryOneByTaskName(taskName string) (TaskControl, error) {
	selector := bson.M{}
	condition := bson.M{
		TCFieldTaskName: taskName,
	}
	return d.findOne(selector, condition)
}

func (d TaskControl) Save(record TaskControl) error {
	q := orm.NewQuery()
	defer q.Release()

	c := q.GetDb().C(d.Name())

	now := time.Now().Unix()
	record.CreateAt = now

	if err := c.Insert(record); err != nil {
		return err
	} else {
		return nil
	}
}

func (d TaskControl) UpdateByPK(update TaskControl) error {
	q := orm.NewQuery()
	defer q.Release()

	c := q.GetDb().C(d.Name())

	if err := c.Update(update.PkKvPair(), update); err != nil {
		return err
	}

	return nil
}

func (d TaskControl) LockTaskControl(taskName string) error {
	q := orm.NewQuery()
	defer q.Release()

	c := q.GetDb().C(d.Name())

	selector := bson.M{
		TCFieldTaskName:    taskName,
		TCFieldIsInProcess: false,
	}
	update := bson.M{
		"$set": bson.M{
			TCFieldIsInProcess: true,
		},
	}
	if err := c.Update(selector, update); err != nil {
		return err
	}

	return nil
}

func (d TaskControl) UnlockTaskControl(taskName string) error {
	q := orm.NewQuery()
	defer q.Release()

	c := q.GetDb().C(d.Name())

	selector := bson.M{
		TCFieldTaskName:    taskName,
		TCFieldIsInProcess: true,
	}
	update := bson.M{
		"$set": bson.M{
			TCFieldIsInProcess:    false,
			TCFieldLatestExecTime: time.Now().Unix(),
		},
	}

	if err := c.Update(selector, update); err != nil {
		return err
	}

	return nil
}
