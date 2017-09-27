package main

import (
	"log"

	"github.com/Syfaro/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
