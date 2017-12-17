package slv

import (
	"github.com/ryym/slv/slv/probdir"
	"github.com/ryym/slv/slv/t"
)

func MakeExecConf(srcPath string) (conf t.ExecConf, err error) {
	pbd, err := probdir.NewFromSrcPath(srcPath)
	if err != nil {
		return conf, err
	}

	return t.ExecConf{
		Probdir: pbd,
	}, nil
}
