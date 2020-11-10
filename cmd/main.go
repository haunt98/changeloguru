package main

import (
	"log"
	"os"

	"github.com/haunt98/changeloguru/pkg/convention"
	"github.com/haunt98/changeloguru/pkg/git"
	"github.com/urfave/cli/v2"
)

const (
	currentPath = "."
	fromFlag    = "from"
	toFlag      = "to"
	verboseFlag = "verbose"
)

func main() {
	a := &action{
		verbose: false,
	}

	app := &cli.App{
		Name:  "changeloguru",
		Usage: "description",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  fromFlag,
				Usage: "from commit revision",
			},
			&cli.StringFlag{
				Name:  toFlag,
				Usage: "to commit revision",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "show what is going on",
			},
		},
		Action: a.Run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type action struct {
	verbose bool
}

func (a *action) Run(c *cli.Context) error {
	a.verbose = c.Bool(verboseFlag)

	path := currentPath
	if c.NArg() > 0 {
		path = c.Args().Get(0)
	}
	a.log("path %s", path)

	commits, err := a.getCommits(c, path)
	if err != nil {
		return err
	}
	a.log("commits %+v", commits)

	conventionalCommits := a.getConventionalCommits(c, commits)
	a.log("conventional commits %+v", conventionalCommits)

	return nil
}

func (a *action) getCommits(c *cli.Context, path string) ([]git.Commit, error) {
	r, err := git.NewRepository(path)
	if err != nil {
		return nil, err
	}

	fromRev := c.String(fromFlag)
	a.log("from revision %s", fromRev)

	toRev := c.String(toFlag)
	a.log("to revision %s", toRev)

	commits, err := r.LogExcludeTo(fromRev, toRev)
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func (a *action) getConventionalCommits(c *cli.Context, commits []git.Commit) []convention.Commit {
	conventionalCommits := make([]convention.Commit, 0, len(commits))

	for _, commit := range commits {
		conventionalCommit, err := convention.NewCommit(commit)
		if err != nil {
			a.log("failed to new conventional commits %+v: %s", commit, err)
			continue
		}

		conventionalCommits = append(conventionalCommits, conventionalCommit)
	}

	return conventionalCommits
}

func (a *action) log(format string, v ...interface{}) {
	if a.verbose {
		log.Printf(format, v...)
	}
}
