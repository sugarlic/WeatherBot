package main

import (
	"log"
	"os"

	"example.com/m/configs"
	"example.com/m/handlers"
	"example.com/m/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := utils.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	waiting_states := configs.WaitingStates{
		Waiting_city:   false,
		Waiting_coords: false}

	const def_reply = "Hello, there is my options: \n/info\n/forecast\n/forecast_by_coords"

	for update := range updates {

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
			waiting_states.Waiting_city = true
			continue
		case "forecast_by_coords":
			msg.Text = "Please enter the coords in format: \"lat lon\""
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			waiting_states.Waiting_coords = true
			continue
		case "info":
			msg.Text = def_reply
		}

		if waiting_states.Waiting_city {
			handlers.ReplyToCityForecast(db, update, &msg)
			waiting_states.Waiting_city = false
		}

		if waiting_states.Waiting_coords {
			handlers.ReplyToCoordsForecast(update, &msg)
			waiting_states.Waiting_coords = false
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
