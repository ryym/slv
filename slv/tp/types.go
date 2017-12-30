// This package defines all structs and interfaces
// used in other several packages to avoid cyclic imports.

package tp

//go:generate moq -out mocks.go . Program

import "io"

type CmdNewOpts struct {
	Name string
}

type Slv struct {
	ProbDir ProbDir
	Program ProgramFactory
}

type ProbDir interface {
	WorkDir() string
	SrcDir() string
	TestDir() string
	DestDir() string
	SrcFile() string
	SrcPath() string
}

type ProgramFactory interface {
	NewProgram(srcPath string, destDir string) (Program, error)
}

type CompileResult struct {
	Compiled bool
	ExecPath string
}

type Program interface {
	Compile() (CompileResult, error)
	Run(input string) (string, error)
	RunWithPipes(stdin io.ReadCloser, stdout io.WriteCloser) error
}