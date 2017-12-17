package slv

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/prgs"
	"github.com/ryym/slv/slv/t"
)

type testCase struct {
	In  string
	Out string
}

type testCases struct {
	Test []testCase
}

func TestAll(c *t.ExecConf) error {
	prg, err := prgs.FindProgram(c.SrcPath)
	if err != nil {
		return err
	}

	execPath := fmt.Sprintf("%s/%s.compiled", c.WorkDir, c.SrcFile)
	if cmds := prg.GetCompileCmds(c.SrcPath, execPath); cmds != nil {
		cmds := prg.GetCompileCmds(c.SrcPath, execPath)
		cmd := exec.Command(cmds[0], cmds[1:]...)
		_, err = cmd.Output()
		if err != nil {
			return errors.Wrapf(err, "Failed to compile %s", c.SrcFile)
		}
	} else {
		execPath = c.SrcPath
	}

	execCmds := prg.GetExecCmds(execPath)

	testdir := filepath.Join(c.RootDir, "test")
	fs, err := ioutil.ReadDir(testdir)
	if err != nil {
		return errors.Wrap(err, "Failed to read test directory")
	}
	for _, entry := range fs {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".toml") {
			continue
		}

		tomlFile, err := os.Open(filepath.Join(testdir, entry.Name()))
		if err != nil {
			return errors.Wrapf(err, "Failed to open %s", entry.Name())
		}

		tomlData, err := ioutil.ReadAll(tomlFile)
		if err != nil {
			return errors.Wrapf(err, "Failed to read %s", entry.Name())
		}

		var t testCases
		err = toml.Unmarshal(tomlData, &t)
		if err != nil {
			return errors.Wrapf(err, "Failed to parse TOML content of %s", entry.Name())
		}

		for _, inout := range t.Test {
			cmd := exec.Command(execCmds[0], execCmds[1:]...)

			stdin, err := cmd.StdinPipe()
			if err != nil {
				return errors.Wrapf(err, "Failed to run %s", c.SrcFile)
			}

			go func() {
				defer stdin.Close()
				io.WriteString(stdin, inout.In)
			}()

			out, err := cmd.CombinedOutput()
			if err != nil {
				return errors.Wrapf(err, "Failed to run %s", c.SrcFile)
			}

			expected := inout.Out
			if !strings.HasSuffix(expected, "\n") {
				expected += "\n"
			}

			// TODO: Compare results instead of just printing them.
			fmt.Printf("IN: %s, GOT: %sWANT: %s %v\n", inout.In, out, expected, string(out) == expected)
		}
	}

	return nil
}
