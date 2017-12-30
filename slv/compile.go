package slv

import (
	"errors"

	"github.com/ryym/slv/slv/prgs"
	"github.com/ryym/slv/slv/tp"
)

func Compile(app *tp.Slv) (string, error) {
	srcPath := app.ProbDir.SrcPath()
	prg, err := prgs.FindProgram(srcPath)
	if err != nil {
		return "", err
	}

	destDir := app.ProbDir.DestDir()
	cmds := prg.GetCompileCmds(srcPath, destDir)

	if cmds.Cmds == nil {
		return "", errors.New("This does not need compilation")
	} else {
		err = compileProgram(&cmds, destDir)
		if err != nil {
			return "", err
		}
	}

	return cmds.ExecPath, nil
}
