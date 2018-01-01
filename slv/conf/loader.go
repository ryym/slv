package conf

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/tp"
)

const CONF_FILE_NAME = ".slv.toml"

func NewConfigLoader() tp.ConfigLoader {
	return &configLoaderImpl{}
}

type configLoaderImpl struct{}

func (cl *configLoaderImpl) Load(r io.Reader) (*tp.Config, error) {
	var conf tp.Config
	_, err := toml.DecodeReader(r, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

// Traverse parent directories from the given path.
// `baseDir` must be an absolute path.
func (cl *configLoaderImpl) LoadFromFiles(baseDir string) ([]*tp.Config, error) {
	f, err := os.Stat(baseDir)
	if os.IsNotExist(err) {
		return nil, err
	}
	if !f.IsDir() {
		return nil, fmt.Errorf("%s is not directory", baseDir)
	}

	var confs []*tp.Config
	d := baseDir
	for true {
		confPath := filepath.Join(d, CONF_FILE_NAME)
		stat, err := os.Stat(confPath)

		if err == nil && !stat.IsDir() {
			file, err := os.Open(confPath)
			if err != nil {
				return nil, errors.Wrapf(err, "Failed to open %s", confPath)
			}

			conf, err := cl.Load(file)
			if err != nil {
				return nil, err
			}
			confs = append(confs, conf)
		} else if err != nil && !os.IsNotExist(err) {
			return nil, errors.Wrapf(err, "Failed to get stat of %s", confPath)
		}

		if isRootDir(d) {
			break
		}
		d = filepath.Dir(d)
	}

	return confs, nil
}

func (cl *configLoaderImpl) Merge(a *tp.Config, b *tp.Config) *tp.Config {
	conf := *a
	for name, lang := range b.Langs {
		conf.Langs[name] = lang
	}
	return &conf
}

func isRootDir(dir string) bool {
	if runtime.GOOS == "windows" {
		return strings.HasSuffix(dir, ":"+string(os.PathSeparator))
	} else {
		return dir == "/"
	}
}
