package slv

import (
	"fmt"
	"os"

	"github.com/ryym/slv/slv/tp"
)

func Run(app *tp.Slv) error {
	fmt.Printf("running %s...\n", app.ProbDir.SrcFile())
	return app.Program.RunWithPipes(os.Stdin, os.Stdout)
}
