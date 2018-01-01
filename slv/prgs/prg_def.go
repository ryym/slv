package prgs

import (
	"bytes"
	"path/filepath"
	"runtime"
	"strings"
	tmpl "text/template"
)

type pathsArg struct {
	Src  string
	Dest string
}

func (pa *pathsArg) Join(paths ...string) string {
	return filepath.Join(paths...)
}

func (pa *pathsArg) Out() string {
	if runtime.GOOS == "windows" {
		return "out.exe"
	} else {
		return "out"
	}
}

type dynamicProgramDef struct {
	name        string
	compileTmpl *tmpl.Template
	execTmpl    *tmpl.Template
	aliases     []string
	exts        []string
}

func (p *dynamicProgramDef) GetCompileCmds(srcPath string, destDir string) ([]string, error) {
	if p.compileTmpl == nil {
		return nil, nil
	}
	return p.execCmdTmpl(p.compileTmpl, &pathsArg{srcPath, destDir})
}

func (p *dynamicProgramDef) GetExecCmds(srcPath string, destDir string) ([]string, error) {
	return p.execCmdTmpl(p.execTmpl, &pathsArg{srcPath, destDir})
}

func (p *dynamicProgramDef) execCmdTmpl(t *tmpl.Template, data interface{}) ([]string, error) {
	var out bytes.Buffer
	err := t.Execute(&out, data)
	if err != nil {
		return nil, err
	}
	return strings.Split(out.String(), " "), nil
}
