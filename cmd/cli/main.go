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
	flagHost     = flag.String("host", ``, `InfluxDB host, format: http[s]://<host>:port`)
	flagUser     = flag.String(`user`, ``, `InfluxDB user name`)
	flagPassword = flag.String(`pwd`, ``, `InfluxDB password`)
	flagDB       = flag.String(`db`, ``, `set connection default DB name`)

	flagPrompt     = flag.String("prompt", "influx-cli", `set connection prompt string`)
	flagDisableNil = flag.Bool(`disable-nil`, false, `when show InfluxDB data, disable nil print, you can switch it during run time`)

	flagEncrypt = flag.String("encrypt", ``, `used to update InfluxDB password in .ifclirc`)
	flagIfCliRC = flag.String(`ifclirc`, ``, `use specified .ifclirc path instead of ~/.ifclirc`)
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	if *flagEncrypt != "" {
		en, err := ifcli.DoEncrypt(*flagEncrypt)
		if err != nil {
			fmt.Printf("failed to encrypt: %s\n", err.Error())
		} else {
			fmt.Println(en)
		}

		return
	}

	if *flagDisableNil {
		ifcli.DisableNil = true
	}

	if *flagIfCliRC != "" {
		if err := ifcli.SetIfCliRC(*flagIfCliRC); err != nil {
			log.Fatal(err)
		}
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
