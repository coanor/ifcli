package ifcli

import (
	"fmt"
	"os"
	"strings"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

func Executor(t string) {

	t = strings.TrimSpace(t)

	switch strings.ToUpper(t) {
	case `EXIT`, `Q`:
		fmt.Println("Bye!")
		os.Exit(0)

	case ``: // ignore empty line
		return

	case `ENABLE_NIL`:
		DisableNil = false
		return

	case `DISABLE_NIL`:
		DisableNil = true
		return

	case `RESET_SUG`:
		ResetSug()
		return

	default:
		// pass
		if strings.HasPrefix(strings.ToUpper(t), `USE`) {
			t = strings.Join(strings.Fields(t), " ") // remove dup spaces
			elems := strings.Split(t, " ")
			if len(elems) != 2 {
				fmt.Println("[error] invalid USE statement")
				return
			}

			CurDB = elems[1]
			PromptStr = Prompt + "." + CurDB + " > "
			return
		}
	}

	q := client.NewQuery(t, CurDB, ``)
	start := time.Now()

	if resp, err := IflxCli.Query(q); err == nil && resp.Error() == nil {
		n := ShowResp(resp)
		fmt.Printf("\n%d rows in set\n", n)
		fmt.Printf("time: %v\n", time.Since(start))
	} else {
		if err == nil {
			fmt.Printf("[error] resp Err: %s\n", resp.Error())
		} else {
			fmt.Printf("[error] %s, resp Err: %s\n", err.Error(), resp.Error())
		}
	}
}
