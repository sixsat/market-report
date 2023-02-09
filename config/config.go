package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
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
			BotToken:  "Bot " + os.Getenv("BOT_TOKEN"),
			ChannelID: os.Getenv("CHANNEL_ID"),
		},
	}
}
