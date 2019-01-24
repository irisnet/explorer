package logger

import (
	"os"
)

type Config struct {
	Filename          string
	MaxSize           int
	MaxAge            int
	Compress          bool
	EnableAtomicLevel bool
}

var (
	conf = Config{
		Filename:          os.ExpandEnv("$HOME/explorer/explorer.log"),
		MaxSize:           20,
		MaxAge:            7,
		Compress:          true,
		EnableAtomicLevel: true,
	}
)
