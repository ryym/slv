package slv

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/prgs"
	"github.com/ryym/slv/slv/probdir"
	"github.com/ryym/slv/slv/tp"
)

func NewSlvApp(pathOrLang string, baseDir string) (slv tp.Slv, err error) {
	srcPath, err := findSrc(pathOrLang, baseDir)
	if err != nil {
		return slv, err
	}

	pd, err := probdir.NewProbDir(srcPath)
	if err != nil {
		return slv, err
	}

	return tp.Slv{
		ProbDir: pd,
	}, nil
}

func findSrc(pathOrLang string, baseDir string) (string, error) {
	if _, err := os.Stat(pathOrLang); err == nil {
		return pathOrLang, nil
	}

	srcDir := probdir.GetSrcDir(baseDir)
	if _, err := os.Stat(srcDir); err != nil {
		return "", errors.New("Could not find src")
	}

	return findFirstSrcByLang(srcDir, pathOrLang)
}

func findFirstSrcByLang(srcDir string, lang string) (string, error) {
	ext := prgs.FindExtByLang(lang)
	if ext == "" {
		return "", fmt.Errorf("Unknown language: %s", lang)
	}

	fs, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read src dir")
	}

	for _, f := range fs {
		if filepath.Ext(f.Name()) == ext {
			return filepath.Join(srcDir, f.Name()), nil
		}
	}
	return "", fmt.Errorf("Could not find src")
}
