package bot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sixsat/market-report/set"
)

func messageHandler(ds *discordgo.Session, dm *discordgo.MessageCreate) {
	if isBot(ds, dm) {
		return
	}

	switch dm.Content {
	case "!bot":
		log.Println("Channel ID:", dm.ChannelID)
		_, err := ds.ChannelMessageSend(dm.ChannelID, "Hello!")
		if err != nil {
			log.Println(err)
			return
		}
	case "!set":
		_, err := ds.ChannelMessageSend(dm.ChannelID, summaryMessage("set"))
		if err != nil {
			log.Println(err)
			return
		}
	case "!mai":
		_, err := ds.ChannelMessageSend(dm.ChannelID, summaryMessage("mai"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func isBot(ds *discordgo.Session, dm *discordgo.MessageCreate) bool {
	return dm.Author.ID == ds.State.User.ID
}

func summaryMessage(market string) string {
	var b strings.Builder
	res := set.GetPrettySummary(market)

	b.WriteString("```\n")

	// Index
	b.WriteString(
		res.Index.NameEN + " Index on " + res.Index.MarketDateTime + " " + res.Index.MarketStatus + "\n" +
			res.Index.Last + " " + res.Index.Change + " (" + res.Index.PercentChange + "%)\n" +
			"Val. " + res.Index.Value + " M฿, Vol. " + res.Index.Volume + " M\n\n",
	)

	// Investor summary
	// TODO: Add YTD net value
	for i := range res.InvestorSummary.Investors {
		b.WriteString(
			res.InvestorSummary.Investors[i].Type + " " + res.InvestorSummary.Investors[i].NetValue + " M฿\n",
		)
	}

	// Most active value
	b.WriteString("\nMost Active Value\n")
	for i := range res.Rankings[0].Stocks {
		b.WriteString(
			res.Rankings[0].Stocks[i].Symbol + " " +
				res.Rankings[0].Stocks[i].Last + " " +
				res.Rankings[0].Stocks[i].Change + " (" +
				res.Rankings[0].Stocks[i].PercentChange + "%)\n",
		)
	}

	// Top gainer
	b.WriteString("\nTop Gainer\n")
	for i := range res.Rankings[1].Stocks {
		b.WriteString(
			res.Rankings[1].Stocks[i].Symbol + " " +
				res.Rankings[1].Stocks[i].Last + " " +
				res.Rankings[1].Stocks[i].Change + " (" +
				res.Rankings[1].Stocks[i].PercentChange + "%)\n",
		)
	}

	// Top loser
	b.WriteString("\nTop Loser\n")
	for i := range res.Rankings[2].Stocks {
		b.WriteString(
			res.Rankings[2].Stocks[i].Symbol + " " +
				res.Rankings[2].Stocks[i].Last + " " +
				res.Rankings[2].Stocks[i].Change + " (" +
				res.Rankings[2].Stocks[i].PercentChange + "%)\n",
		)
	}

	b.WriteString("\nSource: Settrade\n```")

	return b.String()
}
