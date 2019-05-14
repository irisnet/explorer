package model

import (
	"github.com/irisnet/explorer/backend/orm/document"
)

type TokenFlows struct {
	Total int                  `json:"total"`
	Items []document.TokenFlow `json:"items"`
}
