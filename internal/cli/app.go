package cli

import (
	"os"

	"github.com/make-go-great/color-go"
	"github.com/urfave/cli/v2"
)

const (
	name  = "changeloguru"
	usage = "generate changelog from conventional commits"

	flagVerbose     = "verbose"
	flagFrom        = "from"
	flagTo          = "to"
	flagVersion     = "version"
	flagScope       = "scope"
	flagRepository  = "repository"
	flagOutput      = "output"
	flagFilename    = "filename"
	flagFiletype    = "filetype"
	flagDryRun      = "dry-run"
	flagInteractive = "interactive"

	commandGenerate = "generate"

	usageGenerate    = "generate changelog"
	usageVerbose     = "show what is going on"
	usageFrom        = "from `COMMIT`, which is kinda new commit"
	usageTo          = "to `COMMIT`, which is kinda old commit"
	usageVersion     = "`VERSION` to generate, follow Semantic Versioning"
	usageScope       = "scope to generate"
	usageRepository  = "`REPOSITORY` directory path"
	usageOutput      = "`OUTPUT` directory path"
	usageFilename    = "output `FILENAME`"
	usageFiletype    = "output `FILETYPE`"
	usageDryRun      = "demo run without actually changing anything"
	usageInteractive = "interactive mode"
)

var (
	aliasGenerate    = []string{"gen"}
	aliasVerbose     = []string{"v"}
	aliasInteractive = []string{"i"}
)

type App struct {
	cliApp *cli.App
}

func NewApp() *App {
	a := &action{}

	cliApp := &cli.App{
		Name:  name,
		Usage: usage,
		Commands: []*cli.Command{
			{
				Name:    commandGenerate,
				Aliases: aliasGenerate,
				Usage:   usageGenerate,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    flagVerbose,
						Aliases: aliasVerbose,
						Usage:   usageVerbose,
					},
					&cli.StringFlag{
						Name:  flagFrom,
						Usage: usageFrom,
					},
					&cli.StringFlag{
						Name:  flagTo,
						Usage: usageTo,
					},
					&cli.StringFlag{
						Name:  flagVersion,
						Usage: usageVersion,
					},
					&cli.StringSliceFlag{
						Name:  flagScope,
						Usage: usageScope,
					},
					&cli.StringFlag{
						Name:  flagRepository,
						Usage: usageRepository,
						Value: defaultRepository,
					},
					&cli.StringFlag{
						Name:  flagOutput,
						Usage: usageOutput,
						Value: defaultOutput,
					},
					&cli.StringFlag{
						Name:  flagFilename,
						Usage: usageFilename,
						Value: defaultFilename,
					},
					&cli.StringFlag{
						Name:  flagFiletype,
						Usage: usageFiletype,
						Value: defaultFiletype,
					},
					&cli.BoolFlag{
						Name:  flagDryRun,
						Usage: usageDryRun,
					},
					&cli.BoolFlag{
						Name:    flagInteractive,
						Usage:   usageInteractive,
						Aliases: aliasInteractive,
					},
				},
				Action: a.RunGenerate,
			},
		},
		Action: a.RunHelp,
	}

	return &App{
		cliApp: cliApp,
	}
}

func (a *App) Run() {
	if err := a.cliApp.Run(os.Args); err != nil {
		color.PrintAppError(name, err.Error())
	}
}
