package lcd

import "fmt"


type HtlcData struct {
	Sender       string `json:"sender"`
	To           string `json:"to"`
	Amount       []Coin `json:"amount"`
	Secret       []byte `json:"secret"`
	ExpireHeight uint64 `json:"expire_height,string"`
	State        string `json:"state"`
}

func HtlcInfo(hashlock string) (result HtlcData, err error) {
	htlcinfo, err := client.Htlc().QueryHtlc(hashlock)
	if err != nil {
		return result, err
	}
	result = HtlcData{
		Sender:       htlcinfo.Sender,
		To:           htlcinfo.To,
		Amount:       LoadCoins(htlcinfo.Amount),
		Secret:       []byte(htlcinfo.Secret),
		State:        fmt.Sprint(htlcinfo.State),
		ExpireHeight: htlcinfo.ExpirationHeight,
	}
	return result, nil
}
