# TG_REACTION
A Telegram bot that sends reactions based on regex patterns defined in JSON configuration.

## Documentation

- [CONTRIBUTING.md](CONTRIBUTING.md) - developer documentation
- [REACTIONS.md](REACTIONS.md) - detailed configuration documentation

## Requirements

* The bot must have permissions to use reactions in groups
* The bot must be an administrator in the group

## Running via Docker
```bash
docker pull ghcr.io/fromsi/tg_reaction:latest

docker run --rm \
    -e APP_TELEGRAM_TOKEN="your_bot_token" \
    ghcr.io/fromsi/tg_reaction:latest
```

## Environment Variables
* APP_TELEGRAM_TOKEN - your Telegram bot token