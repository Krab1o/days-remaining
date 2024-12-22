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
"Я рассылаю эту информацию раз в день\\.\n\n" + 
"С помощью команды `/set\\_time ЧЧ:ММ` укажи, во сколько ты хотел бы получать рассылку\\. " +
"\"ЧЧ:ММ\" — часы и минуты \\(без кавычек\\)\\.\n\n" +
"С помощью команды `/set\\_date ДД\\.ММ\\.ГГГГ` укажи, когда у тебя заканчивается университет\\. " +
"\"ДД\\.ММ\\.ГГГГ\" — день, месяц и год \\(без кавычек\\)\\.\n\n" +
"Примеры:\n" +
"`/set\\_time 14:10` установит время ежедневной рассылки на 14:10\n" +
"`/set\\_date 25\\.06\\.2025` установит дату окончания университета на 25 июня 2025 года\n"

const SuccessTimeChangeText = 
"Хорошо\\! Каждый день в %s я буду оповещать тебя о том, сколько " + 
"дней осталось до конца университета\\!"

const SuccessDateChangeText =
"Хорошо\\! Буду считать, сколько времени осталось до %s"

const FailureTimeChangeText =
"Кажется, ты не указал время \\(или сделал это неправильно\\)\\." +
"С помощью команды `/set\\_time ЧЧ:ММ` укажи, во сколько ты хотел бы получать рассылку\\. " +
"\"ЧЧ:ММ\" — часы и минуты \\(без кавычек\\)\\."

const FailureDateChangeText =
"Кажется, ты не указал дату \\(или сделал это неправильно\\)\\." +
"С помощью команды `/set\\_date ДД\\.ММ\\.ГГГГ` укажи, когда у тебя заканчивается университет\\. " +
"\"ДД\\.ММ\\.ГГГГ\" — день, месяц и год \\(без кавычек\\)\\."