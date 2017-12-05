package cmd

import (
	"errors"

	"github.com/urfave/cli"
)

func CreateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "slv"
	app.Usage = "Helps you solve programming problems"
	app.Commands = []cli.Command{}
	app.Action = func(c *cli.Context) error {
		return errors.New("Not implemented yet")
	}
	return app
}
