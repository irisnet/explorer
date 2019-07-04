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

func TestGenHexAddrFromPubKey(t *testing.T) {

	vaList := []string{"fva1x292qss22x4rls6ygr7hhnp0et94vwwrdxhezx", "fva1d7rxfhyhqzgudv5nmne548t3xrk4tvazpg94cu",
		"fva1xtstdchjyzkddaptgyug62g23cta7eyzq49svq", "faa1x292qss22x4rls6ygr7hhnp0et94vwwrchaklp",
		"fcp1zcjduepq3dgenw6hha8kh0r4z5emqqv7u2k2w6qevytxwgdqkyugrsqmyykqq0zycl",
		"fcp1zcjduepqujegylnmyuymekm7sk70g5up0xq6p5r8zaj5rpjey7tqxxp0z3zqjns5cy",
		"fcp1zcjduepqcjgmderfahnlyrse563r2hcc3d4vjpafw03axzn3e87kfuqznjcsur8xrq",
		"fcp1zcjduepq3hy0q3ltgktya5gzvjfmkrhkeqzqau9t927ne8wdts36en490q7s8lxjq2",
		"fcp1zcjduepqgs9ffj72djsf7vytvm8vq3d2gzavytw0pkl5lee0dp4maej97x9swzlt60",
		"fcp1zcjduepqevwqk73gun8pp59wz6raddnsg2fczvs237cefl9ve7f94feh6lzsdr4qrf"}

	for k, v := range vaList {

		t.Logf("k: %v  v: %v hashAddr: %v \n", k, v, GenHexAddrFromPubKey(v))
	}

}
