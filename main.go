package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/haunt98/changeloguru/pkg/cli"
)

var (
	fmtErr = color.New(color.FgRed)
)

func main() {
	app := cli.NewApp()
	if err := app.Run(os.Args); err != nil {
		// Highlight error
		fmtErr.Printf("[%s error]: ", cli.AppName)
		fmt.Printf("%s\n", err.Error())
	}
}
