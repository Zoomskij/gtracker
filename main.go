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

	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-upd:
			ChatID := update.Message.Chat.ID
			Text := update.Message.Text

			if Text == "/go" {
				msg := tgbotapi.NewMessage(ChatID, "go")
				bot.Send(msg)
			}

		}
	}
}
