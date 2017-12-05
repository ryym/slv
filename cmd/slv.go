package cmd

import (
	"github.com/ryym/slv/slv"
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
	}
	return app
}

func cmdNew(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelpAndExit(c, "new", 0)
	}
	name := c.Args()[0]
	return slv.MakeDir(slv.CmdNewOpts{
		Name: name,
	})
}
