package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"

	"github.com/coanor/ifcli"
)

var (
	flagHost       = flag.String("host", ``, ``)
	flagUser       = flag.String(`user`, ``, ``)
	flagPassword   = flag.String(`pwd`, ``, ``)
	flagDB         = flag.String(`db`, ``, ``)
	flagPrompt     = flag.String("prompt", "influx-cli ", ``)
	flagDisableNil = flag.Bool(`disable-nil`, false, ``)
)

func main() {
	flag.Parse()

	if *flagDisableNil {
		ifcli.DisableNil = true
	}

	var err error

	ifcli.IflxCli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:      *flagHost,
		Username:  *flagUser,
		Password:  *flagPassword,
		UserAgent: "ifcli",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Please use `EXIT|exit|Q|q` to exit this program.")
	defer fmt.Println("Bye!")

	if *flagDB != `` {
		ifcli.CurDB = *flagDB
	}

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
