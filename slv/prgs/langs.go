package prgs

import "path/filepath"

func newProgramDef(fileName string) programDef {
	switch filepath.Ext(fileName) {
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

func (pg *cmdsGo) GetCompileCmds(src string, destDir string) compileCmds {
	bin := filepath.Join(destDir, "out")
	return compileCmds{
		Cmds:     []string{"go", "build", "-o", bin, src},
		ExecPath: bin,
	}
}
func (pg *cmdsGo) GetExecCmds(execPath string) []string {
	return []string{execPath}
}

type cmdsRuby struct{}

func (pg *cmdsRuby) GetCompileCmds(src string, _destDir string) compileCmds {
	return compileCmds{
		Cmds:     nil,
		ExecPath: src,
	}
}
func (pg *cmdsRuby) GetExecCmds(execPath string) []string {
	return []string{"ruby", execPath}
}

type cmdsScala struct{}

func (pg *cmdsScala) GetCompileCmds(src string, destDir string) compileCmds {
	return compileCmds{
		Cmds:     []string{"scalac", "-d", destDir, src},
		ExecPath: destDir,
	}
}
func (pg *cmdsScala) GetExecCmds(execPath string) []string {
	// It is better if users can configure output class name
	// (currently 'Main').
	return []string{"scala", "-cp", execPath, "Main"}
}
