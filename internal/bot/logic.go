package bot

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"days-remaining/internal/data"
)

func parseDate(msg string) (time.Time, bool) {
	expr, err := regexp.Compile(`\/set_date\s\d{1,2}\.\d{1,2}\.\d{4}`)
	if (err != nil) {
		log.Println(err)
	}
	
	if index := expr.FindStringIndex(msg); index != nil {
		//We take second argument right after "/set_date"
		sendDate := strings.Split(msg[index[0]:index[1]], " ")[1]
		if parsedDate, err := time.Parse(data.DateLayout, sendDate); err != nil {
			log.Println(err)
		} else {
			return parsedDate, true
		}
	}
	return time.Time{}, false
}

func parseTime(msg string) (time.Time, bool) {
	expr, err := regexp.Compile(`\/set_time\s\b\d{1,2}:\d{2}\b`)
	if (err != nil) {
		log.Println(err)
	}
	
	if index := expr.FindStringIndex(msg); index != nil {
		//We take second argument right after "/set_time"
		sendTime := strings.Split(msg[index[0]:index[1]], " ")[1]
		if parsedTime, err := time.Parse(data.TimeLayout, sendTime); err != nil {
			log.Println(err)
		} else {
			return parsedTime, true
		}
	}
	return time.Time{}, false
}

func dailyMessage(mailing data.Mailing) {
	duration := mailing.Date.Sub(time.Now())
	msg := fmt.Sprintf(
		data.DailyText,
		int(duration.Hours()) / 24,
	)
	sendMessage(mailing.ChatID, msg, data.None)
}

