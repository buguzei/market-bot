package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WorkWithUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, msgID map[int64]int) {

	if update.CallbackQuery != nil {
		WorkWithCallback(bot, update, msgID)
		return
	}

	if update.Message != nil {
		if update.Message.Chat.IsPrivate() {
			WorkWithMsg(bot, update, msgID)
		}
		return
	}
}
