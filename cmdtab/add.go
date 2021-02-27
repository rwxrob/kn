package cmdtab

import (
	"os"

	"github.com/rwxrob/cmdtab"
	"github.com/rwxrob/conf-go"
)

func init() {
	x := cmdtab.New("add")
	x.Summary = `add a new root local knowledge node`
	x.Usage = `[<dir>]`

	x.Description = `
		When called with no argument assumes the current directory is to be
		added as a root knowledge node. When called with one argument,
		assumes a path to a directory, determines the absolute path and
		adds.`

	x.Method = func(args []string) error {
		var dir string
		config := conf.New()
		err := config.Load()
		if err != nil {
			return err
		}
		switch len(args) {
		case 0:
			dir, err = os.Getwd()
			if err != nil {
				return err
			}
			err = config.SetSave("current", dir)
			return err
		case 1:
		}
		return nil
	}
}
