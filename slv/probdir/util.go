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
		fileutil.IsDir(filepath.Join(dir, ".slv")) &&
		fileutil.IsDir(filepath.Join(dir, "src")) &&
		fileutil.IsDir(filepath.Join(dir, "test"))
}

func GetWorkDir(rootPath string) string {
	return filepath.Join(rootPath, ".slv")
}
