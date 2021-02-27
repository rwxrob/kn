package cmdtab

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("cd")
	x.Summary = `changes current working directory to current root node`
	x.Description = ``
	x.Method = func(args []string) error {
		shell, err := exec.LookPath(os.Getenv("SHELL"))
		if err != nil {
			return fmt.Errorf("SHELL not found: %v", err)
		}
		os.Chdir("/tmp")
		return syscall.Exec(shell, []string{shell}, os.Environ())

		/*
			var cmd *exec.Cmd
				cd, err := exec.LookPath("cd")
				cmd = exec.Command(cd, "/tmp")
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				return cmd.Run()
		*/

		// TODO exec `cd`  command so that it affects current shell (not
		// subshell)
	}
}
