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
	Program Program
}

type ProbDir interface {
	WorkDir() string
	SrcDir() string
	TestDir() string
	DestDir() string
	SrcFile() string
	SrcPath() string
}

type CompileResult struct {
	Compiled bool
	Output   []byte
}

type Program interface {
	Compile() (CompileResult, error)
	Run(input string) (string, error)
	RunWithPipes(stdin io.ReadCloser, stdout io.WriteCloser) error
}

type ConfigLoader interface {
	Load(r io.Reader) (*Config, error)
	LoadFromFiles(baseDir string) ([]*Config, error)
	Merge(a *Config, b *Config) *Config
}

type Config struct {
	Langs map[string]ConfLang `toml:"lang"`
}

type ConfLang struct {
	Compile string
	Run     string
	Exts    []string
}

type ProgramDict interface {
	FindExts(lang string) ([]string, bool)
	FindDefByExt(ext string) ProgramDef
}
type ProgramDef interface {
	GetCompileCmds(srcPath string, destDir string) ([]string, error)
	GetExecCmds(srcPath string, destDir string) ([]string, error)
}
