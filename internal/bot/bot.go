package bot

import (
	"days-remaining/internal/data"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func InitBot() {
	setCommands()
}

func SendMessage(update data.Update, msg string) {
	sendMessage(update.Message.Chat.ID, msg, data.None)
}

func SetDate(sc gocron.Scheduler, update data.Update) {
	var msg string
	parsedDate, ok := parseDate(update.Message.Text); if ok {
		log.Println(parsedDate)
		addSendJob(sc, update.Message.Chat.ID, time.Time{}, parsedDate)
		msg = fmt.Sprintf(
			data.SuccessDateChangeText, 
			fmt.Sprintf("%02d\\.%02d\\.%04d",
				parsedDate.Day(),
				parsedDate.Month(),
				parsedDate.Year(), 
			),
		)
	} else {
		msg = data.FailureDateChangeText
	}
	sendMessage(update.Message.Chat.ID, msg, data.MarkdownV2)
}

func SetTime(sc gocron.Scheduler, update data.Update) {
	var msg string
	parsedTime, ok := parseTime(update.Message.Text); if ok {
		log.Println(parsedTime)
		addSendJob(sc, update.Message.Chat.ID, parsedTime, time.Time{})
		msg = fmt.Sprintf(
			data.SuccessTimeChangeText, 
			fmt.Sprintf("%02d:%02d",
				parsedTime.Hour(),
				parsedTime.Minute(), 
			),
		)
	} else {
		msg = data.FailureTimeChangeText
	}
	sendMessage(update.Message.Chat.ID, msg, data.MarkdownV2)
}

func Start(update data.Update) {
	sendMessage(update.Message.Chat.ID, data.StartText, data.MarkdownV2)
}