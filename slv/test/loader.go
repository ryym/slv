package test

import (
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
	tomlData, err := ioutil.ReadFile(filepath.Join(tl.testDir, filename))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s", filename)
	}

	testData := testData{}
	err = toml.Unmarshal(tomlData, &testData)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse TOML content of %s", filename)
	}

	return testData.Test, nil
}
