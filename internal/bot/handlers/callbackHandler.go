package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func WorkWithCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, msgID map[int64]int, stage map[int64]string) {
	data := update.CallbackData()
	userID := update.CallbackQuery.From.ID
	//userId := update.Message.From.ID
	//username := update.Message.From.UserName

	switch data {
	case "replenish":
		err := NewEditMsgAndMarkup(bot, userID, msgID[userID], "Выберите валюту:", tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("BTC", "bitcoin"),
				tgbotapi.NewInlineKeyboardButtonData("LTC", "litecoin"),
				tgbotapi.NewInlineKeyboardButtonData("USDT", "usdt"),
			)), true)
		if err != nil {
			log.Println(err)
			return
		}
	case "bitcoin":
		err := NewEditMsgAndMarkup(bot, userID, msgID[userID], "Введите сумму:", tgbotapi.NewInlineKeyboardMarkup(), false)
		if err != nil {
			log.Println(err)
			return
		}

		stage[userID] = "replenish"
	}
}
