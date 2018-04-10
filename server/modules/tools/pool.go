package tools

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/rpc/client"
	"log"
)

type NodePool struct {
	nodes          []Node
	available      int64
	used           int64
	maxConnection  int64
	initConnection int64
}

type Node struct {
	Client client.Client
	used   bool
	id     int64
}

var pool = NodePool{}

func GetNode() Node {
	c, err := getNode()
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func Init() {
	len := viper.GetInt64(InitConnectionNum)
	pool.initConnection = len
	pool.nodes = make([]Node, len)
	for i := int64(0); i < len; i++ {
		createConnection(i)
	}
}

func createConnection(id int64) Node {
	node := Node{
		Client: commands.GetNode(),
		used:   false,
		id:     id,
	}
	pool.nodes[id] = node
	pool.available++
	pool.maxConnection = viper.GetInt64(MaxConnectionNum)
	return node
}

func getNode() (Node, error) {
	if pool.available == 0 {
		maxConnNum := viper.GetInt64(MaxConnectionNum)
		if pool.used < maxConnNum {
			var node Node
			for i := int64(len(pool.nodes)); i < maxConnNum; i++ {
				node = createConnection(i)
			}
			return node, nil
		} else {
			log.Fatal("client pool has no available connection")
		}
	}

	for _, node := range pool.nodes {
		if !node.used {
			node.used = true
			pool.nodes[node.id] = node
			pool.available--
			pool.used++
			log.Printf("current available coonection ：%d", pool.available)
			return node, nil
		}
	}
	return Node{}, errors.New("pool is empty")
}

func (n Node) Release() {
	n.used = false
	pool.nodes[n.id] = n
	pool.available++
	pool.used--
}
