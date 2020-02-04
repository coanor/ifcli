package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"

	"github.com/coanor/ifcli"
)

var (
	flagHost       = flag.String("host", ``, ``)
	flagUser       = flag.String(`user`, ``, ``)
	flagPassword   = flag.String(`pwd`, ``, ``)
	flagDB         = flag.String(`db`, ``, ``)
	flagPrompt     = flag.String("prompt", "influx-cli", ``)
	flagDisableNil = flag.Bool(`disable-nil`, false, ``)
)

func main() {
	flag.Parse()

	if *flagDisableNil {
		ifcli.DisableNil = true
	}

	ifcli.LoadHist()

	if *flagHost != "" {
		c := &ifcli.Conn{
			Host:      *flagHost,
			User:      *flagUser,
			Password:  *flagPassword,
			DefaultDB: *flagDB,
			Created:   time.Now(),
			Prompt:    *flagPrompt,
		}

		if err := c.Connect(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Please use `EXIT|exit|Q|q` to exit this program.")
	defer fmt.Println("Bye!")

	p := prompt.New(
		ifcli.Executor,
		ifcli.SugCompleter,
		prompt.OptionTitle("ifcli: interactive InfluxDB client"),
		prompt.OptionPrefix(ifcli.PromptStr),
		prompt.OptionLivePrefix(ifcli.LivePromptPrefix),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)

	p.Run()
}
