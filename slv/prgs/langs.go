package prgs

import "path/filepath"

func FindExtByLang(lang string) string {
	pairs := []struct {
		lang string
		ext  string
	}{
		{"go", "go"},
		{"ruby", "rb"},
		{"scala", "scala"},
	}

	for _, p := range pairs {
		if p.lang == lang {
			return "." + p.ext
		}
	}
	return ""
}

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

func (pg *cmdsGo) GetCompileCmds(src string, destDir string) []string {
	bin := filepath.Join(destDir, "out")
	return []string{"go", "build", "-o", bin, src}
}
func (pg *cmdsGo) GetExecCmds(_src string, destDir string) []string {
	return []string{filepath.Join(destDir, "out")}
}

type cmdsRuby struct{}

func (pg *cmdsRuby) GetCompileCmds(src string, _destDir string) []string {
	return nil
}
func (pg *cmdsRuby) GetExecCmds(src string, _destDir string) []string {
	return []string{"ruby", src}
}

type cmdsScala struct{}

func (pg *cmdsScala) GetCompileCmds(src string, destDir string) []string {
	return []string{"scalac", "-d", destDir, src}
}
func (pg *cmdsScala) GetExecCmds(_src string, destDir string) []string {
	// It is better if users can configure output class name
	// (currently 'Main').
	return []string{"scala", "-cp", destDir, "Main"}
}
