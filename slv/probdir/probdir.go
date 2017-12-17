package probdir

import (
	"path/filepath"

	"github.com/ryym/slv/slv/t"
)

func NewFromSrcPath(srcPath string) (pbd t.Probdir, err error) {
	srcPath, err = filepath.Abs(srcPath)
	if err != nil {
		return pbd, nil
	}

	root, err := GetRootPath(srcPath)
	if err != nil {
		return pbd, err
	}

	return t.Probdir{
		RootDir: root,
		SrcFile: filepath.Base(srcPath),
		SrcPath: srcPath,
		WorkDir: GetWorkDir(root),
	}, nil
}