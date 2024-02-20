package main

import (
	"log"
	"os"

	"example.com/m/configs"
	"example.com/m/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	waiting_city := false
	waiting_coords := false

	for update := range updates {
		def_reply := "Hello, there is my options: \n/info\n/forecast\n/forecast_by_coords"

		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, def_reply)

		switch update.Message.Command() {
		case "forecast":
			msg.Text = "Please enter the city:"
			msg.ReplyMarkup = configs.NumericKeyboardCity
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			waiting_city = true
			continue
		case "forecast_by_coords":
			msg.Text = "Please enter the coords in format: \"lat lon\""
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			waiting_coords = true
			continue
		case "info":
			msg.Text = def_reply
		}

		if waiting_city {
			reply, err := handlers.MakeRequestByCity(update.Message.Text)
			if err == nil {
				msg.Text = reply
			} else {
				msg.Text = err.Error()
			}
			waiting_city = false
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if waiting_coords {
			reply, err := handlers.MakeRequestByCoords(update.Message.Text)
			if err == nil {
				msg.Text = reply
			} else {
				msg.Text = err.Error()
			}
			waiting_coords = false
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
