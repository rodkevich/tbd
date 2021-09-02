package env

import (
	"log"
	"os"
)

// GetOrDefault for local usage or whatever =)
func GetOrDefault(key string, alt string) string {
	val, presented := os.LookupEnv(key)
	log.Println(val)
	if !presented {
		return alt
	}
	return val
}
