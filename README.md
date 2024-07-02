# TG_REACTION
Sending reaction by regex via Telegram.

* У бота должны быть права на использование групп
* Бот в группе должен быть администратором

Запуск через докер
```bash
docker run --rm \
    -e TG_REACTION_TOKEN="secret" \
    -e TG_REACTION_REGEX="(?i)(вечер|утро)" \
    ghcr.io/fromsi/tg_reaction:latest
```

## ENV
* TG_EMOJI_TOKEN - токен бота
* TG_EMOJI_REGEX - регулярное выражение для триггера
