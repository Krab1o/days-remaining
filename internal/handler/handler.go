package handler

import (
	"days-remaining/internal/data"
	"fmt"
	"log"
	"strings"

	"regexp"
	"time"

	sc "days-remaining/internal/schedule"

	"github.com/go-co-op/gocron/v2"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var finalDate = time.Date(2025, time.June, 25, 18, 0, 0, 0, time.FixedZone("+04", 4 * 3600))

func parseTime(msg string) (time.Time, bool) {
	expr, err := regexp.Compile(`\/set_time\s\b\d{1,2}:\d{2}\b`)
	if (err != nil) {
		log.Println(err)
	}
	
	if index := expr.FindStringIndex(msg); index != nil {
		//We take second argument right after "/set_time"
		sendTime := strings.Split(msg[index[0]:index[1]], " ")[1]
		msg = fmt.Sprintf(data.SuccessChangeText, sendTime)
		if parsedTime, err := time.Parse("15:04", sendTime); err != nil {
			log.Println(err)
		} else {
			return parsedTime, true
		}
	}
	return time.Time{}, false
}

func SendMessage(bot *tgbotapi.BotAPI, chatID int64) {
	duration := finalDate.Sub(time.Now())
	text := fmt.Sprintf(
		data.DailyText,
		int(duration.Hours()) / 24,
	)
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
}

func HandleUpdates(bot *tgbotapi.BotAPI, scheduler gocron.Scheduler) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		var msg tgbotapi.MessageConfig
		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, data.IntroText)
			case "set_time":
				parsedTime, ok := parseTime(update.Message.Text); if ok {
					sc.AddSendJob(SendMessage, bot, scheduler, update.Message.Chat.ID, parsedTime)
					msg = tgbotapi.NewMessage(
						update.Message.Chat.ID, 
						fmt.Sprintf(
							data.SuccessChangeText, 
							fmt.Sprintf("%d:%d",
								parsedTime.Hour(),
								parsedTime.Minute(), 
							),
						),
					)
				} else {
					msg = tgbotapi.NewMessage(
						update.Message.Chat.ID,
						data.FailureChangeText,
					)
				}
			default:
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, data.FailureChangeText)
			}
			bot.Send(msg)
		}
	}
}