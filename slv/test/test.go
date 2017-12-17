package test

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
	err = compileIfNeed(&cmds, destDir)
	if err != nil {
		return err
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

		t, err := loadTestCases(testdir, entry.Name())
		if err != nil {
			return err
		}

		for _, inout := range t.Test {
			cmd := makeCommand(execCmds)
			err = runTestCase(c, &inout, cmd)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func compileIfNeed(cmds *t.CompileCmds, destDir string) error {
	if cmds.Cmds != nil {
		_, err := os.Stat(destDir)
		if os.IsNotExist(err) {
			err = os.Mkdir(destDir, 0755)
		}
		if err != nil {
			return errors.Wrap(err, "Failed to create work dir")
		}

		out, err := makeCommand(cmds.Cmds).CombinedOutput()
		if err != nil {
			return errors.Wrap(err, string(out))
		}
	}
	return nil
}

func makeCommand(cmds []string) *exec.Cmd {
	return exec.Command(cmds[0], cmds[1:]...)
}

func loadTestCases(dir string, filename string) (tcs testCases, err error) {
	tomlFile, err := os.Open(filepath.Join(dir, filename))
	if err != nil {
		return tcs, errors.Wrapf(err, "Failed to open %s", filename)
	}

	tomlData, err := ioutil.ReadAll(tomlFile)
	if err != nil {
		return tcs, errors.Wrapf(err, "Failed to read %s", filename)
	}

	err = toml.Unmarshal(tomlData, &tcs)
	if err != nil {
		return tcs, errors.Wrapf(err, "Failed to parse TOML content of %s", filename)
	}

	return tcs, nil
}

func runTestCase(c *t.ExecConf, inout *testCase, cmd *exec.Cmd) error {
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

	return nil
}
