package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/c-bata/go-prompt"
	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	flagHost       = flag.String("host", ``, ``)
	flagUser       = flag.String(`user`, ``, ``)
	flagPassword   = flag.String(`pwd`, ``, ``)
	flagDB         = flag.String(`db`, ``, ``)
	flagPrompt     = flag.String("prompt", "influx-cli ", ``)
	flagDisableNil = flag.Bool(`disable-nil`, false, ``)
)

var (
	curDB      = ``
	curFMT     = `` // not used
	disableNil = false
)

func main() {
	flag.Parse()

	if *flagDisableNil {
		disableNil = true
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:      *flagHost,
		Username:  *flagUser,
		Password:  *flagPassword,
		UserAgent: "ifcli",
	})

	if err != nil {
		log.Fatal(err)
	}

	if *flagDB != `` {
		curDB = *flagDB
	}

	promptStr := *flagPrompt + "." + curDB + " > "

	for {
	goon:
		t := prompt.Input(promptStr, completer, prompt.OptionSwitchKeyBindMode(prompt.EmacsKeyBind))

		t = strings.TrimSpace(t)

		switch strings.ToUpper(t) {
		case `EXIT`, `Q`:
			return

		case ``: // ignore empty line
			goto goon

		case `ENABLE_NIL`:
			disableNil = false
			goto goon

		case `DISABLE_NIL`:
			disableNil = true
			goto goon

		default:
			// pass
			if strings.HasPrefix(strings.ToUpper(t), `USE`) {
				t = strings.Join(strings.Fields(t), " ") // remove dup spaces
				elems := strings.Split(t, " ")
				if len(elems) != 2 {
					log.Printf("[error] invalid USE statement")
					goto goon
				}

				curDB = elems[1]
				promptStr = *flagPrompt + "." + curDB + " > "
				goto goon
			}
		}

		q := client.NewQuery(t, curDB, ``)
		if resp, err := c.Query(q); err == nil && resp.Error() == nil {
			showResp(resp)
		} else {
			if err == nil {
				fmt.Printf("[error] resp Err: %s\n", resp.Error())
			} else {
				fmt.Printf("[error] %s, resp Err: %s\n", err.Error(), resp.Error())
			}
		}
	}
}
