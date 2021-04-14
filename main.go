package main

import (
	"os"

	"github.com/haunt98/changeloguru/pkg/cli"
	"github.com/haunt98/color"
)

func main() {
	app := cli.NewApp()
	if err := app.Run(os.Args); err != nil {
		color.PrintAppError(cli.AppName, err.Error())
	}
}
