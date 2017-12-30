package slv

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ryym/slv/slv/tp"
)

func Run(app *tp.Slv) error {
	pb := app.ProbDir
	fmt.Printf("running %s...\n", pb.SrcFile())

	execCmds, err := findAndCompile(pb.SrcPath(), pb.DestDir())
	if err != nil {
		return err
	}

	cmd := exec.Command(execCmds[0], execCmds[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
