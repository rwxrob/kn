package cmdtab

import (
	"fmt"
	"time"

	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("tstamp")
	x.Summary = `prints an RFC3338 (IS08601) stamp`

	x.Description = `Prints an RFC3338 (ISO8601) timestamp suitable for
		importing into any modern database and as used by JSON extensions
		and YAML structured data files.`

	x.Method = func(args []string) error {
		fmt.Println(time.Now().Format(time.RFC3339))
		return nil
	}
}
