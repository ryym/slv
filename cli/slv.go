package cli

import (
	"os"

	"github.com/ryym/slv/slv"
	"github.com/ryym/slv/slv/tp"
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
			ArgsUsage: "[src|lang]",
			Action:    cmdTest,
		},
		{
			Name:      "compile",
			Usage:     "Compile without running",
			ArgsUsage: "[src|lang]",
			Action:    cmdCompile,
		},
		{
			Name:      "run",
			Usage:     "Run the specified source code",
			ArgsUsage: "[src|lang]",
			Action:    cmdRun,
		},
	}
	return app
}

func cmdNew(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelpAndExit(c, "new", 0)
	}
	name := c.Args()[0]
	return slv.New(&tp.CmdNewOpts{
		Name: name,
	})
}

func cmdTest(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelpAndExit(c, "test", 0)
	}

	wd, err := os.Getwd()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	app, err := slv.NewSlvApp(c.Args()[0], wd)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	ok, err := slv.Test(&app)

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	if !ok {
		os.Exit(1)
	}
	return nil
}

func cmdCompile(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelpAndExit(c, "compile", 0)
	}

	wd, err := os.Getwd()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	app, err := slv.NewSlvApp(c.Args()[0], wd)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = slv.Compile(&app)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}

func cmdRun(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelpAndExit(c, "run", 0)
	}

	wd, err := os.Getwd()
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	app, err := slv.NewSlvApp(c.Args()[0], wd)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	err = slv.Run(&app)

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
