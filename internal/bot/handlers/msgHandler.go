package handlers

import (
	"fmt"
	"github.com/buguzei/market-bot/config"
	"github.com/buguzei/market-bot/internal/server"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func WorkWithMsg(bot *tgbotapi.BotAPI, update tgbotapi.Update, msgID map[int64]int, stage map[int64]string) {
	text := update.Message.Text
	chatID := update.Message.Chat.ID
	//userId := update.Message.From.ID
	username := update.Message.From.UserName

	switch stage[chatID] {
	case "replenish":
		intText, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			return
		}
		server.Amount = intText

	default:
		switch text {
		case "/start":
			msgConfig, err := NewMsgAndMarkup(bot, chatID, fmt.Sprintf("Привет, %s!\nВ этом боте ты можешь", username), config.MainKB)
			msgID[chatID] = msgConfig.MessageID
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
