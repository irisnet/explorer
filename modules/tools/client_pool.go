package tools

import (
	"net/http"
	"fmt"
	"github.com/tendermint/go-wire/data"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/cosmos/cosmos-sdk/client/commands"
	"errors"
	"github.com/spf13/viper"
)

const (
	ClientMaxCon = "client-max-conn"
)

func FmtOutPutResult(w http.ResponseWriter, res interface{}) error {
	json, err := data.ToJSON(res)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "%s\n", json)
	return err
}

var Pool = NodePool{}

type NodePool struct {
	nodes []Node
}

func (pool NodePool) Init() {
	len := viper.GetInt(ClientMaxCon)
	pool.nodes = make([]Node, len)
	for i := 0; i < len; i++ {
		node := Node{
			client: commands.GetNode(),
			used:   false,
			id:     i,
		}
		pool.nodes[i] = node
	}
}

func (pool NodePool) GetNode() (client.Client, error) {
	for _, node := range pool.nodes {
		if !node.used {
			node.used = true
			pool.nodes[node.id] = node
			return node.client, nil
		}
	}
	return nil, errors.New("pool is empty")
}

type Node struct {
	client client.Client
	used   bool
	id     int
}

func (n Node) Release() {
	n.used = false
	Pool.nodes[n.id] = n
}



