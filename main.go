package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/sixsat/market-report/config"
)

func main() {
	cfg := config.New()

	dg, err := discordgo.New(cfg.Discord.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(messageHandler)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dg.Close()

	log.Println("Bot is now running.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, os.Interrupt)
	<-sc

	log.Println("bye~")
}

func messageHandler(ds *discordgo.Session, dm *discordgo.MessageCreate) {
	// Ignore bot messages
	if dm.Author.ID == ds.State.User.ID {
		return
	}

	switch dm.Content {
	case "!bot":
		log.Println("Channel ID:", dm.ChannelID)
		_, err := ds.ChannelMessageSend(dm.ChannelID, "Hello!")
		if err != nil {
			log.Println(err)
		}
	}
}
