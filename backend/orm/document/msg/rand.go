package msg

type TxMsgRequestRand struct {
	Consumer      string `bson:"consumer"`       // request address
	BlockInterval int64  `bson:"block-interval"` // block interval after which the requested random number will be generated
}

func (msg TxMsgRequestRand) Nil() {
}
