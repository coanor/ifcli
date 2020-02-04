package ifcli

import (
	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	CurFMT     = `` // not used
	DisableNil = false
	PromptStr  string
	Prompt     = `influx-cli`

	IflxCli client.Client
)

func LivePromptPrefix() (string, bool) {
	if curConn == nil {
		PromptStr = Prompt + "." + `[not connected]` + " > "
		return PromptStr, true
	}

	switch curConn.curDB {
	case ``:
		PromptStr = Prompt + "." + `[no DB]` + " > "
		return PromptStr, true
	default:
		PromptStr = Prompt + "." + curConn.curDB + " > "
		return PromptStr, true
	}
}
