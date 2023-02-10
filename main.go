package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sixsat/market-report/bot"
)

func main() {
	session, err := bot.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot is now running.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	err = session.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("bye~")
}
