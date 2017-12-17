package slv

import (
	"github.com/ryym/slv/slv/t"
	"github.com/ryym/slv/slv/test"
)

func TestAll(c *t.ExecConf) error {
	result, err := test.TestAll(c)
	if err != nil {
		return err
	}
	return test.ShowResult(&result)
}
