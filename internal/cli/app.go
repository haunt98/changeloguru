package cli

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/make-go-great/color-go"
)

const (
	name  = "changeloguru"
	usage = "generate changelog from conventional commits"

	commandGenerateName  = "generate"
	commandGenerateUsage = "generate changelog"

	flagVerboseName  = "verbose"
	flagVerboseUsage = "show what is going on"

	flagVersionName  = "version"
	flagVersionUsage = "`VERSION` to generate, follow Semantic Versioning"

	flagFromName  = "from"
	flagFromUsage = "from `COMMIT`, which is kinda new commit, default is latest commit"

	flagToName  = "to"
	flagToUsage = "to `COMMIT`, which is kinda old commit, default is oldest commit"

	flagScopeName  = "scope"
	flagScopeUsage = "scope to generate"

	flagRepositoryName  = "repository"
	flagRepositoryUsage = "`REPOSITORY` directory path"

	flagOutputName  = "output"
	flagOutputUsage = "`OUTPUT` directory path, relative to `REPOSITORY` path"

	flagFilenameName  = "filename"
	flagFilenameUsage = "output `FILENAME`"

	flagFiletypeName  = "filetype"
	flagFiletypeUsage = "output `FILETYPE`"

	flagDryRunName  = "dry-run"
	flagDryRunUsage = "demo run without actually changing anything"

	flagInteractiveName  = "interactive"
	flagInteractiveUsage = "interactive mode"

	flagInteractiveFromName  = "interactive-from"
	flagInteractiveFromUsage = "enable ask from in interactive mode"

	flagAutoGitCommitName  = "auto-commit"
	flagAutoGitCommitUsage = "enable auto git commit after generating changelog"

	flagAutoGitTagName  = "auto-tag"
	flagAutoGitTagUsage = "enable auto git tag after generating changelog, only works if auto-commit is enabled"

	flagAutoGitPushName  = "auto-push"
	flagAutoGitPushUsage = "enable auto git push after generating changelog, only works if auto-commit is enabled, if auto-tag is enabled will auto git push tag too"
)

var (
	commandGenerateAliases = []string{"g", "gen"}
	flagVerboseAliases     = []string{"v"}
	flagInteractiveAliases = []string{"i"}
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
				Name:    commandGenerateName,
				Aliases: commandGenerateAliases,
				Usage:   commandGenerateUsage,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    flagVerboseName,
						Aliases: flagVerboseAliases,
						Usage:   flagVerboseUsage,
					},
					&cli.StringFlag{
						Name:  flagVersionName,
						Usage: flagVersionUsage,
					},
					&cli.StringFlag{
						Name:  flagFromName,
						Usage: flagFromUsage,
					},
					&cli.StringFlag{
						Name:  flagToName,
						Usage: flagToUsage,
					},
					&cli.StringSliceFlag{
						Name:  flagScopeName,
						Usage: flagScopeUsage,
					},
					&cli.StringFlag{
						Name:  flagRepositoryName,
						Usage: flagRepositoryUsage,
					},
					&cli.StringFlag{
						Name:  flagOutputName,
						Usage: flagOutputUsage,
					},
					&cli.StringFlag{
						Name:  flagFilenameName,
						Usage: flagFilenameUsage,
					},
					&cli.StringFlag{
						Name:  flagFiletypeName,
						Usage: flagFiletypeUsage,
					},
					&cli.BoolFlag{
						Name:  flagDryRunName,
						Usage: flagDryRunUsage,
					},
					&cli.BoolFlag{
						Name:    flagInteractiveName,
						Usage:   flagInteractiveUsage,
						Aliases: flagInteractiveAliases,
					},
					&cli.BoolFlag{
						Name:  flagInteractiveFromName,
						Usage: flagInteractiveFromUsage,
					},
					&cli.BoolFlag{
						Name:  flagAutoGitCommitName,
						Usage: flagAutoGitCommitUsage,
					},
					&cli.BoolFlag{
						Name:  flagAutoGitTagName,
						Usage: flagAutoGitTagUsage,
					},
					&cli.BoolFlag{
						Name:  flagAutoGitPushName,
						Usage: flagAutoGitPushUsage,
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
