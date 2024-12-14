package main

import (
	"days-remaining/internal/handler"
	"log"
	"os"

	sc "days-remaining/internal/schedule"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
  	if err != nil {
    	log.Fatalf("Error loading .env file")
  	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	
	log.Printf("Authorized on account %s", bot.Self.UserName)

	s := sc.SetupScheduler(handler.SendMessage, bot)
	defer func() { _ = s.Shutdown() }()
	handler.HandleUpdates(bot, s)
}