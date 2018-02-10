package prgs

import (
	tmpl "text/template"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/tp"
)

type programDictImpl struct {
	prgs []*dynamicProgramDef
}

func (pd *programDictImpl) FindDefByExt(ext string) tp.ProgramDef {
	for _, p := range pd.prgs {
		for _, e := range p.exts {
			if e == ext {
				return p
			}
		}
	}
	return nil
}

func (pd *programDictImpl) FindExts(nameOrExt string) ([]string, bool) {
	maybeExt := "." + nameOrExt
	for _, p := range pd.prgs {
		found := false

		if p.name == nameOrExt {
			found = true
		} else {
			for _, ext := range p.exts {
				if ext == maybeExt {
					found = true
					break
				}
			}
		}

		if found {
			exts := make([]string, len(p.exts))
			copy(exts, p.exts)
			return exts, true
		}
	}
	return nil, false
}

func MakeProgramDict(conf *tp.Config) (tp.ProgramDict, error) {
	prgs := make([]*dynamicProgramDef, len(conf.Langs))
	i := 0

	for name, def := range conf.Langs {
		var compileTmpl *tmpl.Template
		var err error

		if def.Compile != "" {
			compileTmpl, err = tmpl.New("compile").Parse(def.Compile)
		}
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse 'compile' of %s", name)
		}

		execTmpl, err := tmpl.New("exec").Parse(def.Run)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse 'run' of %s", name)
		}

		prgs[i] = &dynamicProgramDef{
			name:        name,
			compileTmpl: compileTmpl,
			execTmpl:    execTmpl,
			exts:        def.Exts,
		}
		i += 1
	}

	return &programDictImpl{prgs}, nil
}
