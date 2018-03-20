package tools

import (
	"net/http"
	"fmt"
	"github.com/tendermint/go-wire/data"
)

func FmtOutPutResult(w http.ResponseWriter, res interface{}) error {
	json, err := data.ToJSON(res)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "%s\n", json)
	return err
}
