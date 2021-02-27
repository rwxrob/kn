package cmdtab

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("epoch")
	x.Summary = `prints seconds since Jan 1, 1970`
	x.Description = ``
	x.Method = func(args []string) error {
		// TODO add modifer support
		fmt.Println(time.Now().Unix())
		return nil
	}
}
