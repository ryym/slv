package slv

import (
	"errors"

	"github.com/ryym/slv/slv/tp"
)

func Compile(app *tp.Slv) (string, error) {
	pd := app.ProbDir
	prg, err := app.Program.NewProgram(pd.SrcPath(), pd.DestDir())
	if err != nil {
		return "", err
	}

	ret, err := prg.Compile()
	if err != nil {
		return "", err
	}
	if !ret.Compiled {
		return "", errors.New("This does not need compilation")
	}

	return ret.ExecPath, nil
}
