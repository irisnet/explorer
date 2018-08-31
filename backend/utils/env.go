package utils

import (
	"os"
	"fmt"
)

func GetEnv(key string,result string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(key, "=", value)
		return value
	} else {
		return result
	}
}
