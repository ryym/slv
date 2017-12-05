package main

import (
	"os"

	"github.com/ryym/slv/cmd"
)

func main() {
	app := cmd.CreateApp()
	app.Run(os.Args)
}
