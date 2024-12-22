package data

type ParseMode int

const (
	None ParseMode = iota
	MarkdownV2
)

func (mode ParseMode) String() string {
	switch mode {
	case None:
		return ""
	case MarkdownV2:
		return "MarkdownV2"
	default:
		return "unknown"
	}
}

const TimeLayout = "15:04"
const DateLayout = "02.01.2006"

const URL = "https://api.telegram.org/bot"

const StartDescription = "Starts this bot and prints info about it"
const SetTimeDescription = "Sets time for an daily message receiving"
const SetDateDescription = "Sets date when some event is over"

const DailyText = 
`До конца университета осталось %d дней`

const StartText =
"Привет\\! Я бот, который считает, сколько дней осталось до конца университета\\. " +
"Я рассылаю эту информацию раз в день\\. Пожалуйста, укажи, во сколько ты хотел " +
"бы получать это сообщение? Время укажи с помощью команды `/set\\_time` в формате " +
"\"ЧЧ:ММ\" — часы и минуты \\(без кавычек\\)\n" +
"Например, `/set\\_time 14:10`\nустановит время ежедневной рассылки на 14:10"

const SuccessTimeChangeText = 
"Хорошо\\! Каждый день в %s я буду оповещать тебя о том, сколько " + 
"дней осталось до конца университета\\!"

const SuccessDateChangeText =
"Хорошо\\! Я буду считать, сколько времени осталось до даты %s"

const FailureTimeChangeText =
"Кажется, ты не указал время \\(или сделал это неправильно\\)\\. " +
"Сделай это с помощью команды `/set\\_time ЧЧ:ММ` \\(часы и минуты, без кавычек\\)"
