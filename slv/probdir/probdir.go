package probdir

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ryym/slv/slv/tp"
)

const (
	WORK_DIR = ".slv"
	SRC_DIR  = "src"
	TEST_DIR = "test"
)

func NewProbDir(srcPath string) (pd tp.ProbDir, err error) {
	srcPath, err = filepath.Abs(srcPath)
	if err != nil {
		return pd, nil
	}

	newPd := &probDirImpl{
		rootDir: filepath.Dir(filepath.Dir(srcPath)),
		srcPath: srcPath,
	}

	if !isValidProbDir(newPd) {
		return pd, errors.New("invalid directory structure")
	}

	return newPd, nil
}

type probDirImpl struct {
	rootDir string
	srcPath string
}

func (pd *probDirImpl) WorkDir() string {
	return filepath.Join(pd.rootDir, WORK_DIR)
}

func (pd *probDirImpl) SrcDir() string {
	return filepath.Join(pd.rootDir, SRC_DIR)
}

func (pd *probDirImpl) TestDir() string {
	return filepath.Join(pd.rootDir, TEST_DIR)
}

func (pd *probDirImpl) DestDir() string {
	return filepath.Join(pd.WorkDir(), pd.SrcFile()+".built")
}

func (pd *probDirImpl) SrcPath() string {
	return pd.srcPath
}

func (pd *probDirImpl) SrcFile() string {
	return filepath.Base(pd.srcPath)
}

func isValidProbDir(pd tp.ProbDir) bool {
	return isDir(pd.SrcDir()) && isDir(pd.TestDir())
}

func isDir(p string) bool {
	s, err := os.Stat(p)
	return err == nil && s.IsDir()
}
