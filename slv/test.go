package slv

import (
	"fmt"

	"github.com/ryym/slv/slv/test"
	"github.com/ryym/slv/slv/tp"
)

func TestAll(app *tp.Slv) (bool, error) {
	pd := app.ProbDir
	fmt.Printf("testing %s...\n", pd.SrcFile())

	execCmds, err := findAndCompile(pd.SrcPath(), pd.DestDir())
	if err != nil {
		return false, err
	}

	printer := test.NewResultPrinter()
	return test.TestAll(execCmds, pd.TestDir(), printer)
}
