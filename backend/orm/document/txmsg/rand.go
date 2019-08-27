package txmsg

type TxMsgRequestRand struct {
	Consumer      string `bson:"consumer"`       // request address
	BlockInterval int64  `bson:"block_interval"` // block interval after which the requested random number will be generated
}
