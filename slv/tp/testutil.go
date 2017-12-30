package tp

import "io"

// This file defines some mock structs for testing.
// `Fk` is an abbriviation of `Fake`.

type FkProgram struct {
	FkCompile      func() (CompileResult, error)
	FkRun          func(string) (string, error)
	FkRunWithPipes func(io.ReadCloser, io.WriteCloser) error
}

func (p *FkProgram) Compile() (CompileResult, error) {
	return p.FkCompile()
}
func (p *FkProgram) Run(input string) (string, error) {
	return p.FkRun(input)
}
func (p *FkProgram) RunWithPipes(stdin io.ReadCloser, stdout io.WriteCloser) error {
	return p.FkRunWithPipes(stdin, stdout)
}
