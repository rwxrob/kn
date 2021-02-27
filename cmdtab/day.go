package cmdtab

import (
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("day")
	x.Summary = `outputs the current date with optional delimiter`
	x.Description = ``
	x.Method = func(args []string) error {
		// TODO add modifer
		/*
			var cmd *exec.Cmd
			if len(os.Args) > 0 {
				a := append([]string{"+%4Y/%m/%d", "-d"}, os.Args...)
				cmd = exec.Command("date", a...)
			} else {
				cmd = exec.Command("date", "+%4Y/%m/%d")
			}
			return cmd.Exec()
		*/
		return nil
	}
}
