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

	destDir := fmt.Sprintf("%s/%s.built", c.WorkDir, c.SrcFile)
	cmds := prg.GetCompileCmds(c.SrcPath, destDir)

	if cmds.Cmds != nil {
		_, err = os.Stat(destDir)
		if os.IsNotExist(err) {
			err = os.Mkdir(destDir, 0755)
		}
		if err != nil {
			return errors.Wrapf(err, "Failed to create work dir for %s", c.SrcFile)
		}

		cmd := exec.Command(cmds.Cmds[0], cmds.Cmds[1:]...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return errors.Wrap(err, string(out))
		}
	}

	execCmds := prg.GetExecCmds(cmds.ExecPath)

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
				return errors.Wrapf(err, "Failed to pipe stdin to %s", c.SrcFile)
			}

			go func() {
				defer stdin.Close()
				io.WriteString(stdin, inout.In)
			}()

			out, err := cmd.CombinedOutput()
			if err != nil {
				return errors.Wrap(err, string(out))
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
