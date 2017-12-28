package slv

import (
	"fmt"
	"path/filepath"

	"github.com/ryym/slv/slv/t"
	"github.com/ryym/slv/slv/test"
)

func TestAll(c *t.ExecConf) (bool, error) {
	fmt.Printf("testing %s...\n", c.SrcFile)

	execCmds, err := findAndCompile(c)
	if err != nil {
		return false, err
	}

	testdir := filepath.Join(c.RootDir, "test")
	printer := test.NewResultPrinter()
	return test.TestAll(execCmds, testdir, printer)
}
