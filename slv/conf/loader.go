package conf

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/ryym/slv/slv/tp"
)

func NewConfigLoader() tp.ConfigLoader {
	return &configLoaderImpl{}
}

type configLoaderImpl struct{}

func (cl *configLoaderImpl) Load(r io.Reader) (conf tp.Config, err error) {
	_, err = toml.DecodeReader(r, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
