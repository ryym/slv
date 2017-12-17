package cmd

import (
	"github.com/ryym/slv/slv"
	"github.com/ryym/slv/slv/t"
	"github.com/urfave/cli"
)

func CreateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "slv"
	app.Usage = "Helps you solve programming problems"
	app.Commands = []cli.Command{
		{
			Name:      "new",
			Usage:     "Create new problem directory",
			ArgsUsage: "[directory]",
			Action:    cmdNew,
		},
		{
			Name:      "test",
			Usage:     "Run tests for the specified source code",
			ArgsUsage: "[src]",
			Action:    cmdTest,
		},
	}
	return app
}

func cmdNew(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelpAndExit(c, "new", 0)
	}
	name := c.Args()[0]
	return slv.MakeDir(t.CmdNewOpts{
		Name: name,
	})
}

func cmdTest(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelpAndExit(c, "test", 0)
	}

	conf, err := slv.MakeExecConf(c.Args()[0])
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = slv.TestAll(&conf)

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
