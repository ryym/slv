package slv

import (
	"os"

	"github.com/pkg/errors"
)

func Def() {}

func MakeDir(opts CmdNewOpts) (err error) {
	name := opts.Name
	err = os.Mkdir(name, 0755)
	if err != nil {
		return errors.Wrapf(err, "Failed to create %s directory", name)
	}
	for _, d := range []string{"src", "test", ".slv"} {
		err = os.Mkdir(name+"/"+d, 0755)
		if err != nil {
			return errors.Wrapf(err, "Failed to create %s directory", d)
		}
	}
	return nil
}
