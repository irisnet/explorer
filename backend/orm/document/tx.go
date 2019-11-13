package document

import (
	"fmt"
	"time"

	"github.com/irisnet/explorer/backend/orm"
	"github.com/irisnet/explorer/backend/types"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmCommonTx = "tx_common"
	TxStatusSuccess      = "success"
	TxStatusFail         = "fail"

	Tx_Field_Time                 = "time"
	Tx_Field_Height               = "height"
	Tx_Field_Hash                 = "tx_hash"
	Tx_Field_From                 = "from"
	Tx_Field_To                   = "to"
	Tx_Field_Signers              = "signers"
	Tx_Field_Amount               = "amount"
	Tx_Field_Type                 = "type"
	Tx_Field_Fee                  = "fee"
	Tx_Field_Memo                 = "memo"
	Tx_Field_Status               = "status"
	Tx_Field_Code                 = "code"
	Tx_Field_Log                  = "log"
	Tx_Field_GasUsed              = "gas_used"
	Tx_Field_GasPrice             = "gas_price"
	Tx_Field_ActualFee            = "actual_fee"
	Tx_Field_ProposalId           = "proposal_id"
	Tx_Field_Tags                 = "tags"
	Tx_Field_Msgs                 = "msgs"
	Tx_Field_StakeCreateValidator = "stake_create_validator"
	Tx_Field_StakeEditValidator   = "stake_edit_validator"

	Tx_Field_Msgs_UdInfo         = "msgs.msg.ud_info.source"
	Tx_Field_Msgs_Moniker        = "msgs.msg.moniker"
	Tx_Field_Msgs_UdInfo_Symbol  = "msgs.msg.ud_info.symbol"
	Tx_Field_Msgs_UdInfo_Gateway = "msgs.msg.ud_info.gateway"
	Tx_Field_Msgs_Hashcode       = "msgs.msg.hash_lock"
	Tx_AssetType_Native          = "native"
	Tx_AssetType_Gateway         = "gateway"

	Tx_Asset_TxType_Issue                = "IssueToken"
	Tx_Asset_TxType_Edit                 = "EditToken"
	Tx_Asset_TxType_Mint                 = "MintToken"
	Tx_Asset_TxType_TransferOwner        = "TransferTokenOwner"
	Tx_Asset_TxType_TransferGatewayOwner = "TransferGatewayOwner"
)

type Signer struct {
	AddrHex    string `bson:"addr_hex"`
	AddrBech32 string `bson:"addr_bech32"`
}

type Coin struct {
	Denom  string  `bson:"denom"`
	Amount float64 `bson:"amount"`
}

func (c Coin) Add(a Coin) Coin {
	if c.Denom == a.Denom {
		return Coin{
			Denom:  c.Denom,
			Amount: c.Amount + a.Amount,
		}
	}
	return c
}

type Coins []Coin

type Fee struct {
	Amount Coins `bson:"amount"`
	Gas    int64 `bson:"gas"`
}

type ActualFee struct {
	Denom  string  `bson:"denom"`
	Amount float64 `bson:"amount"`
}

type CommonTx struct {
	Time                 time.Time            `bson:"time"`
	Height               int64                `bson:"height"`
	TxHash               string               `bson:"tx_hash"`
	From                 string               `bson:"from"`
	To                   string               `bson:"to"`
	Amount               Coins                `bson:"amount"`
	Type                 string               `bson:"type"`
	Fee                  Fee                  `bson:"fee"`
	Memo                 string               `bson:"memo"`
	Status               string               `bson:"status"`
	Code                 uint32               `bson:"code"`
	Log                  string               `bson:"log"`
	GasUsed              int64                `bson:"gas_used"`
	GasWanted            int64                `bson:"gas_wanted"`
	GasPrice             float64              `bson:"gas_price"`
	ActualFee            ActualFee            `bson:"actual_fee"`
	ProposalId           uint64               `bson:"proposal_id"`
	Tags                 map[string]string    `bson:"tags"`
	Msgs                 []MsgItem            `bson:"msgs"`
	Signers              []Signer             `bson:"signers"`
}

func (tx CommonTx) String() string {
	return ""

}

type (
	Msg interface {
		Type() string
		String() string
	}
	MsgItem struct {
		Type    string      `bson:"type"`
		MsgData interface{} `bson:"msg"`
	}
	StakeCreateValidator struct {
		PubKey      string         `bson:"pub_key"`
		Description ValDescription `bson:"description"`
		Commission  CommissionMsg  `bson:"commission"`
	}
	CommissionMsg struct {
		Rate          string `bson:"rate"`            // the commission rate charged to delegators
		MaxRate       string `bson:"max_rate"`        // maximum commission rate which validator can ever charge
		MaxChangeRate string `bson:"max_change_rate"` // maximum daily increase of the validator commission
	}
	StakeEditValidator struct {
		CommissionRate string         `bson:"commission_rate"`
		Description    ValDescription `bson:"description"`
	}
)

// Description
type ValDescription struct {
	Moniker  string `bson:"moniker"`
	Identity string `bson:"identity"`
	Website  string `bson:"website"`
	Details  string `bson:"details"`
}

func (_ CommonTx) QueryByAddr(addr string, pageNum, pageSize int, istotal bool) (int, []CommonTx, error) {
	var data []CommonTx
	query := bson.M{}
	query["$or"] = []bson.M{{"from": addr}, {"to": addr}, {"signers": bson.M{"$elemMatch": bson.M{"addr_bech32": addr}}}}
	var typeArr []string
	typeArr = append(typeArr, types.BankList...)
	typeArr = append(typeArr, types.DeclarationList...)
	typeArr = append(typeArr, types.StakeList...)
	typeArr = append(typeArr, types.GovernanceList...)
	query[Tx_Field_Type] = bson.M{
		"$in": typeArr,
	}
	query = FilterUnknownTxs(query)

	total, err := pageQuery(CollectionNmCommonTx, nil, query, desc(Tx_Field_Time), pageNum, pageSize, istotal, &data)

	return total, data, err
}

func (_ CommonTx) QueryByPage(query bson.M, pageNum, pageSize int, istotal bool) (int, []CommonTx, error) {
	var data []CommonTx

	total, err := pageQuery(CollectionNmCommonTx, nil, query, desc(Tx_Field_Time), pageNum, pageSize, istotal, &data)

	return total, data, err
}

func (_ CommonTx) QueryHashActualFeeType() ([]CommonTx, error) {

	var selector = bson.M{"time": 1, "tx_hash": 1, "actual_fee": 1, "type": 1}
	var txs []CommonTx

	err := queryAll(CollectionNmCommonTx, selector, nil, desc(Tx_Field_Time), 10, &txs)
	return txs, err
}

func (_ CommonTx) QueryTxByHash(hash string) (CommonTx, error) {
	dbm := getDb()
	defer dbm.Session.Close()

	var result CommonTx
	query := bson.M{}
	query[Tx_Field_Hash] = hash
	err := dbm.C(CollectionNmCommonTx).Find(query).Sort(desc(Tx_Field_Time)).One(&result)

	return result, err
}

func (_ CommonTx) QueryHtlcTx(query bson.M) (CommonTx, error) {
	dbm := getDb()
	defer dbm.Session.Close()

	var result CommonTx
	err := dbm.C(CollectionNmCommonTx).Find(query).Sort(desc(Tx_Field_Time)).One(&result)

	return result, err
}

type Counter []struct {
	Type  string `bson:"_id,omitempty"`
	Count int
}

func (cArr Counter) String() string {
	res := ""
	for k, v := range cArr {
		res += fmt.Sprintf("idx: %v Type  :%v  \t	Count :%v \n", k, v.Type, v.Count)
	}
	return res
}

func (_ CommonTx) CountByType(query bson.M) (Counter, error) {

	var typeArr []string
	typeArr = append(typeArr, types.BankList...)
	typeArr = append(typeArr, types.DeclarationList...)
	typeArr = append(typeArr, types.StakeList...)
	typeArr = append(typeArr, types.GovernanceList...)
	query[Tx_Field_Type] = bson.M{
		"$in": typeArr,
	}
	query = FilterUnknownTxs(query)

	counter := Counter{}

	c := getDb().C(CollectionNmCommonTx)
	defer c.Database.Session.Close()

	pipe := c.Pipe(
		[]bson.M{
			{"$match": query},
			{"$group": bson.M{
				"_id":   "$type",
				"count": bson.M{"$sum": 1},
			}},
		},
	)

	err := pipe.All(&counter)

	return counter, err
}

func (_ CommonTx) GetTxlistByDuration(startTime, endTime string) ([]TxNumStat, error) {

	query := bson.M{}
	query["date"] = bson.M{"$gte": startTime, "$lt": endTime}

	var selector = bson.M{"date": 1, "num": 1}
	var txNumStatList []TxNumStat

	q := orm.NewQuery()
	q.SetCollection(CollectionTxNumStat)
	q.SetCondition(query)
	q.SetSelector(selector).SetSort("date")
	q.SetResult(&txNumStatList)

	defer q.Release()

	err := q.Exec()
	return txNumStatList, err
}

func (_ CommonTx) GetTxCountByDuration(startTime, endTime time.Time) (int, error) {

	db := orm.GetDatabase()
	defer db.Session.Close()

	txStore := db.C(CollectionNmCommonTx)

	query := bson.M{}
	query = FilterUnknownTxs(query)
	query["time"] = bson.M{"$gte": startTime, "$lt": endTime}

	return txStore.Find(query).Count()
}

func (_ CommonTx) QueryProposalTxFromById(idArr []uint64) (map[uint64]string, error) {

	selector := bson.M{Tx_Field_From: 1, Tx_Field_ProposalId: 1}
	condition := bson.M{Tx_Field_Type: "SubmitProposal", Tx_Field_Status: "success", Tx_Field_ProposalId: bson.M{"$in": idArr}}
	var txs []CommonTx
	condition = FilterUnknownTxs(condition)

	err := queryAll(CollectionNmCommonTx, selector, condition, desc(Tx_Field_Time), 0, &txs)

	proposerById := map[uint64]string{}

	for _, v := range txs {
		proposerById[v.ProposalId] = v.From
	}

	return proposerById, err
}

func (_ CommonTx) QueryProposalTxListById(idArr []uint64) ([]CommonTx, error) {

	selector := bson.M{Tx_Field_Amount: 1, Tx_Field_ProposalId: 1}
	condition := bson.M{Tx_Field_Type: "SubmitProposal", Tx_Field_Status: "success", Tx_Field_ProposalId: bson.M{"$in": idArr}}
	var txs []CommonTx
	condition = FilterUnknownTxs(condition)

	err := queryAll(CollectionNmCommonTx, selector, condition, desc(Tx_Field_Time), 0, &txs)

	return txs, err
}

func (_ CommonTx) QueryProposalInitAmountTxById(id int) (CommonTx, error) {
	selector := bson.M{Tx_Field_Amount: 1, Tx_Field_ProposalId: 1}
	condition := bson.M{Tx_Field_Type: "SubmitProposal", Tx_Field_Status: "success", Tx_Field_ProposalId: id}
	condition = FilterUnknownTxs(condition)
	var txs CommonTx

	err := queryOne(CollectionNmCommonTx, selector, condition, &txs)
	return txs, err
}

func (_ CommonTx) QueryProposalTxById(proposalId int64, page, size int, total bool, iaaAddrs []string) (int, []CommonTx, error) {

	txs := []CommonTx{}

	selector := bson.M{
		Tx_Field_Height: 1,
		Tx_Field_Time:   1,
		Tx_Field_Hash:   1,
		Tx_Field_From:   1,
	}
	condition := bson.M{
		Tx_Field_Status:     types.TxTypeStatus,
		Tx_Field_ProposalId: proposalId,
		Tx_Field_Type:       types.TxTypeVote,
		Tx_Field_From:       bson.M{"$in": iaaAddrs},
	}
	condition = FilterUnknownTxs(condition)
	sort := fmt.Sprintf("-%v", Tx_Field_Height)

	num, err := pageQuery(CollectionNmCommonTx, selector, condition, sort, page, size, total, &txs)

	return num, txs, err
}

func (_ CommonTx) QueryDepositedProposalTxByValidatorWithSubmitOrDepositType(validatorAddrAcc string, page, size int, total bool) (int, []CommonTx, error) {

	txs := []CommonTx{}
	selector := bson.M{
		Tx_Field_Hash:       1,
		Tx_Field_From:       1,
		Tx_Field_Amount:     1,
		Tx_Field_Type:       1,
		Tx_Field_ProposalId: 1,
	}
	condition := bson.M{
		Tx_Field_Status: types.TxTypeStatus,
		Tx_Field_From:   validatorAddrAcc,
		Tx_Field_Type: bson.M{
			"$in": []string{types.TxTypeSubmitProposal, types.TxTypeDeposit},
		},
	}
	condition = FilterUnknownTxs(condition)
	sort := fmt.Sprintf("-%v", Tx_Field_Height)
	num, err := pageQuery(CollectionNmCommonTx, selector, condition, sort, page, size, total, &txs)

	return num, txs, err
}

func (_ CommonTx) QueryProposalTxByIdWithSubmitOrDepositType(proposalId int64, page, size int, total bool) (int, []CommonTx, error) {

	txs := []CommonTx{}
	selector := bson.M{
		Tx_Field_Hash:   1,
		Tx_Field_From:   1,
		Tx_Field_Amount: 1,
		Tx_Field_Type:   1,
		Tx_Field_Time:   1,
	}
	condition := bson.M{
		Tx_Field_Status:     types.TxTypeStatus,
		Tx_Field_ProposalId: proposalId,
		Tx_Field_Type: bson.M{
			"$in": []string{types.TxTypeSubmitProposal, types.TxTypeDeposit},
		},
	}
	condition = FilterUnknownTxs(condition)
	sort := fmt.Sprintf("-%v", Tx_Field_Height)
	num, err := pageQuery(CollectionNmCommonTx, selector, condition, sort, page, size, total, &txs)

	return num, txs, err
}

func (_ CommonTx) QueryTxAsset(assetType, tokenType, symbol, gateway string, page, size int, total bool) (int, []CommonTx, error) {
	txs := []CommonTx{}
	selector := bson.M{
		Tx_Field_Hash:      1,
		Tx_Field_Height:    1,
		Tx_Field_From:      1,
		Tx_Field_To:        1,
		Tx_Field_Amount:    1,
		Tx_Field_Type:      1,
		Tx_Field_Status:    1,
		Tx_Field_ActualFee: 1,
		Tx_Field_Tags:      1,
		Tx_Field_Msgs:      1,
		Tx_Field_Time:      1,
	}
	condition := bson.M{
		Tx_Field_Msgs_UdInfo: assetType,
	}
	condition = FilterUnknownTxs(condition)
	if tokenType != "" {
		condition[Tx_Field_Type] = tokenType
	} else {
		condition[Tx_Field_Type] = bson.M{
			"$in": []string{types.TxTypeIssueToken, types.TxTypeEditToken, types.TxTypeMintToken, types.TxTypeTransferTokenOwner},
		}
	}
	if symbol != "" {
		condition[Tx_Field_Msgs_UdInfo_Symbol] = bson.M{
			"$regex":   symbol,
			"$options": "$i",
		}
	}
	if gateway != "" {
		condition[Tx_Field_Msgs_UdInfo_Gateway] = bson.M{
			"$regex":   gateway,
			"$options": "$i",
		}
	}
	sort := fmt.Sprintf("-%v", Tx_Field_Height)
	num, err := pageQuery(CollectionNmCommonTx, selector, condition, sort, page, size, total, &txs)
	return num, txs, err
}

func (_ CommonTx) QueryTxTransferGatewayOwner(moniker string, page, size int, total bool) (int, []CommonTx, error) {
	txs := []CommonTx{}
	selector := bson.M{
		Tx_Field_Hash:      1,
		Tx_Field_Height:    1,
		Tx_Field_From:      1,
		Tx_Field_To:        1,
		Tx_Field_Amount:    1,
		Tx_Field_Type:      1,
		Tx_Field_Status:    1,
		Tx_Field_ActualFee: 1,
		Tx_Field_Tags:      1,
		Tx_Field_Msgs:      1,
		Tx_Field_Time:      1,
	}

	condition := bson.M{
		Tx_Field_Type: Tx_Asset_TxType_TransferGatewayOwner,
	}
	condition = FilterUnknownTxs(condition)
	if moniker != "" {
		condition[Tx_Field_Msgs_Moniker] = moniker
	}

	sort := fmt.Sprintf("-%v", Tx_Field_Height)
	num, err := pageQuery(CollectionNmCommonTx, selector, condition, sort, page, size, total, &txs)
	return num, txs, err
}

func FilterUnknownTxs(query bson.M) bson.M {

	if status, ok := query["status"]; ok {
		query["status"] = status.(string)
	} else {
		query["status"] = bson.M{
			"$in": []string{TxStatusSuccess, TxStatusFail},
		}
	}
	return query
}
