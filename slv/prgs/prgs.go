package prgs

import (
	"fmt"
	"path/filepath"

	"github.com/ryym/slv/slv/t"
)

func FindExtByLang(lang string) string {
	pairs := []struct {
		lang string
		ext  string
	}{
		{"go", "go"},
		{"ruby", "rb"},
		{"scala", "scala"},
	}

	for _, p := range pairs {
		if p.lang == lang {
			return "." + p.ext
		}
	}
	return ""
}

func FindProgram(srcPath string) (t.Program, error) {
	ext := filepath.Ext(srcPath)
	prg := NewProgramByExt(ext)

	if prg != nil {
		return prg, nil
	} else {
		file := filepath.Base(srcPath)
		return nil, fmt.Errorf("Unsupported source code: %s", file)
	}
}

func NewProgramByExt(ext string) t.Program {
	switch ext {
	case ".go":
		return &ProgramGo{}
	case ".rb":
		return &ProgramRuby{}
	case ".scala":
		return &ProgramScala{}
	}
	return nil
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
