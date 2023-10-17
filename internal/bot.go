package internal

import (
	"github.com/buguzei/market-bot/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	msgID := make(map[int64]int)

	for update := range updates {
		handlers.WorkWithUpdate(bot, update, msgID)
	}
}
