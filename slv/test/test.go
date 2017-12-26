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
)

func TestAll(execCmds []string, testdir string, printer TestResultPrinter) error {
	fs, err := ioutil.ReadDir(testdir)
	if err != nil {
		return errors.Wrap(err, "Failed to read test directory")
	}

	totalResult := totalTestResult{}

	for _, entry := range fs {
		filename := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(filename, ".toml") {
			continue
		}

		t, err := loadTestCases(testdir, filename)
		if err != nil {
			return err
		}

		totalResult.CaseCnt += len(t.Test)
		for i, inout := range t.Test {
			cmd := exec.Command(execCmds[0], execCmds[1:]...)
			out, err := runTestCase(inout.In, cmd)
			if err != nil {
				return err
			}

			if !strings.HasSuffix(inout.Out, "\n") {
				inout.Out += "\n"
			}
			if inout.Name == "" {
				inout.Name = fmt.Sprintf("%s[%d]", filename, i)
			}

			result := testResult{
				Ok:       inout.Out == out,
				TestCase: inout,
				Actual:   out,
				Filename: filename,
			}

			printer.ShowResult(&result)

			if result.Ok {
				totalResult.PassedCnt += 1
			} else {
				totalResult.Fails = append(totalResult.Fails, result)
			}
		}
	}

	printer.ShowFailures(totalResult.Fails)
	printer.ShowSummary(&totalResult)

	return nil
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
