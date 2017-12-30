package prgs

import (
	"fmt"
	"path/filepath"

	"github.com/ryym/slv/slv/tp"
)

func NewProgramFactory() tp.ProgramFactory {
	return &programFactoryImpl{newProgramDef}
}

type programFactoryImpl struct {
	newProgramDef func(string) programDef
}

func (pf *programFactoryImpl) NewProgram(srcPath string, destDir string) (tp.Program, error) {
	def := pf.newProgramDef(srcPath)
	if def == nil {
		return nil, fmt.Errorf("Unsupported source code: %s", filepath.Base(srcPath))
	}

	return &programImpl{
		srcPath: srcPath,
		destDir: destDir,
		def:     def,
	}, nil
}
