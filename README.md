# TG_REACTION
A Telegram bot that sends reactions based on regex patterns defined in JSON configuration.

## Documentation

- [CONTRIBUTING.md](CONTRIBUTING.md) - developer documentation
- [REACTIONS.md](REACTIONS.md) - detailed configuration documentation

## Requirements

* The bot must have permissions to use reactions in groups
* The bot must be an administrator in the group

## Reaction Features

The bot can send several types of emoji reactions in response to different events in Telegram chats:

### 1. Text Messages
- Sends various reactions based on regex patterns defined in the configuration
- Can use different reactions during holidays and special events
- Supports multiple regex patterns and random selection from multiple emojis
- Updates reactions on edited messages: when a message is edited, the bot removes the previous reaction and applies a new one based on the updated text

### 2. Group Events
- Sends 🎉 (Party) reaction when:
  - A new member joins the group
  - Group photo is updated
  - Group title is changed
  
- Sends 😭 (Crying) reaction when:
  - A member leaves the group
  - Group photo is deleted

### 3. Message Events
- Sends 👀 (Eyes) reaction when a message is pinned

### 4. Interactive Content
- Sends 🏆 (Trophy) reaction when a dice roll results in a winning value (6)

## Running via Docker
```bash
docker pull ghcr.io/fromsi/tg_reaction:latest

docker run --rm \
    -e APP_TELEGRAM_TOKEN="your_bot_token" \
    ghcr.io/fromsi/tg_reaction:latest
```

## Environment Variables
* APP_TELEGRAM_TOKEN - your Telegram bot token