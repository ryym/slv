package slv

import (
	"errors"

	"github.com/ryym/slv/slv/prgs"
	"github.com/ryym/slv/slv/probdir"
	"github.com/ryym/slv/slv/tp"
)

func Compile(c *tp.ExecConf) (string, error) {
	prg, err := prgs.FindProgram(c.SrcPath)
	if err != nil {
		return "", err
	}

	destDir := probdir.GetDestDir(c)
	cmds := prg.GetCompileCmds(c.SrcPath, destDir)

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
