// This package defines all structs and interfaces
// used in other several packages to avoid cyclic imports.

package t

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
