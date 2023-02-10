package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/sixsat/market-report/config"
)

// Start creates a new Discord session, setups an event handler, and opens a
// websocket connection to Discord. Use Session.Close() to close the connection.
func Start() (*discordgo.Session, error) {
	cfg := config.New()

	session, err := discordgo.New(cfg.Discord.BotToken)
	if err != nil {
		return nil, fmt.Errorf("creating discord session: %w", err)
	}

	session.AddHandler(messageHandler)

	err = session.Open()
	if err != nil {
		return nil, fmt.Errorf("opening websocket connection: %w", err)
	}

	return session, nil
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
