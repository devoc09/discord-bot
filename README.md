# Discord-bot
A command line tool for send message to Discord webhook URL.

## Installation
- Build from this Sources.

if you have a Go env
- Use Makefile or self build

Author's env
`go version go1.17 darwin/amd64`

## Set up
`config.json` is required<br>
1. make dir `$HOME/.config/discord-bot` and put `config.json` to `~/.config/discord-bot/`
1. make `config.yaml` into `~/.config/gtodo/`<br>
```yaml
# config.json
[
    {
        "avatar_url": "BOT's icons image url",
        "username": "BOT's user name(anything)",
        "webhook_url": "your discord webhook url"
    }
]

```

## How to use
```
Usage: discord-bot [options]
    Post message to Discord webhook URL.

Command:
    -i <message>    send <message> and host CPU & Memory Info to discord webhook url.
    -s <message>    send <message> to discord webhook url.
    -h              Print Help message
    -v              Print the version of this application
```

I have installed the command on my server and run it periodically by cron. By doing so, I can check the survival of the server.
