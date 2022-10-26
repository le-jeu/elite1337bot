# 1337 telegram bot

[![GitHub Actions - Test](https://github.com/pascalroose/elite1337bot/actions/workflows/test.yaml/badge.svg)](https://github.com/PascalRoose/elite1337bot/actions/workflows/test.yaml)
[![GitHub Actions - Publish](https://github.com/pascalroose/elite1337bot/actions/workflows/publish.yaml/badge.svg)](https://github.com/PascalRoose/elite1337bot/actions/workflows/publish.yaml)
[![Telegram bot](https://img.shields.io/badge/Telegram-@elite1337bot-blue.svg)](https://t.me/elite1337bot)

Telegram bot for providing a 1337 (leet) scoreboard for chats. 

# Installation

You can install the bot by simply the bot using Docker. There are two environment variables that are used:

- `TGBOT_TOKEN` - Get the bot token from [@Botfather](https://t.me/Botfather)
- `TZ` - Name of the [timezone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) the bot will use for server time
  -  If not provided the default timezone is `Europe/Brussels`

```bash
docker run --name 1337bot -e TGBOT_TOKEN=insert-token-here -d ghcr.io/pascalroose/elite1337bot
```
