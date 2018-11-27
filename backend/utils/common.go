package utils

import (
	"strconv"
	"time"
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

func FmtUTCTime(time time.Time) string {
	if time.IsZero() {
		return ""
	}
	return time.UTC().Format("2006/01/02 15:04:05+UTC")
}
