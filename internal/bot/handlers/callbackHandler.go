package handlers

import (
	"github.com/buguzei/market-bot/config"
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
		err := NewEditMsgAndMarkup(bot, userID, msgID[userID], "Выберите валюту:", config.Currencies, true)
		if err != nil {
			log.Println(err)
			return
		}
	case "city":
		err := NewEditMsgAndMarkup(bot, userID, msgID[userID], "Введите сумму:", config.Cities, false)
		if err != nil {
			log.Println(err)
			return
		}
	case "replace":
		err := NewEditMsgAndMarkup(bot, userID, msgID[userID], "Введите описание товара", config.Cities, false)
		if err != nil {
			log.Println(err)
			return
		}
		stage[userID] = "replace"
	case "bitcoin":
		err := NewEditMsgAndMarkup(bot, userID, msgID[userID], "Введите сумму:", tgbotapi.NewInlineKeyboardMarkup(), false)
		if err != nil {
			log.Println(err)
			return
		}

		stage[userID] = "replenish"
	}
}
