package data

import "time"

type Mailing struct {
	ChatID	int			`json:"ID"`
	SendTime 	time.Time	`json:"time"`
	Date	time.Time	`json:"date"`
}

type Command struct{
	Title 	string	`json:"command"`
	Desc	string	`json:"description"`
}

type Update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

type APIResponseUpdates struct {
	Ok     bool 	`json:"ok"`
	Result []Update `json:"result"`
}

type Config struct {
	Offset		int
	Timeout		int
}