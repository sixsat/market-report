package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sixsat/market-report/config"
)

var session *discordgo.Session

// Start creates a new Discord session, setups an event handler, and opens a
// websocket connection to Discord.
func Start() error {
	var err error
	session, err = discordgo.New("Bot " + config.Getenv("BOT_TOKEN"))
	if err != nil {
		return fmt.Errorf("creating discord session: %w", err)
	}

	session.AddHandler(messageHandler)

	err = session.Open()
	if err != nil {
		return fmt.Errorf("opening websocket connection: %w", err)
	}

	return nil
}

// Stop closes a websocket and stops all listening/heartbeat goroutines.
func Stop() error {
	err := session.Close()
	if err != nil {
		return fmt.Errorf("closing websocket: %w", err)
	}

	return nil
}
