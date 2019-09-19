package service

import (
	"github.com/irisnet/explorer/backend/utils"
	"testing"
)

func TestBaseService_QueryBlackList(t *testing.T) {
	service := BaseService{}

	res := service.QueryBlackList()
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}
