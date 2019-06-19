package utils

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeAndConvert(t *testing.T) {
	bech32Str := "fcp1zcjduepq5vpny3xgk0wkweza6vfgchmz5rylqv4f3p50huzwa5dne3prxl9s02mecs"
	wanted := "1624DE6420A3033244C8B3DD67645DD3128C5F62A0C9F032A98868FBF04EED1B3CC42337CB"
	if prefix, bytes, err := DecodeAndConvert(bech32Str); err != nil {
		t.Fatal(err)
	} else {
		hexStr := strings.ToUpper(hex.EncodeToString(bytes))
		t.Log(prefix)
		t.Log(hexStr)
		assert.Equal(t, wanted, hexStr)
	}
}

func TestConvert(t *testing.T) {
	dst := "fva"
	bech32Str := "faa17cjdg63thy2vfqvvgj5lfv5dp339t0lr99wc8p"
	wanted := "fva17cjdg63thy2vfqvvgj5lfv5dp339t0lrs5yh6x"
	res := Convert(dst, bech32Str)
	t.Log(res)
	assert.Equal(t, res, wanted)
}
