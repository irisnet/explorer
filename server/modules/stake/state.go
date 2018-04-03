package stake

import "github.com/tendermint/go-crypto"

var (
	// Keys for store prefixes
	CandidatesPubKeysKey = []byte{0x01} // key for all candidates' pubkeys
	ParamKey             = []byte{0x02} // key for global parameters relating to staking

	// Key prefixes
	CandidateKeyPrefix      = []byte{0x03} // prefix for each key to a candidate
	DelegatorBondKeyPrefix  = []byte{0x04} // prefix for each key to a delegator's bond
	DelegatorBondsKeyPrefix = []byte{0x05} // prefix for each key to a delegator's bond
)

// DelegatorBond represents the bond with tokens held by an account.  It is
// owned by one delegator, and is associated with the voting power of one
// pubKey.
type DelegatorBond struct {
	PubKey crypto.PubKey
	Shares uint64
}
