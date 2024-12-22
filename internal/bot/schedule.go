package bot

import (
	"days-remaining/internal/data"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func remove(s []data.Mailing, i int) []data.Mailing {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func updateTime(chatID int, newSendTime time.Time) data.Mailing {
	file, err := os.ReadFile(dataPath)
	if (err != nil) {
		log.Println(err)
	}

	mailings := []data.Mailing{}
	json.Unmarshal(file, &mailings)
	var updatedMailing data.Mailing

	for i := 0; i < len(mailings); i++ {
		if (mailings[i].ChatID == chatID) {
			updatedMailing = mailings[i]
			mailings = remove(mailings, i)
		}
	}

	if (updatedMailing == data.Mailing{}) {
		currentDate, err := time.Parse(data.DateLayout, time.Now().String())
		if err != nil {
			log.Println(err)
		}
		updatedMailing = data.Mailing{
			ChatID: chatID,
			Date: currentDate,
			SendTime: newSendTime,
		}
	} else {
		updatedMailing.SendTime = newSendTime
	}

	mailings = append(mailings, updatedMailing)

	dataBinary, err := json.MarshalIndent(mailings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	writeFile(dataBinary)

	return updatedMailing
}

func updateDate(chatID int, newDate time.Time) {
	file, err := os.ReadFile(dataPath)
	if (err != nil) {
		log.Println(err)
	}

	mailings := []data.Mailing{}
	json.Unmarshal(file, &mailings)
	var updatedMailing data.Mailing

	for i := 0; i < len(mailings); i++ {
		if (mailings[i].ChatID == chatID) {
			updatedMailing = mailings[i]
			mailings = remove(mailings, i)
		}
	}

	if (updatedMailing == data.Mailing{}) {
		currentTime, err := time.Parse(data.TimeLayout, time.Now().String())
		if err != nil {
			log.Println(err)
		}
		updatedMailing = data.Mailing{
			ChatID: chatID,
			Date: newDate,
			SendTime: currentTime,
		}
	} else {
		updatedMailing.Date = newDate
	}

	mailings = append(mailings, updatedMailing)

	dataBinary, err := json.MarshalIndent(mailings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	writeFile(dataBinary)
}

func SetupScheduler() gocron.Scheduler {
	s, _ := gocron.NewScheduler()

	if _, err := os.Stat(directoryPath); errors.Is(err, os.ErrNotExist) {
		createDirectory()
	}

	file, err := os.ReadFile(dataPath)
	if (err != nil) {
		log.Println(err)
	}

	mailing := []data.Mailing{}
	json.Unmarshal(file, &mailing)

	for _, val := range mailing {
		_, _ = s.NewJob(
			gocron.DailyJob(
				1,
				gocron.NewAtTimes(
					gocron.NewAtTime(
						uint(val.SendTime.Hour()), 
						uint(val.SendTime.Minute()), 
						0,
					),
				),
			),
			gocron.NewTask(
				sendDaily,
				val,
			),
		)
	}

	s.Start()
	return s
}

func addSendJob(
		scheduler gocron.Scheduler,
		chatID int,
		newSendTime time.Time,
	) {
	hours, minutes := uint(newSendTime.Hour()), uint(newSendTime.Minute())

	updatedMailing := updateTime(chatID, newSendTime)

	_, _ = scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(hours, minutes, 0),
			),
		),
		gocron.NewTask(
			sendDaily,
			updatedMailing,
		),
	)

	
}