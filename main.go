package main

import (
	"log"
	"os"

	// "example.com/m/configs"

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

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Text {
		case "/Hello":
			msg.Text = "Hello, I am test bot!"
		default:
			msg.Text = "Hello, there is my options: \n/Hello\n/Broadcast"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
