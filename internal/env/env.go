package env

import (
	"log"
	"os"
)

// EnvGetOrDefault ...
func EnvGetOrDefault(key string, alt string) string {
	val, presented := os.LookupEnv(key)
	log.Println(val)
	if !presented {
		return alt
	}
	return val
}
