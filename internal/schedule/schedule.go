package schedule

import (
	"days-remaining/internal/data"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func SetupScheduler(function any, bot any) gocron.Scheduler {
	s, _ := gocron.NewScheduler()

	if _, err := os.Stat(directoryPath); errors.Is(err, os.ErrNotExist) {
		createDirectory()
	}
	file, err := os.ReadFile(dataPath)
	if (err != nil) {
		log.Println(err)
	}

	data := []data.Sending{}
	json.Unmarshal(file, &data)

	for _, val := range data {
		_, _ = s.NewJob(
			gocron.DailyJob(
				1,
				gocron.NewAtTimes(
					gocron.NewAtTime(val.Hours, val.Minutes, 0),
				),
			),
			gocron.NewTask(
				function,
				bot,
				val.ChatID,
			),
		)
	}
	
	s.Start()
	return s
}

func AddSendJob(
		function any, 		//function to do in task
		bot any, 			//tg-bot sending messages
		scheduler gocron.Scheduler,
		chatID int64, 
		sendTime time.Time,
	) {
	hours, minutes := uint(sendTime.Hour()), uint(sendTime.Minute())

	file, err := os.ReadFile(dataPath)
	if (err != nil) {
		log.Println(err)
	}

	vals := []data.Sending{}
	json.Unmarshal(file, &vals)
	log.Print(vals)

	for i := 0; i < len(vals); i++ {
		if (vals[i].ChatID == chatID) {
			vals = remove(vals, i)
		}
	}	
	_, _ = scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(hours, minutes, 0),
			),
		),
		gocron.NewTask(
			function,
			bot,
			chatID,
		),
	)

	vals = append(vals, data.Sending{
		ChatID: chatID, 
		Hours: hours, 
		Minutes: minutes})
	dataBinary, err := json.MarshalIndent(vals, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	writeFile(dataBinary)
}