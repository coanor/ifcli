package ifcli

import (
	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	CurFMT     = `` // not used
	DisableNil = false
	PromptStr  string

	IflxCli client.Client
)

func LivePromptPrefix() (string, bool) {
	if curConn == nil {
		PromptStr = "influx-cli." + `[not connected]` + " > "
		return PromptStr, true
	}

	switch curConn.curDB {
	case ``:
		PromptStr = curConn.Prompt + "." + `[no DB]` + " > "
		return PromptStr, true
	default:
		PromptStr = curConn.Prompt + "." + curConn.curDB + " > "
		return PromptStr, true
	}
}
