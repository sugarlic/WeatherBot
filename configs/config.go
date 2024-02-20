package configs

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
