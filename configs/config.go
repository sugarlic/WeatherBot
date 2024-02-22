package configs

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WaitingStates struct {
	Waiting_city   bool
	Waiting_coords bool
}

var Start time.Time

var NumericKeyboardCity = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Kazan"),
		tgbotapi.NewKeyboardButton("Moscow"),
		tgbotapi.NewKeyboardButton("Petersburg"),
		tgbotapi.NewKeyboardButton("Novgorod"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Samara"),
		tgbotapi.NewKeyboardButton("Omsk"),
		tgbotapi.NewKeyboardButton("Vladimir"),
		tgbotapi.NewKeyboardButton("Ufa"),
	),
)
