package slv

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/conf"
	"github.com/ryym/slv/slv/prgs"
	"github.com/ryym/slv/slv/probdir"
	"github.com/ryym/slv/slv/tp"
)

func NewSlvApp(pathOrLang string, baseDir string) (slv tp.Slv, err error) {
	confL := conf.NewConfigLoader()

	conf, err := confL.Load(strings.NewReader(DEFAULT_CONF))
	if err != nil {
		return slv, err
	}

	pdict, err := prgs.MakeProgramDict(&conf)
	if err != nil {
		return slv, err
	}

	var ext, srcPath string

	// Identify the source file path and its extension.
	if ext = filepath.Ext(pathOrLang); ext != "" {
		srcPath = pathOrLang
		if _, err := os.Stat(srcPath); err != nil {
			return slv, err
		}
	} else {
		exts, found := pdict.FindExts(pathOrLang)
		if !found {
			return slv, fmt.Errorf("Unknown language: %s", pathOrLang)
		}
		if len(exts) == 0 {
			return slv, fmt.Errorf("No extentions are defined for %s", pathOrLang)
		}

		srcDir := filepath.Join(baseDir, probdir.SRC_DIR)
		srcPath, ext, err = findFirstSrc(exts, srcDir)
		if err != nil {
			return slv, err
		}
	}

	probDir, err := probdir.NewProbDir(srcPath)
	if err != nil {
		return slv, err
	}

	def := pdict.FindDefByExt(ext)
	return tp.Slv{
		ProbDir: probDir,
		Program: prgs.NewProgram(def, srcPath, probDir.DestDir()),
	}, nil
}

func findFirstSrc(exts []string, srcDir string) (src string, ext string, err error) {
	fs, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return "", "", errors.Wrap(err, "Failed to read src directory")
	}

	for _, f := range fs {
		ext = filepath.Ext(f.Name())
		for _, e := range exts {
			if e == ext {
				return filepath.Join(srcDir, f.Name()), ext, nil
			}
		}
	}

	return "", "", fmt.Errorf("Could not find src")
}
