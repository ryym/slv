package slv

import (
	"fmt"

	"github.com/ryym/slv/slv/test"
	"github.com/ryym/slv/slv/tp"
)

func TestAll(app *tp.Slv) (bool, error) {
	pd := app.ProbDir
	fmt.Printf("testing %s...\n", pd.SrcFile())

	return test.TestAll(app.Program, pd.TestDir())
}
