package ifcli

import (
	client "github.com/influxdata/influxdb1-client/v2"
)

var (
	CurDB      = ``
	CurFMT     = `` // not used
	DisableNil = false
	PromptStr  string
	Prompt     = `influx-cli`

	IflxCli client.Client
)

func LivePromptPrefix() (string, bool) {
	switch CurDB {
	case ``:
		PromptStr = Prompt + "." + `<no DB>` + " > "
		return PromptStr, true
	default:
		PromptStr = Prompt + "." + CurDB + " > "
		return PromptStr, true
	}
}
