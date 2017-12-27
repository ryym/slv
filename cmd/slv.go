package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		{
			Name:      "compile",
			Usage:     "Compile without running",
			ArgsUsage: "[src]",
			Action:    cmdCompile,
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

	ok, err := slv.TestAll(&conf)

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

	conf, err := slv.MakeExecConf(c.Args()[0])
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	execPath, err := slv.Compile(&conf)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	// Try to use relative path.
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		relPath, err := filepath.Rel(pwd, execPath)
		if err == nil && len(relPath) < len(execPath) {
			execPath = relPath
		}
	}

	fmt.Printf("Compiled to %s\n", execPath)
	return nil
}
