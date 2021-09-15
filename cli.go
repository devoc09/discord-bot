package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/devoc09/discord-bot/internal"
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
	var info, send, help, version bool

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = func() {
		fmt.Fprintf(cli.errStream, usage, Name)
	}
	flags.BoolVar(&info, "i", false, "send Server Info(CPU, Memory) to discord webhook url")
	flags.BoolVar(&send, "s", false, "send Message to discord webhook url")
	flags.BoolVar(&help, "h", false, "display help message")
	flags.BoolVar(&version, "v", false, "display the version")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if info {
		arg := flags.Arg(0)
		homedir := os.Getenv("HOME")
		username, avatar_url, webhook_url, err := webhook.ReadConfig(homedir + "/.config/discord-bot/config.json")
		if err != nil {
			return ExitCodeReadConfigError
		}
		sys := &internal.System{}
		si := sys
		comcutil := si.CpuUtil(0, true) // combine cpu util
		var sum float64
		for _, v := range comcutil {
			sum += v
		}
		mutil := si.MemUtil() // memory util

		feilds := make([]webhook.Field, 2, 2)

		// feilds[0] = webhook.Field{Name: "CPU INFO", Value: fmt.Sprintf("Per: %v%%", strconv.FormatFloat(comcutil[0], 'f', 2, 64))}
		feilds[0] = webhook.Field{Name: "CPU INFO", Value: fmt.Sprintf("Per: %v%%", strconv.FormatFloat(sum/8, 'f', 2, 64))}
		feilds[1] = webhook.Field{Name: "Memory INFO", Value: fmt.Sprintf("Total: %v, Free: %v, UserdPercent: %f%%\n", mutil.Total, mutil.Free, mutil.UsedPercent)}

		embeds := make([]webhook.Embed, 1, 1)
		embeds[0] = webhook.Embed{Color: 5620992, Fields: feilds}

		message := &webhook.EmbedMessage{Username: username, Content: arg, Avatar_url: avatar_url, Embeds: embeds}
		doneCh, postErrCh, err := webhook.SendEmbedMessage(webhook_url, message)
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

	if send {
		arg := flags.Arg(0)
		homedir := os.Getenv("HOME")
		username, avatar_url, webhook_url, err := webhook.ReadConfig(homedir + "/.config/discord-bot/config.json")
		if err != nil {
			return ExitCodeReadConfigError
		}
		message := &webhook.MinMessage{Username: username, Content: arg, Avatar_url: avatar_url}
		doneCh, postErrCh, err := webhook.SendMinMessage(webhook_url, message)
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
