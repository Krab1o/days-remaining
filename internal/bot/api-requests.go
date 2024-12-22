package bot

import (
	"days-remaining/internal/data"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getUpdates(conf *data.Config) ([]data.Update, error) {
	getUpdatesRequest, err := http.NewRequest("GET", data.URL + os.Getenv("TELEGRAM_TOKEN") + "/getUpdates", nil)
	getUpdatesRequestQuers := getUpdatesRequest.URL.Query()
	getUpdatesRequestQuers.Add("timeout", strconv.Itoa(conf.Timeout))
	getUpdatesRequestQuers.Add("offset", strconv.Itoa(conf.Offset))
	getUpdatesRequest.URL.RawQuery = getUpdatesRequestQuers.Encode()

	getUpdatesResponse, err := http.DefaultClient.Do(getUpdatesRequest)
	if err != nil {
		return []data.Update{}, err
	}
	defer getUpdatesResponse.Body.Close()
	
	response := data.APIResponseUpdates{}
	json.NewDecoder(getUpdatesResponse.Body).Decode(&response)

	log.Println(response.Result)
	return response.Result, nil
}

func GetUpdatesChan(config *data.Config) chan data.Update {
	channelSize := 100
	ch := make(chan data.Update, channelSize)
	
	go func() {
		for {
			updates, err := getUpdates(config)
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)
				continue
			}
			
			for _, update := range updates {
				if update.UpdateID >= config.Offset {
					config.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func sendMessage(chatID int, text string, parseMode data.ParseMode) {
	sendMessageRequest, err := http.NewRequest("POST", data.URL + os.Getenv("TELEGRAM_TOKEN") + "/sendMessage", nil)
	if err != nil {
		log.Println(err)
	}
	
	queryParams := sendMessageRequest.URL.Query()
	queryParams.Add("chat_id", strconv.Itoa(chatID))
	queryParams.Add("text", text)
	if parseMode != data.None {
		queryParams.Add("parse_mode", parseMode.String())	
	}
	sendMessageRequest.URL.RawQuery = queryParams.Encode()
	response, err := http.DefaultClient.Do(sendMessageRequest)
	test, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(test))
}

func setCommands() {
	setMyCommandsRequest, err := http.NewRequest("POST", data.URL + os.Getenv("TELEGRAM_TOKEN") + "/setMyCommands", nil)
	if err != nil {
		log.Println(err)
	}

	commands := []data.Command{
		{Title: "/start", Desc:	data.StartDescription},
		{Title: "/set_time", Desc: data.SetTimeDescription},
		{Title: "/set_date", Desc: data.SetDateDescription},
	}

	jsonString, err := json.MarshalIndent(commands, "", "    ")

	queryParams := setMyCommandsRequest.URL.Query()
	queryParams.Add("commands", string(jsonString))
	setMyCommandsRequest.URL.RawQuery = queryParams.Encode()
	http.DefaultClient.Do(setMyCommandsRequest)
}