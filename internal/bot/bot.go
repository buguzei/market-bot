package bot

import (
	"github.com/buguzei/market-bot/internal/bot/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	msgID := make(map[int64]int)
	stage := make(map[int64]string)

	for update := range updates {
		go handlers.WorkWithUpdate(bot, update, msgID, stage)
	}
}
