package ifcli

import (
	"fmt"
	"io"
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
			handleUseStmt(t)
			return
		}

		if strings.HasPrefix(strings.ToUpper(t), `CONN`) { // connect to another influxdb
			handleUseStmt(t)
			return
		}

		if strings.HasPrefix(strings.ToUpper(t), `TEE`) { // forward output to another file
			handleTeeStmt(t)
			return
		}
	}

	DoQuery(t)
}

func handleTeeStmt(t string) {
	var err error
	elems := strings.SplitN(t, " ", 2)
	if len(elems) != 2 {
		fmt.Println("[error] invalid TEE statement")
		return
	}

	if teeFile != nil {
		teeFile.Close()
	}

	teeFile, err = os.Create(elems[1])
	if err != nil {
		fmt.Printf("[error] %s", err.Error())
		return
	}

	tee = io.MultiWriter(os.Stdout, teeFile)
}

func handleUseStmt(t string) {

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
}

func handleConnStmt(t string) {
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
