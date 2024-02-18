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

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "/forecast":
			msg.Text = "Please enter the city:"
			msg.ReplyMarkup = configs.NumericKeyboardCity
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			waiting_city = true
			continue
		case "/info":
			msg.Text = "Hello, there is my options: \n/info\n/forecast"
		case "/start":
			msg.Text = "Hello, there is my options: \n/info\n/forecast"
		}

		if waiting_city {
			reply, err := handlers.MakeRequestToOpWether(update.Message.Text)
			if err == nil {
				msg.Text = reply
			} else {
				msg.Text = err.Error()
			}
			waiting_city = false
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
