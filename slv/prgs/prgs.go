package prgs

import (
	"fmt"
	"path/filepath"

	"github.com/ryym/slv/slv/tp"
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

func FindProgram(srcPath string) (tp.ProgramCmds, error) {
	ext := filepath.Ext(srcPath)
	prg := NewProgramByExt(ext)

	if prg != nil {
		return prg, nil
	} else {
		file := filepath.Base(srcPath)
		return nil, fmt.Errorf("Unsupported source code: %s", file)
	}
}

func NewProgramByExt(ext string) tp.ProgramCmds {
	switch ext {
	case ".go":
		return &cmdsGo{}
	case ".rb":
		return &cmdsRuby{}
	case ".scala":
		return &cmdsScala{}
	}
	return nil
}

type cmdsGo struct{}

func (pg *cmdsGo) GetCompileCmds(src string, destDir string) tp.CompileCmds {
	bin := filepath.Join(destDir, "out")
	return tp.CompileCmds{
		Cmds:     []string{"go", "build", "-o", bin, src},
		ExecPath: bin,
	}
}
func (pg *cmdsGo) GetExecCmds(execPath string) []string {
	return []string{execPath}
}

type cmdsRuby struct{}

func (pg *cmdsRuby) GetCompileCmds(src string, _destDir string) tp.CompileCmds {
	return tp.CompileCmds{
		Cmds:     nil,
		ExecPath: src,
	}
}
func (pg *cmdsRuby) GetExecCmds(execPath string) []string {
	return []string{"ruby", execPath}
}

type cmdsScala struct{}

func (pg *cmdsScala) GetCompileCmds(src string, destDir string) tp.CompileCmds {
	return tp.CompileCmds{
		Cmds:     []string{"scalac", "-d", destDir, src},
		ExecPath: destDir,
	}
}
func (pg *cmdsScala) GetExecCmds(execPath string) []string {
	// It is better if users can configure output class name
	// (currently 'Main').
	return []string{"scala", "-cp", execPath, "Main"}
}
