package cli

import (
	"errors"
	"os"

	"github.com/ryym/slv/slv"
	"github.com/ryym/slv/slv/tp"
	"github.com/urfave/cli"
)

func CreateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "slv"
	app.Version = "0.1.0"
	app.Usage = "Helps you solve programming problems"
	app.Commands = []cli.Command{
		{
			Name:      "new",
			Aliases:   []string{"n"},
			Usage:     "Create new problem directory",
			ArgsUsage: "[directory]",
			Action:    cmdNew,
		},
		{
			Name:      "test",
			Aliases:   []string{"t"},
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
			Aliases:   []string{"r"},
			Usage:     "Run the specified source code",
			ArgsUsage: "[src|lang]",
			Action:    cmdRun,
		},
	}
	return app
}

func cmdNew(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "new")
		return nil
	}
	name := c.Args()[0]
	return slv.New(&tp.CmdNewOpts{
		Name: name,
	})
}

func cmdTest(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "test")
		return nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	app, err := slv.NewSlvApp(c.Args()[0], wd)
	if err != nil {
		return err
	}

	ok, err := slv.Test(&app)

	if err != nil {
		return err
	}
	if !ok {
		return errors.New("")
	}
	return nil
}

func cmdCompile(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "compile")
		return nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	app, err := slv.NewSlvApp(c.Args()[0], wd)
	if err != nil {
		return err
	}

	err = slv.Compile(&app)
	if err != nil {
		return err
	}

	return nil
}

func cmdRun(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "run")
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	app, err := slv.NewSlvApp(c.Args()[0], wd)
	if err != nil {
		return err
	}

	err = slv.Run(&app)

	if err != nil {
		return err
	}
	return nil
}
