// This package defines all structs and interfaces
// used in other several packages to avoid cyclic imports.

package tp

type CmdNewOpts struct {
	Name string
}

type Slv struct {
	ProbDir ProbDir
}

type ProbDir interface {
	WorkDir() string
	SrcDir() string
	TestDir() string
	DestDir() string
	SrcFile() string
	SrcPath() string
}

type ProgramCmds interface {
	GetCompileCmds(srcPath string, destDir string) CompileCmds
	GetExecCmds(execPath string) []string
}

type CompileCmds struct {
	Cmds     []string
	ExecPath string
}
