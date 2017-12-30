package probdir

import (
	"fmt"
	"path/filepath"

	"github.com/ryym/slv/slv/tp"
)

func NewProbDir(srcPath string) (pd tp.ProbDir, err error) {
	srcPath, err = filepath.Abs(srcPath)
	if err != nil {
		return pd, nil
	}

	root, err := GetRootPath(srcPath)
	if err != nil {
		return pd, err
	}

	return &probDirImpl{
		rootDir: root,
		srcPath: srcPath,
	}, nil
}

type probDirImpl struct {
	rootDir string
	srcPath string
}

func (pd *probDirImpl) WorkDir() string {
	return filepath.Join(pd.rootDir, ".slv")
}
func (pd *probDirImpl) SrcDir() string {
	return filepath.Join(pd.rootDir, "src")
}
func (pd *probDirImpl) TestDir() string {
	return filepath.Join(pd.rootDir, "test")
}
func (pd *probDirImpl) DestDir() string {
	return fmt.Sprintf("%s/%s.built", pd.WorkDir(), pd.SrcFile())
}
func (pd *probDirImpl) SrcPath() string {
	return pd.srcPath
}
func (pd *probDirImpl) SrcFile() string {
	return filepath.Base(pd.srcPath)
}
