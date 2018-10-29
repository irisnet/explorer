package utils

import (
	"strconv"
)

func ParseInt(text string) (i int64, b bool) {
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		return i, false
	}
	return i, true
}

func ParseUint(text string) (i int64, b bool) {
	i, ok := ParseInt(text)
	if ok {
		return i, i > 0
	}
	return i, ok
}
