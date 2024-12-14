package data

var DailyText = 
`До конца университета осталось %d дней`

// TODO: Change to bytes.Buffer version
var IntroText =
`Привет! Я бот, который считает, сколько дней осталось до конца университета. Я рассылаю эту информацию раз в день. Пожалуйста, укажи, во сколько ты хотел бы получать это сообщение? Время укажи с помощью команды /set_time в формате "ЧЧ:ММ" — часы и минуты (без кавычек)
Например, "/set_time 14:10"
установит время ежедневной рассылки на 14:10` 

var SuccessChangeText = 
`Хорошо! Каждый день в %s я буду оповещать тебя о том, сколько дней осталось до конца университета!`

var FailureChangeText =
`Кажется, ты не указал время (или сделал это неправильно). Сделай это с помощью команды "/set_time ЧЧ:ММ" (часы и минуты, без кавычек)`

type Sending struct {
	ChatID	int64	`json:"ID"`
	Hours	uint	`json:"hours"`
	Minutes uint	`json:"minutes"`
}