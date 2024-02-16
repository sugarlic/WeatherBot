package configs

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var NumericKeyboard = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("Hello"),
    ),
)