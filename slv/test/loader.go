package test

import (
	"io/ioutil"
	"os"
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

	names := make([]string, len(fs))
	for i, f := range fs {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".toml") {
			names[i] = f.Name()
		}
	}
	return names, nil
}

type testData struct {
	Test []testCase
}

func (tl *testLoaderImpl) Load(filename string) ([]testCase, error) {
	tomlFile, err := os.Open(filepath.Join(tl.testDir, filename))
	if err != nil {
		return nil, err
	}

	tomlData, err := ioutil.ReadAll(tomlFile)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to read %s", filename)
	}

	testData := testData{}
	err = toml.Unmarshal(tomlData, &testData)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to parse TOML content of %s", filename)
	}

	return testData.Test, nil
}
