package config

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Пополнить", "replenish")),
)
