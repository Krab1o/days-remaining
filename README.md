# Days-remaining

It's a telegram bot which counts how many days remaining to a particular date.
It sends this information every day in particular time. Time and date can be
specified by user.

# Build

1. Clone repository.
2. Create an .env-file, where you'll specify your own environment variables:
```
TELEGRAM_TOKEN=<YOUR TOKEN>
# UTC +4:00
TZ=Asia/Dubai
```

After timezone change all messages are going to be sent according to specified timezone

3. Run `docker-compose up --build` and test your tg-bot.