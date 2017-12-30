package main

import (
	"os"

	"github.com/ryym/slv/cli"
)

func main() {
	app := cli.CreateApp()
	app.Run(os.Args)
}
