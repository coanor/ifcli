package ifcli

import (
	"fmt"
	"os"
	"strings"
)

func Executor(t string) {

	t = strings.TrimSpace(t)

	switch strings.ToUpper(t) {
	case `EXIT`, `Q`:
		if curConn != nil {
			curConn.Close()
		}

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

		t = strings.Join(strings.Fields(t), " ") // remove dup spaces

		// pass
		if strings.HasPrefix(strings.ToUpper(t), `USE`) {
			elems := strings.SplitN(t, " ", 2)
			if len(elems) != 2 {
				fmt.Println("[error] invalid USE statement")
				return
			}

			if curConn == nil {
				fmt.Println("[error] not connected")
				return
			}

			curConn.curDB = elems[1]
			return
		} else if strings.HasPrefix(strings.ToUpper(t), `CONN`) { // connect to another influxdb
			elems := strings.SplitN(t, " ", 2)
			if len(elems) != 2 {
				fmt.Println("[error] invalid CONN statement")
				return
			}

			c, ok := curConnections[elems[1]]
			if !ok {
				fmt.Printf("[error] CONN %s not exist", elems[1])
				return
			}

			if err := c.Connect(); err != nil {
				fmt.Printf("[error] connect to %s failed: %s", c.Key(), err.Error())
				return
			}

			return
		}
	}

	DoQuery(t)

}
