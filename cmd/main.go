package main

import (
	"github.com/buguzei/market-bot/internal/bot"
	"github.com/buguzei/market-bot/internal/server"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")

	Bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	server.MakeWalletAddress(server.USDT)

	bot.StartBot(Bot)
}
