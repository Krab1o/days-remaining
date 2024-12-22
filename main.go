package main

import (
	"days-remaining/internal/bot"
	"days-remaining/internal/data"
	"log"
	"regexp"

	"github.com/joho/godotenv"
)

func parseCommand(message string) (string, bool) {
	rg, err := regexp.Compile(`^\/[a-zA-Z_]{1,}\b`)
	if err != nil {
		log.Println(err)
	}
	if index := rg.FindStringIndex(message); index != nil {
		return message[index[0]:index[1]], true
	}
	return data.None.String(), false
}

func main() {
	err := godotenv.Load(".env")
  	if err != nil {
    	log.Fatalf("Error loading .env file")
  	}

	conf := &data.Config{
		Offset: 0,
		Timeout: 60,
	}

	bot.InitBot()
	s := bot.SetupScheduler()
	updates := bot.GetUpdatesChan(conf)

	for update := range updates {
		if command, ok := parseCommand(update.Message.Text); ok != false {
			log.Println(command)
			switch command {
			case "/start":
				bot.Start(update)
			case "/set_time":
				bot.SetTime(s, update)
			case "/set_date":
				bot.SetDate(update)
			default:
				// bot.SendMessage(update.Message.Chat.ID, update.Message.Text, "")
			}
		} else {
			bot.SendMessage(update, "Я понимаю только команды :(")
		}
	}
}