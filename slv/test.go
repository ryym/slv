package slv

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/prgs"
	"github.com/ryym/slv/slv/probdir"
	"github.com/ryym/slv/slv/t"
	"github.com/ryym/slv/slv/test"
)

func TestAll(c *t.ExecConf) (bool, error) {
	execCmds, err := findAndCompile(c)
	if err != nil {
		return false, err
	}

	testdir := filepath.Join(c.RootDir, "test")
	printer := test.NewResultPrinter()
	return test.TestAll(execCmds, testdir, printer)
}

func findAndCompile(c *t.ExecConf) ([]string, error) {
	prg, err := prgs.FindProgram(c.SrcPath)
	if err != nil {
		return nil, err
	}

	destDir := probdir.GetDestDir(c)
	cmds := prg.GetCompileCmds(c.SrcPath, destDir)

	if cmds.Cmds != nil {
		err = compileProgram(&cmds, destDir)
		if err != nil {
			return nil, err
		}
	}

	return prg.GetExecCmds(cmds.ExecPath), nil
}

func compileProgram(cmds *t.CompileCmds, destDir string) error {
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
