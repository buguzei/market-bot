package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WorkWithMsg(bot *tgbotapi.BotAPI, update tgbotapi.Update, msgID map[int64]int) {
	text := update.Message.Text
	chatID := update.Message.Chat.ID
	//userId := update.Message.From.ID
	username := update.Message.From.UserName

	switch text {
	case "/start":
		msgConfig, err := NewMsgAndMarkup(bot, chatID, fmt.Sprintf("Привет, %s. Если ты тут впервые, то вызови команду /help", username), tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("LETS'GO!", "start"))))
		msgID[chatID] = msgConfig.MessageID
		CheckErr(err)
	case "/help":
		_, err := NewMsg(bot, chatID, "Вот все мои команды:\n")
		CheckErr(err)
	}
}
