package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/devoc09/discord-bot/webhook"
)

const (
	ExitCodeOK               = 0
	ExitCodeError            = 1
	ExitCodeParseFlagsError  = 2
	ExitCodeBadArgs          = 3
	ExitCodeReadConfigError  = 4
	ExitCodeSendMessageError = 5
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var send, help, version bool

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = func() {
		fmt.Fprintf(cli.errStream, usage, Name)
	}
	flags.BoolVar(&send, "s", false, "send message to discord webhook rul")
	flags.BoolVar(&help, "h", false, "display help message")
	flags.BoolVar(&version, "v", false, "display the version")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if send {
		arg := flags.Arg(0)
		// Proxmox's config "/root/.config/discord-bot/config.json"
		username, avatar_url, webhook_url, err := webhook.ReadConfig("/root/.config/discord-bot/config.json")
		if err != nil {
			return ExitCodeReadConfigError
		}
		message := &webhook.Message{Username: username, Content: arg, Avatar_url: avatar_url}
		doneCh, postErrCh, err := webhook.SendMessage(webhook_url, message)
		if err != nil {
			return ExitCodeSendMessageError
		}
		select {
		case err := <-postErrCh:
			fmt.Fprintf(cli.errStream, "error POST webhook url %s\n", err)
			return ExitCodeSendMessageError
		case msg := <-doneCh:
			fmt.Printf("Succenss sent message!! ExitCode is %d\n", msg)
		}

		return ExitCodeOK
	}

	if help {
		flags.Usage()
		return ExitCodeOK
	}

	if version {
		fmt.Fprintf(cli.errStream, "%s v%s\n", Name, Version)
		return ExitCodeOK
	}

	parsedArgs := flags.Args()

	if len(parsedArgs) != 2 {
		fmt.Fprintf(cli.errStream, "cli: must specify two arguments\n")
		flags.Usage()
		return ExitCodeBadArgs
	}

	return ExitCodeOK
}

const usage = `
Usage: %s [options]
    Post message to Discord webhook URL.

Command:
    -s <message>    send <message> to discord webhook url.
    -h              Print Help message
    -v              Print the version of this application
`
