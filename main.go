package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sixsat/market-report/bot"
)

func main() {
	err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot is now running.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	err = bot.Stop()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("bye~")
}
