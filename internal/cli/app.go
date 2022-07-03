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
	flagVersion     = "version"
	flagFrom        = "from"
	flagTo          = "to"
	flagScope       = "scope"
	flagRepository  = "repository"
	flagOutput      = "output"
	flagFilename    = "filename"
	flagFiletype    = "filetype"
	flagDryRun      = "dry-run"
	flagInteractive = "interactive"
	flagAutoCommit  = "auto-commit"

	commandGenerate = "generate"

	usageCommandGenerate = "generate changelog"
	usageFlagVerbose     = "show what is going on"
	usageFlagVersion     = "`VERSION` to generate, follow Semantic Versioning"
	usageFlagFrom        = "from `COMMIT`, which is kinda new commit, default is latest commit"
	usageFlagTo          = "to `COMMIT`, which is kinda old commit, default is oldest commit"
	usageFlagScope       = "scope to generate"
	usageFlagRepository  = "`REPOSITORY` directory path"
	usageFlagOutput      = "`OUTPUT` directory path"
	usageFlagFilename    = "output `FILENAME`"
	usageFlagFiletype    = "output `FILETYPE`"
	usageFlagDryRun      = "demo run without actually changing anything"
	usageFlagInteractive = "interactive mode"
	usageFlagAutoCommit  = "enable auto commit after generating changelog"
)

var (
	aliasCommandGenerate = []string{"g", "gen"}
	aliasFlagVerbose     = []string{"v"}
	aliasFlagInteractive = []string{"i"}
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
				Aliases: aliasCommandGenerate,
				Usage:   usageCommandGenerate,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    flagVerbose,
						Aliases: aliasFlagVerbose,
						Usage:   usageFlagVerbose,
					},
					&cli.StringFlag{
						Name:  flagVersion,
						Usage: usageFlagVersion,
					},
					&cli.StringFlag{
						Name:  flagFrom,
						Usage: usageFlagFrom,
					},
					&cli.StringFlag{
						Name:  flagTo,
						Usage: usageFlagTo,
					},
					&cli.StringSliceFlag{
						Name:  flagScope,
						Usage: usageFlagScope,
					},
					&cli.StringFlag{
						Name:  flagRepository,
						Usage: usageFlagRepository,
						Value: defaultRepository,
					},
					&cli.StringFlag{
						Name:  flagOutput,
						Usage: usageFlagOutput,
						Value: defaultOutput,
					},
					&cli.StringFlag{
						Name:  flagFilename,
						Usage: usageFlagFilename,
						Value: defaultFilename,
					},
					&cli.StringFlag{
						Name:  flagFiletype,
						Usage: usageFlagFiletype,
						Value: defaultFiletype,
					},
					&cli.BoolFlag{
						Name:  flagDryRun,
						Usage: usageFlagDryRun,
					},
					&cli.BoolFlag{
						Name:    flagInteractive,
						Usage:   usageFlagInteractive,
						Aliases: aliasFlagInteractive,
						Value:   true,
					},
					&cli.BoolFlag{
						Name:  flagAutoCommit,
						Usage: usageFlagAutoCommit,
						Value: true,
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
