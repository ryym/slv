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
	case ".scala":
		return &ProgramScala{}, nil
	default:
		file := filepath.Base(srcPath)
		return nil, fmt.Errorf("Unsupported source code: %s", file)
	}
}

type ProgramGo struct{}

func (pg *ProgramGo) GetCompileCmds(src string, destDir string) t.CompileCmds {
	bin := filepath.Join(destDir, "out")
	return t.CompileCmds{
		Cmds:     []string{"go", "build", "-o", bin, src},
		ExecPath: bin,
	}
}
func (pg *ProgramGo) GetExecCmds(execPath string) []string {
	return []string{execPath}
}

type ProgramRuby struct{}

func (pg *ProgramRuby) GetCompileCmds(src string, _destDir string) t.CompileCmds {
	return t.CompileCmds{
		Cmds:     nil,
		ExecPath: src,
	}
}
func (pg *ProgramRuby) GetExecCmds(execPath string) []string {
	return []string{"ruby", execPath}
}

type ProgramScala struct{}

func (pg *ProgramScala) GetCompileCmds(src string, destDir string) t.CompileCmds {
	return t.CompileCmds{
		Cmds:     []string{"scalac", "-d", destDir, src},
		ExecPath: destDir,
	}
}
func (pg *ProgramScala) GetExecCmds(execPath string) []string {
	// It is better if users can configure output class name
	// (currently 'Main').
	return []string{"scala", "-cp", execPath, "Main"}
}
