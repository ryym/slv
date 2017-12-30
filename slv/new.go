package slv

import (
	"os"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/probdir"
	"github.com/ryym/slv/slv/tp"
)

func NewProblem(opts *tp.CmdNewOpts) (err error) {
	name := opts.Name
	err = os.Mkdir(name, 0755)
	if err != nil {
		return errors.Wrapf(err, "Failed to create %s directory", name)
	}

	expectedDirs := []string{probdir.SRC_DIR, probdir.TEST_DIR, probdir.WORK_DIR}
	for _, d := range expectedDirs {
		err = os.Mkdir(name+"/"+d, 0755)
		if err != nil {
			return errors.Wrapf(err, "Failed to create %s directory", d)
		}
	}
	return nil
}
