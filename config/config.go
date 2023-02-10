package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file:", err)
	}
}

type c struct {
	Discord discord
}

type discord struct {
	BotToken  string
	ChannelID string
}

func New() *c {
	return &c{
		discord{
			BotToken:  "Bot " + getEnv("BOT_TOKEN"),
			ChannelID: getEnv("CHANNEL_ID"),
		},
	}
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("env %s is empty", key)
	}

	return val
}
