package slv

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ryym/slv/slv/tp"
)

func Run(c *tp.ExecConf) error {
	fmt.Printf("running %s...\n", c.SrcFile)

	execCmds, err := findAndCompile(c)
	if err != nil {
		return err
	}

	cmd := exec.Command(execCmds[0], execCmds[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
