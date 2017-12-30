package slv

import (
	"fmt"

	"github.com/ryym/slv/slv/test"
	"github.com/ryym/slv/slv/tp"
)

func TestAll(app *tp.Slv) (bool, error) {
	pd := app.ProbDir
	fmt.Printf("testing %s...\n", pd.SrcFile())

	prg, err := app.Program.NewProgram(pd.SrcPath(), pd.DestDir())
	if err != nil {
		return false, err
	}

	return test.TestAll(prg, pd.TestDir())
}
