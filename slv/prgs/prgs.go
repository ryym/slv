package prgs

import (
	"fmt"
	"path/filepath"

	"github.com/ryym/slv/slv/t"
)

func FindProgram(srcPath string) (t.Program, error) {
	switch filepath.Ext(srcPath) {
	case ".go":
		return &ProgramGo{}, nil
	default:
		file := filepath.Base(srcPath)
		return nil, fmt.Errorf("Unsupported source code: %s", file)
	}
}

type ProgramGo struct{}

func (pg *ProgramGo) ShouldCompile() bool {
	return true
}
func (pg *ProgramGo) GetCompileCmds(src string, dest string) []string {
	return []string{"go", "build", "-o", dest, src}
}
func (pg *ProgramGo) GetExecCmds(execPath string) []string {
	return []string{execPath}
}
