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

func SetupScheduler() gocron.Scheduler {
	s, _ := gocron.NewScheduler()

	if _, err := os.Stat(directoryPath); errors.Is(err, os.ErrNotExist) {
		createDirectory()
	}

	file, err := os.ReadFile(dataPath)
	if (err != nil) {
		log.Println(err)
	}

	mailings := []data.Mailing{}
	json.Unmarshal(file, &mailings)

	for _, mailing := range mailings {
		_, _ = s.NewJob(
			gocron.DailyJob(
				1,
				gocron.NewAtTimes(
					gocron.NewAtTime(
						uint(mailing.SendTime.Hour()), 
						uint(mailing.SendTime.Minute()), 
						0,
					),
				),
			),
			gocron.NewTask(
				dailyMessage,
				mailing,
			),
		)
	}

	s.Start()
	return s
}

func addSendJob(
		scheduler	gocron.Scheduler,
		chatID		int,
		newSendTime time.Time,
		newDate 	time.Time,
	) {
	hours, minutes := uint(newSendTime.Hour()), uint(newSendTime.Minute())

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
			scheduler.RemoveJob(updatedMailing.CronID)
			mailings = remove(mailings, i)
		}
	}

	timeChanged := true 
	dateChanged := true

	if (newSendTime == time.Time{}) {
		timeChanged = false
		newSendTime, err = time.Parse(data.TimeLayout, time.Now().Format(data.TimeLayout))
		log.Print(newSendTime)
		if (err != nil) {
			log.Println(err)
		}
	}
	if (newDate == time.Time{}) {
		dateChanged = false
		newDate, err = time.Parse(data.DateLayout, time.Now().Format(data.DateLayout))
		log.Print(newDate)
		if (err != nil) {
			log.Println(err)
		}
	}

	if (updatedMailing == data.Mailing{}) {
		updatedMailing = data.Mailing{
			ChatID: chatID,
			Date: newDate,
			SendTime: newSendTime,
		}
	} else {
		if (timeChanged) {
			updatedMailing.SendTime = newSendTime
		}
		if (dateChanged) {
			updatedMailing.Date = newDate
		}
	}

	job, _ := scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(hours, minutes, 0),
			),
		),
		gocron.NewTask(
			dailyMessage,
			data.Mailing{
				ChatID: updatedMailing.ChatID,
				SendTime: updatedMailing.SendTime,
				Date: updatedMailing.Date,
			},
		),
	)

	log.Print(updatedMailing)

	updatedMailing.CronID = job.ID()
	mailings = append(mailings, updatedMailing)

	dataBinary, err := json.MarshalIndent(mailings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	writeFile(dataBinary)
}