package model

import "github.com/irisnet/explorer/backend/orm/document"

type BlockRsp struct {
	CandidateMap map[string]string
	Counter      Counter
	document.Block
}

type Counter struct {
	TxCnt    int
	PropoCnt int
}
