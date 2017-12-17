package slv

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/t"
)

func Compile(c *t.ExecConf) (binpath string, err error) {
	binpath = filepath.Join(
		c.WorkDir,
		fmt.Sprintf("%s.%s", c.SrcFile, "compiled"),
	)

	// TODO: Support other languages.
	cmd := exec.Command("go", "build", "-o", binpath, c.SrcPath)

	_, err = cmd.Output()
	if err != nil {
		return "", errors.Wrapf(err, "Failed to compile %s", c.SrcFile)
	}
	return binpath, nil
}
