package main

import (
	"fmt"
	"os"

	"github.com/ryym/slv/cli"
)

func main() {
	app := cli.CreateApp()
	err := app.Run(os.Args)

	if err != nil {
		if err.Error() != "" {
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
}
