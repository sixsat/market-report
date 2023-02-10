package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(ds *discordgo.Session, dm *discordgo.MessageCreate) {
	if botMessage(ds, dm) {
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

func botMessage(ds *discordgo.Session, dm *discordgo.MessageCreate) bool {
	return dm.Author.ID == ds.State.User.ID
}
