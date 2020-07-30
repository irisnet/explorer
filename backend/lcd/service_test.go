package lcd

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestGetServiceBindings(t *testing.T) {
	res := GetServiceBindings("assettransfer")
	data,_ := json.Marshal(res)
	fmt.Println(string(data))
}

func TestGetServiceContexts(t *testing.T) {
	res := GetServiceContexts("48F166919DB0394BA1553C013C2CF26EFC2ED870C27017C016C651BCDF5F97B60000000000000000")
	data,_ := json.Marshal(res)
	fmt.Println(string(data))
}
