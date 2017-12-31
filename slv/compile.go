package slv

import (
	"errors"
	"fmt"

	"github.com/ryym/slv/slv/tp"
)

func Compile(app *tp.Slv) error {
	pd := app.ProbDir
	prg, err := app.Program.NewProgram(pd.SrcPath(), pd.DestDir())
	if err != nil {
		return err
	}

	ret, err := prg.Compile()
	if err != nil {
		return err
	}

	if !ret.Compiled {
		return errors.New("This does not need compilation")
	}

	if len(ret.Output) > 0 {
		fmt.Println(string(ret.Output))
	}

	return nil
}
