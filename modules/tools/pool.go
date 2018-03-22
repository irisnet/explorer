package tools

import (
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"errors"
	"github.com/spf13/viper"
	"log"
)

const (
	ClientMaxCon = "client-max-conn"
)


type NodePool struct {
	nodes []Node
	available int8
	maxConnection int8
}

type Node struct {
	Client client.Client
	used   bool
	id     int
}

var pool = NodePool{}

func GetNode() Node{
	c,err := getNode()
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func Init() {
	len := viper.GetInt(ClientMaxCon)
	pool.nodes = make([]Node, len)
	for i := 0; i < len; i++ {
		node := Node{
			Client: commands.GetNode(),
			used:   false,
			id:     i,
		}
		pool.nodes[i] = node
		pool.available = int8(len)
		pool.maxConnection = int8(len)
	}
}

func getNode() (Node, error) {


	if pool.available == 0 {
		log.Fatal("client pool has no available connection")
	}

	for _, node := range pool.nodes {
		if !node.used {
			node.used = true
			pool.nodes[node.id] = node
			pool.available--

			log.Printf("current available coonection ：%d",pool.available)
			return node, nil
		}
	}
	return Node{}, errors.New("pool is empty")
}

func (n Node) Release() {
	n.used = false
	pool.nodes[n.id] = n
	pool.available++
}






