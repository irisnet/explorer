package utils

import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
)

func TransferPage(r *http.Request) (int, int) {
	args := mux.Vars(r)
	page := args["page"]
	size := args["size"]
	iPage := 0
	iSize := 20
	if page != "" {
		if iPage, _ = strconv.Atoi(page); iPage < 1 {
			iPage = 1
		}

	}
	if size != "" {
		if iSize, _ = strconv.Atoi(size); iSize < 1 {
			iSize = 20
		}
	}
	return iPage, iSize
}
