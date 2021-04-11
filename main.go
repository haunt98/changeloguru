package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	appName = "changeloguru"

	// flags
	fromFlag       = "from"
	toFlag         = "to"
	versionFlag    = "version"
	scopeFlag      = "scope"
	repositoryFlag = "repository"
	outputFlag     = "output"
	filenameFlag   = "filename"
	filetypeFlag   = "filetype"
	verboseFlag    = "verbose"

	// commands
	generateCommand = "generate"
)

var (
	fmtErr = color.New(color.FgRed)

	// flags
	verboseAliases = []string{"v"}

	// commands
	generateAliases = []string{"gen"}
)

func main() {
	a := &action{}

	app := &cli.App{
		Name:  appName,
		Usage: "generate changelog from conventional commits",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    verboseFlag,
				Aliases: verboseAliases,
				Usage:   "show what is going on",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    generateCommand,
				Aliases: generateAliases,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  fromFlag,
						Usage: "generate from `COMMIT`",
					},
					&cli.StringFlag{
						Name:  toFlag,
						Usage: "generate to `COMMIT`",
					},
					&cli.StringFlag{
						Name:  versionFlag,
						Usage: "`VERSION` to generate, follow Semantic Versioning",
					},
					&cli.StringSliceFlag{
						Name:  scopeFlag,
						Usage: "scope to generate",
					},
					&cli.StringFlag{
						Name:        repositoryFlag,
						Usage:       "`REPOSITORY` directory path",
						DefaultText: defaultRepository,
					},
					&cli.StringFlag{
						Name:        outputFlag,
						Usage:       "`OUTPUT` directory path",
						DefaultText: defaultOutput,
					},
					&cli.StringFlag{
						Name:        filenameFlag,
						Usage:       "output `FILENAME`",
						DefaultText: defaultFilename,
					},
					&cli.StringFlag{
						Name:        filetypeFlag,
						Usage:       "output `FILETYPE`",
						DefaultText: defaultFiletype,
					},
				},
				Action: a.runGenerate,
			},
		},
		Action: a.runHelp,
	}

	if err := app.Run(os.Args); err != nil {
		// Highlight error
		fmtErr.Printf("[%s error]: ", appName)
		fmt.Printf("%s\n", err.Error())
	}
}
