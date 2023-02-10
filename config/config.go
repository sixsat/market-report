package config

import (
	"log"
	"os"
)

// Getenv uses os.Getenv, but exit the program on empty value.
func Getenv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("env %s is empty", key)
	}

	return val
}
