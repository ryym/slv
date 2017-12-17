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
	case ".rb":
		return &ProgramRuby{}, nil
	default:
		file := filepath.Base(srcPath)
		return nil, fmt.Errorf("Unsupported source code: %s", file)
	}
}

type ProgramGo struct{}

func (pg *ProgramGo) GetCompileCmds(src string, dest string) []string {
	return []string{"go", "build", "-o", dest, src}
}
func (pg *ProgramGo) GetExecCmds(execPath string) []string {
	return []string{execPath}
}

type ProgramRuby struct{}

func (pg *ProgramRuby) GetCompileCmds(src string, dest string) []string {
	return nil
}
func (pg *ProgramRuby) GetExecCmds(execPath string) []string {
	return []string{"ruby", execPath}
}
