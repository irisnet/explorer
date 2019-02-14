package utils

import "time"

const (
	DateFmtYYYYMMDD       = "2006-01-02"
	DateFmtYYYYMMDDHHmmss = "2006-01-02 15:04:05"
)

const (
	_ Unit = iota
	Day
	Hour
	Min
	Sec
)

type Unit int

func TruncateTime(t time.Time, unit Unit) time.Time {
	switch unit {
	case Day:
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	case Hour:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	case Min:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	case Sec:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
	}
	panic("not exist unit")
}

func ParseDuration(num int, unit Unit) time.Duration {
	switch unit {
	case Day:
		return time.Duration(num*24) * time.Hour
	case Hour:
		return time.Duration(num) * time.Hour
	case Min:
		return time.Duration(num) * time.Minute
	case Sec:
		return time.Duration(num) * time.Second
	}
	panic("not exist unit")
}

func FmtTime(t time.Time, fmt string) string {
	return t.Format(fmt)
}
