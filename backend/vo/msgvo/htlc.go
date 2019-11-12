package msgvo

import (
	"github.com/irisnet/explorer/backend/utils"
	"encoding/json"
)

type TxMsgCreateHTLC struct {
	Sender               string      `json:"sender"`                  // the initiator address
	To                   string      `json:"to"`                      // the destination address
	ReceiverOnOtherChain string      `json:"receiver_on_other_chain"` // the claim receiving address on the other chain
	Amount               utils.Coins `json:"amount"`                  // the amount to be transferred
	HashLock             string      `json:"hash_lock"`               // the hash lock generated from secret (and timestamp if provided)
	Timestamp            uint64      `json:"timestamp"`               // if provided, used to generate the hash lock together with secret
	TimeLock             uint64      `json:"time_lock"`               // the time span after which the HTLC will expire
}

func (vo *TxMsgCreateHTLC) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgClaimHTLC struct {
	Sender   string `json:"sender"`    // the initiator address
	HashLock string `json:"hash_lock"` // the hash lock identifying the HTLC to be claimed
	Secret   string `json:"secret"`    // the secret with which to claim
}

func (vo *TxMsgClaimHTLC) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}

type TxMsgRefundHTLC struct {
	Sender   string `json:"sender"`    // the initiator address
	HashLock string `json:"hash_lock"` // the hash lock identifying the HTLC to be refunded
}

func (vo *TxMsgRefundHTLC) BuildMsgByUnmarshalJson(data []byte) error {
	return json.Unmarshal(data, vo)
}
