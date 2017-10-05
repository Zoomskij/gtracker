package main

import (
	"log"
	"strconv"
	"time"

	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	var start time.Time
	var stop time.Time
	bot, err := tgbotapi.NewBotAPI("Token")
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
				start = time.Now()
				msg := tgbotapi.NewMessage(ChatID, "Started at "+strconv.Itoa(start.Hour())+":"+strconv.Itoa(start.Minute()))
				bot.Send(msg)
			}

			if Text == "/stop" {
				stop = time.Now()
				travelTime := (stop.Hour()*60 + stop.Minute()) - (start.Hour()*60 + start.Minute())
				result := "Finished at " + strconv.Itoa(stop.Hour()) + ":" + strconv.Itoa(stop.Minute()) + "\n" + "Travel time: " + strconv.Itoa(travelTime)
				msg := tgbotapi.NewMessage(ChatID, result)
				bot.Send(msg)
			}

			if Text == "/clear" {
				start = time.Now()
				stop = time.Now()
				msg := tgbotapi.NewMessage(ChatID, "Cleared")
				bot.Send(msg)
			}

		}
	}
}
