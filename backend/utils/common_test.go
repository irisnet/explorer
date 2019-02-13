package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestRunTimer(t *testing.T) {
	RunTimer(1*time.Second, Sec, func() {
		fmt.Printf("test RunTimer:%s\n", time.Now().String())
	})
	time.Sleep(1 * time.Hour)
}
