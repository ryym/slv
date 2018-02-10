package test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type testLoaderImpl struct {
	testDir string
}

func newTestLoader(testDir string) testLoader {
	return &testLoaderImpl{testDir}
}

func (tl *testLoaderImpl) ListFileNames() ([]string, error) {
	fs, err := ioutil.ReadDir(tl.testDir)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, f := range fs {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".toml") {
			names = append(names, f.Name())
		}
	}
	return names, nil
}

type testData struct {
	Test []testCase
}

func (tl *testLoaderImpl) Load(filename string) ([]testCase, error) {
	tomlData, err := ioutil.ReadFile(filepath.Join(tl.testDir, filename))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s", filename)
	}

	testData := testData{}
	err = toml.Unmarshal(tomlData, &testData)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse TOML content of %s", filename)
	}

	// Load input and output from files if neceessary.
	for i, _ := range testData.Test {
		tc := &testData.Test[i]
		if tc.InFile != "" {
			input, err := loadTest(tl.testDir, tc.InFile)
			if err != nil {
				return nil, errors.Wrapf(err, "test case [%d]", i)
			}
			tc.In = input
		}
		if tc.OutFile != "" {
			output, err := loadTest(tl.testDir, tc.OutFile)
			if err != nil {
				return nil, errors.Wrapf(err, "test case [%d]", i)
			}
			tc.Out = output
		}

		if tc.In == "" && tc.Out == "" {
			return nil, fmt.Errorf(
				"test case [%d]: `in` or `out` must be specified", i,
			)
		}
	}

	return testData.Test, nil
}

func loadTest(dir string, path string) (string, error) {
	if !filepath.IsAbs(path) {
		path = filepath.Join(dir, path)
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read %s", path)
	}
	return string(content), nil
}
