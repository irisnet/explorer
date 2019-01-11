package model

import "github.com/irisnet/explorer/backend/orm/document"

type BlockVo struct {
	CandidateMap map[string]string
	Counter      CounterVo
	document.Block
}

type CounterVo struct {
	TxCnt    int
	PropoCnt int
}
