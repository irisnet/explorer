package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func RoundFloat(num float64, bit int) (i float64, b bool) {
	format := "%" + fmt.Sprintf("0.%d", bit) + "f"
	s := fmt.Sprintf(format, num)
	i, err := strconv.ParseFloat(s, 0)
	if err != nil {
		return i, false
	}
	return i, true
}

func FmtTimeToString(time time.Time) string {
	if time.IsZero() {
		return ""
	}
	return time.String()
}

func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, num)
	return buffer.Bytes()
}
