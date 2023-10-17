package main

import (
	"github.com/buguzei/market-bot/internal"
	"github.com/buguzei/market-bot/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("config/.env")
	handlers.CheckErrWithText(err, "Error loading .env file")

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	handlers.CheckErr(err)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	internal.Start(bot)
}
