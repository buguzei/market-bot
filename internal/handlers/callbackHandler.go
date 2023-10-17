package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WorkWithCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, msgID map[int64]int) {
	data := update.CallbackData()
	userID := update.CallbackQuery.From.ID
	//userId := update.Message.From.ID
	//username := update.Message.From.UserName

	switch data {
	case "start":
		err := NewEditMsgAndMarkup(bot, userID, msgID[userID], "Выберите свою роль:", tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Я селлер", "imseller"),
				tgbotapi.NewInlineKeyboardButtonData("Я покупатель", "imbuyer"),
			)), true)
		CheckErr(err)
	}
}
