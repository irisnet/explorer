package service

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBlockService_Query(t *testing.T) {
	result := blockService.Query(1000)
	bz, _ := json.Marshal(result)
	fmt.Println(string(bz))
}
