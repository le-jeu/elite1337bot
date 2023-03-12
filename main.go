package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/le-jeu/elite1337bot/tgbot"
)

func main() {
	// Starting bot
	bot, err := tgbot.NewTelegramBot(os.Getenv("TGBOT_TOKEN"))
	if err != nil {
		log.Fatal("Couldn't open the Telegram Bot. Error: " + err.Error())
	}

	err = bot.Start()
	if err != nil {
		log.Fatal("Couldn't open the Telegram Bot. Error: " + err.Error())
	}

	// Starting the maintenace
	//maintenance.Start(bot)

	// Prevent main from returning immediately. Wait for interrupt.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	<-signalChan
	log.Println("Application closed by user.")
}
