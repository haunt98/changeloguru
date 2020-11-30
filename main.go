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

	pathArgs = "path"
)

func main() {
	a := &action{
		verbose: false,
		flags:   make(map[string]string),
		args:    make(map[string]string),
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
	flags   map[string]string
	args    map[string]string
}

func (a *action) Run(c *cli.Context) error {
	// set up
	a.getFlags(c)
	a.getArgs(c)

	commits, err := a.getCommits()
	if err != nil {
		return err
	}
	a.log("commits %+v", commits)

	conventionalCommits := a.getConventionalCommits(commits)
	a.log("conventional commits %+v", conventionalCommits)

	if err := a.generateChangelog(conventionalCommits); err != nil {
		return err
	}

	return nil
}

func (a *action) getArgs(c *cli.Context) {
	a.args[pathArgs] = currentPath
	if c.NArg() > 0 {
		a.args[pathArgs] = c.Args().Get(0)
	}
}

func (a *action) getFlags(c *cli.Context) {
	a.verbose = c.Bool(verboseFlag)
	a.flags[fromFlag] = c.String(fromFlag)
	a.flags[excludeToFlag] = c.String(excludeToFlag)
	a.flags[includeToFlag] = c.String(includeToFlag)
	a.flags[versionFlag] = c.String(versionFlag)
}

func (a *action) getCommits() ([]git.Commit, error) {
	path := a.args[pathArgs]
	a.log("path %s", path)

	r, err := git.NewRepository(path)
	if err != nil {
		return nil, err
	}

	fromRev := a.flags[fromFlag]
	a.log("from revision %s", fromRev)

	excludeToRev := a.flags[excludeToFlag]
	a.log("exclude to revision %s", excludeToRev)

	includeToRev := a.flags[includeToFlag]
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

func (a *action) getConventionalCommits(commits []git.Commit) []convention.Commit {
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

func (a *action) generateChangelog(commits []convention.Commit) error {
	path := a.args[pathArgs]
	changelogPath := filepath.Join(path, changelogFile)
	a.log("changelog path %s", path)

	version := a.flags[versionFlag]
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
