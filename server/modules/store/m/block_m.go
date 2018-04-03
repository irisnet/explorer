package m

import (
	"time"
	"gopkg.in/mgo.v2"
	"log"
	"github.com/irisnet/irisplorer.io/server/modules/store"
)

const (
	DocsNmBlock = "block"
)

type Block struct {
	Height int64     `json:"height" bson:"height"`
	Time   time.Time `json:"time" bson:"time"`
	TxNum  int64     `json:"tx_num" bson:"tx_num"`
}

func (d Block) Name() string {
	return DocsNmBlock
}

func (d Block) PkKvPair() (string, interface{}) {
	return "height", d.Height
}

func (d Block) Index() mgo.Index {
	return mgo.Index{
		Key:        []string{"height"}, // 索引字段， 默认升序,若需降序在字段前加-
		Unique:     true,              // 唯一索引 同mysql唯一索引
		DropDups:   false,              // 索引重复替换旧文档,Unique为true时失效
		Background: true,               // 后台创建索引
	}
}

func QueryBlockByPage(page int) ([]Block) {
	result := []Block{}
	query := func(c *mgo.Collection) error {
		skip := (page - 1) * store.PageSize
		err := c.Find(nil).Sort("-height").Skip(skip).Limit(store.PageSize).All(&result)
		return err
	}

	if store.ExecCollection(DocsNmBlock, query) != nil {
		log.Printf("Block is Empry")
	}

	return result
}
