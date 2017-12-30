package slv

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/prgs"
	"github.com/ryym/slv/slv/tp"
)

func findAndCompile(srcPath string, destDir string) ([]string, error) {
	prg, err := prgs.FindProgram(srcPath)
	if err != nil {
		return nil, err
	}

	cmds := prg.GetCompileCmds(srcPath, destDir)

	if cmds.Cmds != nil {
		err = compileProgram(&cmds, destDir)
		if err != nil {
			return nil, err
		}
	}

	return prg.GetExecCmds(cmds.ExecPath), nil
}

func compileProgram(cmds *tp.CompileCmds, destDir string) error {
	_, err := os.Stat(destDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(destDir, 0755)
	}
	if err != nil {
		return errors.Wrap(err, "Failed to create work dir")
	}

	out, err := exec.Command(cmds.Cmds[0], cmds.Cmds[1:]...).CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(out))
	}

	return nil
}
