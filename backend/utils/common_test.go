package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/websocket"
	"testing"
	"time"
)

func TestParseInt(t *testing.T) {
	_, ok := ParseInt("1")
	assert.True(t, ok)

	_, ok = ParseUint("-1")
	assert.False(t, ok)

	_, ok = ParseInt("sd")
	assert.False(t, ok)
}

func TestParseStringToFloat(t *testing.T) {
	a, _ := ParseStringToFloat("10")
	b, _ := ParseStringToFloat("4")
	fmt.Println(a / b)
}

func TestRunTimer(t *testing.T) {
	RunTimer(1, Sec, func() {
		fmt.Printf("test RunTimer:%s\n", time.Now().String())
	})
	time.Sleep(1 * time.Hour)
}

func TestWebsocket(t *testing.T) {
	conn, err := websocket.Dial("ws://192.168.150.7:30657/websocket", "", "http://192.168.150.7:30657")
	fmt.Println(err)
	conn.Write([]byte(`{
                    "jsonrpc": "2.0",
                    "method": "subscribe",
                    "id": "0",
                    "params": {
                        "query": "tm.event='NewBlock'"
                    }}`))

	conn.Write([]byte(`{
                    "jsonrpc": "2.0",
                    "method": "subscribe",
                    "id": "1",
                    "params": {
                        "query": "tm.event='CompleteProposal'"
                    }}`))

	for {
		request := make([]byte, 128)
		readLen, _ := conn.Read(request)
		fmt.Println(string(request[0:readLen]))
	}

}

func TestMd5Encryption(t *testing.T) {
	data := []byte("nil")
	t.Log(Md5Encryption(data))
}

func TestFuncAddStr(t *testing.T) {
	data := FuncAddStr("11.1111111", "7.777777")
	t.Log(data.FloatString(18))
	data = FuncAddStr("11.1111111", "")
	if data == nil {
		t.Log("OK")
	}
}