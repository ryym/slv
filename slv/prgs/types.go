package prgs

type programDef interface {
	GetCompileCmds(srcPath string, destDir string) compileCmds
	GetExecCmds(execPath string) []string
}

type compileCmds struct {
	Cmds     []string
	ExecPath string
}
