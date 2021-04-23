package cli

import (
	"os"

	"github.com/haunt98/color"
	"github.com/urfave/cli/v2"
)

const (
	appName  = "changeloguru"
	appUsage = "generate changelog from conventional commits"

	// flags
	verboseFlag    = "verbose"
	fromFlag       = "from"
	toFlag         = "to"
	versionFlag    = "version"
	scopeFlag      = "scope"
	repositoryFlag = "repository"
	outputFlag     = "output"
	filenameFlag   = "filename"
	filetypeFlag   = "filetype"
	dryRunFlag     = "dry-run"

	// commands
	generateCommand = "generate"

	// flag usage
	verboseUsage    = "show what is going on"
	fromUsage       = "generate from `COMMIT`"
	toUsage         = "generate to `COMMIT`"
	versionUsage    = "`VERSION` to generate, follow Semantic Versioning"
	scopeUsage      = "scope to generate"
	repositoryUsage = "`REPOSITORY` directory path"
	outputUsage     = "`OUTPUT` directory path"
	filenameUsage   = "output `FILENAME`"
	filetypeUsage   = "output `FILETYPE`"
	dryRunUsage     = "demo run without actually changing anything"

	// command usage
	generateUsage = "generate changelog"
)

var (
	// flag aliases
	verboseAliases = []string{"v"}

	// command aliases
	generateAliases = []string{"gen"}
)

type App struct {
	cliApp *cli.App
}

func NewApp() *App {
	a := &action{}

	cliApp := &cli.App{
		Name:  appName,
		Usage: appUsage,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    verboseFlag,
				Aliases: verboseAliases,
				Usage:   verboseUsage,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    generateCommand,
				Aliases: generateAliases,
				Usage:   generateUsage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  fromFlag,
						Usage: fromUsage,
					},
					&cli.StringFlag{
						Name:  toFlag,
						Usage: toUsage,
					},
					&cli.StringFlag{
						Name:  versionFlag,
						Usage: versionUsage,
					},
					&cli.StringSliceFlag{
						Name:  scopeFlag,
						Usage: scopeUsage,
					},
					&cli.StringFlag{
						Name:        repositoryFlag,
						Usage:       repositoryUsage,
						DefaultText: defaultRepository,
					},
					&cli.StringFlag{
						Name:        outputFlag,
						Usage:       outputUsage,
						DefaultText: defaultOutput,
					},
					&cli.StringFlag{
						Name:        filenameFlag,
						Usage:       filenameUsage,
						DefaultText: defaultFilename,
					},
					&cli.StringFlag{
						Name:        filetypeFlag,
						Usage:       filetypeUsage,
						DefaultText: defaultFiletype,
					},
					&cli.BoolFlag{
						Name:  dryRunFlag,
						Usage: dryRunUsage,
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
		color.PrintAppError(appName, err.Error())
	}
}
