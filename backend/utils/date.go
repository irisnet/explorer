package utils

import "time"

const (
	DateFmtYYYYMMDD       = "2006-01-02"
	DateFmtYYYYMMDDHHmmss = "2006-01-02 15:04:05"
)

const (
	_ Prec = iota
	Day
	Hour
	Min
	Sec
)

type Prec int

func TruncateTime(t time.Time, prec Prec) time.Time {
	switch prec {
	case Day:
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	case Hour:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	case Min:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	case Sec:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func FmtTime(t time.Time, fmt string) string {
	return t.Format(fmt)
}
