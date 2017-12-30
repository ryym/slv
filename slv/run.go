package slv

import (
	"fmt"
	"os"

	"github.com/ryym/slv/slv/tp"
)

func Run(app *tp.Slv) error {
	pb := app.ProbDir
	fmt.Printf("running %s...\n", pb.SrcFile())

	prg, err := app.Program.NewProgram(pb.SrcPath(), pb.DestDir())
	if err != nil {
		return err
	}

	return prg.RunWithPipes(os.Stdin, os.Stdout)
}
