package types

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/irisnet/explorer/backend/utils"
	"strings"
)

func FromBech32ToAddr(pubKey string) string {
	_, bz, _ := utils.DecodeAndConvert(pubKey)
	addBz := sumTruncated(bz[5:])
	return strings.ToUpper(hex.EncodeToString(addBz))
}

func sumTruncated(bz []byte) []byte {
	hash := sha256.Sum256(bz)
	return hash[:20]
}
