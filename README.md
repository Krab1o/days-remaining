# Days-remaining

Бот, который считает, сколько дней осталось до захардкоженной даты. Присылает
данные об этом раз в сутки. Чтобы указать, во сколько он будет это присылать,
можно запустить команду /set_time

# Build

1. Склонируйте репозиторий.
2. Создайте .env-файл, в котором пропишете собственные переменные окружения:
```
TELEGRAM_TOKEN=<ВАШ ТОКЕН>
# UTC +4:00
TZ=Asia/Dubai
```

При изменении временной зоны все сообщения будут присылаться по времени этой
временной зоны.

3. Запустите docker-compose up --build и тестируйте своего бота.