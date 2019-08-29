package document

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/orm"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmBlock = "block"

	Block_Field_Height          = "height"
	Block_Field_Hash            = "hash"
	Block_Field_Time            = "time"
	Block_Field_NumTxs          = "num_txs"
	Block_Field_Meta            = "meta"
	Block_Field_Block           = "block"
	Block_Field_Validators      = "validators"
	Block_Field_ProposalAddress = "proposal_address"
)

type Block struct {
	Height       int64         `bson:"height"`
	Hash         string        `bson:"hash"`
	Time         time.Time     `bson:"time"`
	NumTxs       int64         `bson:"num_txs"`
	Meta         BlockMeta     `bson:"meta"`
	Block        BlockContent  `bson:"block"`
	Validators   []TmValidator `bson:"validators"`
	Result       BlockResults  `bson:"results"`
	ProposalAddr string        `bson:"proposal_address"`
}

func (b Block) String() string {
	return fmt.Sprintf(`
		Height      :%v
		Hash        :%v
		Time        :%v
		NumTxs      :%v
		Meta        :%v
		Block       :%v
		Validators  :%v
		Result      :%v
		ProposalAddr:%v
		`, b.Height, b.Hash, b.Time, b.NumTxs, b.Meta, b.Block, b.Validators, b.Result, b.ProposalAddr)
}

func (_ Block) QueryBlockHeightTimeHashByHeight(height int64) (Block, error) {

	var block Block
	var selector = bson.M{"height": 1, "time": 1, "hash": 1}
	var query = orm.NewQuery().
		SetCollection(CollectionNmBlock).
		SetSelector(selector).
		SetCondition(bson.M{Block_Field_Height: height}).
		SetResult(&block)

	defer query.Release()

	err := query.Exec()
	return block, err
}

func (_ Block) GetBlockListByOffsetAndSize(offset, size int) ([]Block, error) {

	var selector = bson.M{"height": 1, "time": 1, "num_txs": 1, "hash": 1, "validators.pub_key": 1, "validators.address": 1,
		"validators.voting_power": 1, "block.last_commit.precommits.validator_address": 1, "meta.header.total_txs": 1, "proposal_address": 1}
	var blocks []Block

	sort := desc(Block_Field_Height)

	err := querylistByOffsetAndSize(CollectionNmBlock, selector, nil, sort, offset, size, &blocks)

	return blocks, err
}

func (_ Block) GetBlockListByPage(offset, size int, total bool) (int, []Block, error) {

	var selector = bson.M{"height": 1, "time": 1, "num_txs": 1, "hash": 1, "validators.address": 1, "validators.voting_power": 1, "block.last_commit.precommits.validator_address": 1, "meta.header.total_txs": 1}

	var blocks []Block

	sort := desc(Block_Field_Height)
	var cnt, err = pageQuery(CollectionNmBlock, selector, bson.M{"height": bson.M{"$gt": 0}}, sort, offset, size, total, &blocks)

	return cnt, blocks, err
}

func (_ Block) GetRecentBlockList() ([]Block, error) {
	var blocks []Block
	var selector = bson.M{"height": 1, "time": 1, "num_txs": 1}

	sort := desc(Block_Field_Height)
	err := queryAll(CollectionNmBlock, selector, nil, sort, 10, &blocks)
	return blocks, err
}

func (_ Block) QueryOneBlockOrderByHeightAsc() (Block, error) {

	var blocks []Block

	err := queryAll(CollectionNmBlock, nil, nil, asc(Block_Field_Height), 1, &blocks)

	if len(blocks) == 1 {
		return blocks[0], err
	}

	return Block{}, err
}

func (_ Block) QueryOneBlockOrderByHeightDesc() (Block, error) {

	db := orm.GetDatabase()
	defer db.Session.Close()

	var firstBlock Block

	err := db.C(CollectionNmBlock).Find(nil).Sort("-height").One(&firstBlock)
	return firstBlock, err
}

func (_ Block) QueryBlocksByDurationWithHeightAsc(startTime, endTime time.Time) ([]Block, error) {
	db := orm.GetDatabase()
	defer db.Session.Close()

	blocks := []Block{}
	err := db.C(CollectionNmBlock).Find(bson.M{"time": bson.M{"$gte": startTime, "$lt": endTime}}).Sort("height").All(&blocks)
	return blocks, err
}

func (_ Block) QueryValidatorsByHeightList(hArr []int64) ([]Block, error) {

	var selector = bson.M{Block_Field_Height: 1, Block_Field_Validators: 1}

	sort := desc(Block_Field_Height)
	var blocks []Block
	err := queryAll(CollectionNmBlock, selector, bson.M{"height": bson.M{"$in": hArr}}, sort, 0, &blocks)
	return blocks, err
}

type BlockMeta struct {
	BlockID BlockID `bson:"block_id"`
	Header  Header  `bson:"header"`
}

type BlockID struct {
	Hash        string        `bson:"hash"`
	PartsHeader PartSetHeader `bson:"parts"`
}

type PartSetHeader struct {
	Total int    `bson:"total"`
	Hash  string `bson:"hash"`
}

type Header struct {
	// basic block info
	ChainID string    `bson:"chain_id"`
	Height  int64     `bson:"height"`
	Time    time.Time `bson:"time"`
	NumTxs  int64     `bson:"num_txs"`

	// prev block info
	LastBlockID BlockID `bson:"last_block_id"`
	TotalTxs    int64   `bson:"total_txs"`

	// hashes of block data
	LastCommitHash string `bson:"last_commit_hash"` // commit from validators from the last block
	DataHash       string `bson:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash  string `bson:"validators_hash"`   // validators for the current block
	ConsensusHash   string `bson:"consensus_hash"`    // consensus params for current block
	AppHash         string `bson:"app_hash"`          // state after txs from the previous block
	LastResultsHash string `bson:"last_results_hash"` // root hash of all results from the txs from the previous block

	// consensus info
	EvidenceHash string `bson:"evidence_hash"` // evidence included in the block
}

type BlockContent struct {
	LastCommit Commit `bson:"last_commit"`
}

type Commit struct {
	// NOTE: The Precommits are in order of address to preserve the bonded ValidatorSet order.
	// Any peer with a block can gossip precommits by index with a peer without recalculating the
	// active ValidatorSet.
	BlockID    BlockID `bson:"block_id"`
	Precommits []Vote  `bson:"precommits"`
}

// Represents a prevote, precommit, or commit vote from validators for consensus.
type Vote struct {
	ValidatorAddress string    `bson:"validator_address"`
	ValidatorIndex   int       `bson:"validator_index"`
	Height           int64     `bson:"height"`
	Round            int       `bson:"round"`
	Timestamp        time.Time `bson:"timestamp"`
	Type             byte      `bson:"type"`
	BlockID          BlockID   `bson:"block_id"` // zero if vote is nil.
	Signature        Signature `bson:"signature"`
}

type Signature struct {
	Type  string `bson:"type"`
	Value string `bson:"value"`
}

type TmValidator struct {
	Address     string `bson:"address"`
	PubKey      string `bson:"pub_key"`
	VotingPower int64  `bson:"voting_power"`
	Accum       int64  `bson:"accum"`
}

type BlockResults struct {
	DeliverTx  []ResponseDeliverTx `bson:"deliver_tx"`
	EndBlock   ResponseEndBlock    `bson:""end_block""`
	BeginBlock ResponseBeginBlock  `bson:""begin_block""`
}

type ResponseDeliverTx struct {
	Code      uint32   `bson:"code"`
	Data      string   `bson:"data"`
	Log       string   `bson:"log"`
	Info      string   `bson:"info"`
	GasWanted int64    `bson:"gas_wanted"`
	GasUsed   int64    `bson:"gas_used"`
	Tags      []KvPair `bson:"tags"`
	Codespace string   `bson:"codespace"`
}

type ResponseEndBlock struct {
	ValidatorUpdates      []ValidatorUpdate `bson:"validator_updates"`
	ConsensusParamUpdates ConsensusParams   `bson:"consensus_param_updates"`
	Tags                  []KvPair          `bson:"tags"`
}

type ValidatorUpdate struct {
	PubKey string `bson:"pub_key"`
	Power  int64  `bson:"power"`
}

type ConsensusParams struct {
	BlockSize BlockSizeParams `bson:"block_size"`
	Evidence  EvidenceParams  `bson:"evidence"`
	Validator ValidatorParams `bson:"validator"`
}

type BlockSizeParams struct {
	MaxBytes int64 `bson:"max_bytes"`
	MaxGas   int64 `bson:"max_gas"`
}

type EvidenceParams struct {
	MaxAge int64 `bson:"max_age"`
}

type ValidatorParams struct {
	PubKeyTypes []string `bson:"pub_key_types`
}

type ResponseBeginBlock struct {
	Tags []KvPair `bson:"tags"`
}

type KvPair struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
}

func (d Block) Name() string {
	return CollectionNmBlock
}

func (d Block) PkKvPair() map[string]interface{} {
	return bson.M{Block_Field_Height: d.Height}
}

type ResValidatorPreCommits struct {
	Address       string `bson:"_id"`
	PreCommitsNum int64  `bson:"num"`
}
