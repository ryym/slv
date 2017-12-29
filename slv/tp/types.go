// This package defines all structs and interfaces
// used in other several packages to avoid cyclic imports.

package tp

type CmdNewOpts struct {
	Name string
}

type Probdir struct {
	RootDir string
	SrcFile string
	SrcPath string
	WorkDir string
}

type ExecConf struct {
	Probdir
}

type Program interface {
	GetCompileCmds(srcPath string, destDir string) CompileCmds
	GetExecCmds(execPath string) []string
}

type CompileCmds struct {
	Cmds     []string
	ExecPath string
}
