package fileutil

import "os"

func IsDir(p string) bool {
	s, err := os.Stat(p)
	return err == nil && s.IsDir()
}
