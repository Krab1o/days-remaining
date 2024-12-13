package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var text = 
`До конца университета осталось %d дней`

func createFinalDate() time.Time {
	name := "+04"
	offsetSeconds := 4 * 3600 // 4 hours in seconds
	customZone := time.FixedZone(name, offsetSeconds)
	return time.Date(2025, time.June, 25, 18, 0, 0, 0, customZone)
}

func sendMessage(bot *tgbotapi.BotAPI) {
	finalDate := createFinalDate()
	duration := finalDate.Sub(time.Now())
	text := fmt.Sprintf(
		text,
		int(duration.Hours()) / 24,
	)
	msg := tgbotapi.NewMessage(468919970, text)
	bot.Send(msg)
}

func setupScheduler(bot *tgbotapi.BotAPI) gocron.Scheduler {
	s, _ := gocron.NewScheduler()
	defer func() { _ = s.Shutdown() }()
	_, _ = s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(18, 0, 0),
			),
		),
		gocron.NewTask(
			sendMessage,
			bot,
		),
	)
	s.Start()
	return s
}

func handleUpdates(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		switch {
		case update.Message != nil:
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		case update.EditedMessage != nil:
			log.Printf("[%s] %s", update.EditedMessage.From.UserName, update.EditedMessage.Text)

			msg := tgbotapi.NewMessage(update.EditedMessage.Chat.ID, update.EditedMessage.Text + "YOU EDITED MESSAGE")
			msg.ReplyToMessageID = update.EditedMessage.MessageID

			bot.Send(msg)
		case update.MyChatMember != nil:
			//implement remembering IDs of different chatters
		}
	}
}

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

	// start the scheduler
	s := setupScheduler(bot)
	go handleUpdates(bot)
	select {} // block forever

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}