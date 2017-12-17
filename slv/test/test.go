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

type failedTestCase struct {
	testCase
	Actual string
	Name   string
}

type testResult struct {
	Ok          bool
	CaseCnt     int
	PassedCnt   int
	FailedCases []failedTestCase
}

func TestAll(c *t.ExecConf) (result testResult, err error) {
	prg, err := prgs.FindProgram(c.SrcPath)
	if err != nil {
		return result, err
	}

	destDir := fmt.Sprintf("%s/%s.built", c.WorkDir, c.SrcFile)
	cmds := prg.GetCompileCmds(c.SrcPath, destDir)
	err = compileIfNeed(&cmds, destDir)
	if err != nil {
		return result, err
	}

	execCmds := prg.GetExecCmds(cmds.ExecPath)

	testdir := filepath.Join(c.RootDir, "test")
	fs, err := ioutil.ReadDir(testdir)
	if err != nil {
		return result, errors.Wrap(err, "Failed to read test directory")
	}

	result.FailedCases = []failedTestCase{}
	for _, entry := range fs {
		filename := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(filename, ".toml") {
			continue
		}

		t, err := loadTestCases(testdir, filename)
		if err != nil {
			return result, err
		}

		result.CaseCnt += len(t.Test)
		for i, inout := range t.Test {
			cmd := makeCommand(execCmds)
			out, err := runTestCase(inout.In, cmd)
			if err != nil {
				return result, err
			}

			expected := inout.Out
			if !strings.HasSuffix(expected, "\n") {
				expected += "\n"
			}

			if out == expected {
				result.PassedCnt += 1
			} else {
				result.FailedCases = append(result.FailedCases, failedTestCase{
					testCase: inout,
					Actual:   out,
					Name:     fmt.Sprintf("%s[%d]", filename, i),
				})
			}
		}
	}
	result.Ok = len(result.FailedCases) == 0

	return result, nil
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

func runTestCase(input string, cmd *exec.Cmd) (actual string, err error) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", errors.Wrap(err, "Failed to pipe stdin")
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, string(out))
	}

	return string(out), nil
}
