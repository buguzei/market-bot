package config

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Пополнить", "replenish"),
		tgbotapi.NewInlineKeyboardButtonData("Разместить товар", "replace"),
		tgbotapi.NewInlineKeyboardButtonData("Выбрать город", "city"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Правила", "rules"),
		tgbotapi.NewInlineKeyboardButtonData("Поддержка", "support"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Каталог", "catalog"),
		tgbotapi.NewInlineKeyboardButtonData("Информация", "info"),
	),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Отзывы", "reviews")),
)

var Currencies = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("BTC", "bitcoin"),
		tgbotapi.NewInlineKeyboardButtonData("LTC", "litecoin"),
		tgbotapi.NewInlineKeyboardButtonData("USDT", "usdt"),
	))

var Cities = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Москва", "moscow"),
		tgbotapi.NewInlineKeyboardButtonData("Санкт-Петеребург", "peter"),
		tgbotapi.NewInlineKeyboardButtonData("Варшава", "varshava"),
		tgbotapi.NewInlineKeyboardButtonData("Гданьск", "gdansk"),
	))
