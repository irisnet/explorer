package lcd

type LcdHtlc struct {
	Type  string   `json:"type"`
	Value HtlcData `json:"value"`
}

type HtlcData struct {
	Sender               string  `json:"sender"`
	To                   string  `json:"to"`
	ReceiverOnOtherChain string  `json:"receiver_on_other_chain"`
	Amount               []*Coin `json:"amount"`
	Secret               []byte  `json:"secret"`
	Timestamp            int64   `json:"timestamp,string"`
	ExpireHeight         int64   `json:"expire_height,string"`
	State                string  `json:"state"`
}

func HtlcInfo(hashlock string) (result LcdHtlc, err error) {
	//htlcinfo, err := client.Htlc().QueryHTLC(hashlock)
	//if err != nil {
	//	return result, err
	//}
	//data, _ := json.Marshal(htlcinfo)
	//fmt.Println(data)
	//result = LcdHtlc{
	//	Type:htlcinfo,
	//}
	return result, nil
}
