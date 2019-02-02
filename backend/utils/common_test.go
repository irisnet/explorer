package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInt(t *testing.T) {
	_, ok := ParseInt("1")
	assert.True(t, ok)

	_, ok = ParseUint("-1")
	assert.False(t, ok)

	_, ok = ParseInt("sd")
	assert.False(t, ok)
}

func TestGenesis(t *testing.T) {
	fmt.Printf("%v", Genesis().Result.Genesis.AppState.Guardian.Profilers)
}
