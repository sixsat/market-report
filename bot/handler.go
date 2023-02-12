package bot

import (
	"fmt"
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
	case "!summary":
		_, err := ds.ChannelMessageSend(dm.ChannelID, summaryMessage(""))
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
	res := set.GetSummary(market)

	// Index
	fmt.Fprintf(&b,
		`%s %s %s
%.2f %.2f (%.2f%%)
Val. %.2f M Vol. %.2f M
`,
		res.Index.NameEN, res.Index.Level, res.Index.MarketStatus,
		res.Index.Last, res.Index.Change, res.Index.PercentChange,
		res.Index.Value/1_000_000, res.Index.Volume/1_000_000,
	)

	// Investor summary
	fmt.Fprintf(&b,
		`
INSTITUTION %.2f M
PROP %.2f M
FOREIGN %.2f M
LOCAL %.2f M
`,
		res.InvestorSummary.Investors[0].NetValue/1_000_000,
		res.InvestorSummary.Investors[1].NetValue/1_000_000,
		res.InvestorSummary.Investors[2].NetValue/1_000_000,
		res.InvestorSummary.Investors[3].NetValue/1_000_000,
	)

	// Most active value
	b.WriteString("\nMost Active Value\n")
	for i := range res.Rankings[0].Stocks {
		fmt.Fprintf(&b, "%s %.2f %.2f (%.2f%%)\n",
			res.Rankings[0].Stocks[i].Symbol, res.Rankings[0].Stocks[i].Last, res.Rankings[0].Stocks[i].Change, res.Rankings[0].Stocks[i].PercentChange)
	}

	// Top gainer
	b.WriteString("\nTop Gainer\n")
	for i := range res.Rankings[1].Stocks {
		fmt.Fprintf(&b, "%s %.2f %.2f (%.2f%%)\n",
			res.Rankings[1].Stocks[i].Symbol, res.Rankings[1].Stocks[i].Last, res.Rankings[1].Stocks[i].Change, res.Rankings[1].Stocks[i].PercentChange)
	}

	// Top loser
	b.WriteString("\nTop Loser\n")
	for i := range res.Rankings[2].Stocks {
		fmt.Fprintf(&b, "%s %.2f %.2f (%.2f%%)\n",
			res.Rankings[2].Stocks[i].Symbol, res.Rankings[2].Stocks[i].Last, res.Rankings[2].Stocks[i].Change, res.Rankings[2].Stocks[i].PercentChange)
	}

	b.WriteString("\nSource: Settrade")

	return b.String()
}
