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
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	verbose := c.Bool(verboseFlag)

	path := currentPath
	if c.NArg() > 0 {
		path = c.Args().Get(0)
	}

	if verbose {
		log.Printf("path %s", path)
	}

	r, err := git.NewRepository(path)
	if err != nil {
		return err
	}

	fromRev := c.String(fromFlag)
	if verbose {
		log.Printf("from revision %s", fromRev)
	}

	toRev := c.String(toFlag)
	if verbose {
		log.Printf("to revision %s", toRev)
	}

	commits, err := r.LogExcludeTo(fromRev, toRev)
	if err != nil {
		return err
	}
	if verbose {
		log.Printf("commits %+v", commits)
	}

	conventionalCommits := make([]convention.Commit, 0, len(commits))
	for _, commit := range commits {
		conventionalCommit, err := convention.NewCommit(commit)
		if err != nil {
			if verbose {
				log.Printf("failed to new conventional commits %+v: %s", commit, err)
			}
			continue
		}

		conventionalCommits = append(conventionalCommits, conventionalCommit)
	}
	if verbose {
		log.Printf("conventional commits %+v", conventionalCommits)
	}

	return nil
}
