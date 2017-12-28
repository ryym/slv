package probdir

import (
	"fmt"
	"path/filepath"

	"github.com/ryym/slv/slv/fileutil"
)

// srcPath must be absolute.
func GetRootPath(srcPath string) (string, error) {
	rootDir := filepath.Dir(filepath.Dir(srcPath))
	if IsProbDir(rootDir) {
		return rootDir, nil
	}
	return "", fmt.Errorf("Could not found problem directory of %s", filepath.Base(srcPath))
}

func IsProbDir(dir string) bool {
	return fileutil.IsDir(dir) &&
		fileutil.IsDir(GetWorkDir(dir)) &&
		fileutil.IsDir(GetSrcDir(dir)) &&
		fileutil.IsDir(GetTestDir(dir))
}

func GetWorkDir(rootPath string) string {
	return filepath.Join(rootPath, ".slv")
}

func GetSrcDir(rootPath string) string {
	return filepath.Join(rootPath, "src")
}

func GetTestDir(rootPath string) string {
	return filepath.Join(rootPath, "test")
}
