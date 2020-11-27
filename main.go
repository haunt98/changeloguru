package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/haunt98/changeloguru/pkg/changelog"
	"github.com/haunt98/changeloguru/pkg/convention"
	"github.com/haunt98/changeloguru/pkg/git"
	"github.com/urfave/cli/v2"
	"golang.org/x/mod/semver"
)

const (
	name        = "changeloguru"
	description = "generate changelog from conventional commits"

	currentPath   = "."
	changelogFile = "CHANGELOG.md"

	fromFlag      = "from"
	excludeToFlag = "exclude-to"
	includeToFlag = "include-to"
	versionFlag   = "version"
	verboseFlag   = "verbose"
)

func main() {
	a := &action{
		verbose: false,
	}

	app := &cli.App{
		Name:  name,
		Usage: description,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  fromFlag,
				Usage: "from commit revision",
			},
			&cli.StringFlag{
				Name:  excludeToFlag,
				Usage: "to commit revision (exclude)",
			},
			&cli.StringFlag{
				Name:  includeToFlag,
				Usage: "to commit revision (include)",
			},
			&cli.StringFlag{
				Name:  versionFlag,
				Usage: "version",
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

	if err := a.generateChangelog(c, path, conventionalCommits); err != nil {
		return err
	}

	return nil
}

func (a *action) getCommits(c *cli.Context, path string) ([]git.Commit, error) {
	r, err := git.NewRepository(path)
	if err != nil {
		return nil, err
	}

	fromRev := c.String(fromFlag)
	a.log("from revision %s", fromRev)

	excludeToRev := c.String(excludeToFlag)
	a.log("exclude to revision %s", excludeToRev)

	includeToRev := c.String(includeToFlag)
	a.log("include to revision %s", includeToRev)

	if excludeToRev != "" && includeToRev != "" {
		return nil, errors.New("excludeToFlag and includeToFlag can not appear same time")
	}

	if excludeToRev != "" {
		return r.LogExcludeTo(fromRev, excludeToRev)
	}

	if includeToRev != "" {
		return r.LogIncludeTo(fromRev, includeToRev)
	}

	return r.Log(fromRev)
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

func (a *action) generateChangelog(c *cli.Context, path string, commits []convention.Commit) error {
	changelogPath := filepath.Join(path, changelogFile)

	version := c.String(versionFlag)
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}
	if !semver.IsValid(version) {
		return fmt.Errorf("invalid semver %s", version)
	}

	var oldData string
	bytes, err := ioutil.ReadFile(changelogPath)
	if err == nil {
		oldData = string(bytes)
	}

	markdownGenerator := changelog.NewMarkdownGenerator(oldData, version, time.Now())
	newData := markdownGenerator.Generate(commits)

	if err := ioutil.WriteFile(changelogPath, []byte(newData), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", changelogPath, err)
	}

	return nil
}

func (a *action) log(format string, v ...interface{}) {
	if a.verbose {
		log.Printf(format, v...)
	}
}
