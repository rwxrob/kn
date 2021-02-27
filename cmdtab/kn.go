package cmdtab

import (
	"github.com/rwxrob/cmdtab"
	_ "github.com/rwxrob/cmdtab-config"
	_ "github.com/rwxrob/cmdtab-pomo"
)

func Execute(s string) {
	//	cmdtab.Dump(cmdtab.Index)
	cmdtab.Execute(s)
}

func init() {
	x := cmdtab.New("kn", "day", "cd", "config", "add", "tstamp", "html", "epoch", "pomo")
	x.Summary = ``
	x.Version = "1.0.0"
	x.Author = "Rob Muhlestein <rob@rwx.gg> (rob.rwx.gg)"
	x.Git = "github.com/afkworks/kn/cmd/kn/kn"
	x.Copyright = "(c) Rob Muhlestein"
	x.License = "MPL-2"
	x.Description = ``
}
