package slv

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/probdir"
	"github.com/ryym/slv/slv/test"
	"github.com/ryym/slv/slv/tp"
)

func New(opts *tp.CmdNewOpts) (err error) {
	name := opts.Name
	err = os.Mkdir(name, 0755)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s directory", name)
	}
	return probdir.Mkdirs(name)
}

func Compile(app *tp.Slv) error {
	ret, err := app.Program.Compile()
	if err != nil {
		return err
	}

	if !ret.Compiled {
		return errors.New("this does not need compilation")
	}

	if len(ret.Output) > 0 {
		fmt.Println(string(ret.Output))
	}

	return nil
}

func Run(app *tp.Slv) error {
	fmt.Printf("running %s...\n", app.ProbDir.SrcFile())
	return app.Program.RunWithPipes(os.Stdin, os.Stdout)
}

func Test(app *tp.Slv) (bool, error) {
	pd := app.ProbDir
	fmt.Printf("testing %s...\n", pd.SrcFile())

	return test.TestAll(app.Program, pd.TestDir())
}
