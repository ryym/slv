package prgs

type programDef interface {
	GetCompileCmds(srcPath string, destDir string) []string
	GetExecCmds(srcPath string, destDir string) []string
}
