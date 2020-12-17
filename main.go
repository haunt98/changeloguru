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

	currentPath      = "."
	markdownFiletype = "md"

	defaultPath     = currentPath
	defaultFilename = "CHANGELOG"
	defaultFiletype = markdownFiletype
	defaultVersion  = "0.1.0"

	fromFlag      = "from"
	excludeToFlag = "exclude-to"
	includeToFlag = "include-to"
	versionFlag   = "version"
	filenameFlag  = "filename"
	filetypeFlag  = "filetype"
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
				Usage: "generate from `COMMIT`",
			},
			&cli.StringFlag{
				Name:  excludeToFlag,
				Usage: "generate to `COMMIT` (exclude that commit)",
			},
			&cli.StringFlag{
				Name:  includeToFlag,
				Usage: "generate to `COMMIT` (include that commit)",
			},
			&cli.StringFlag{
				Name:  versionFlag,
				Usage: "generate new `VERSION`",
			},
			&cli.StringFlag{
				Name:  filenameFlag,
				Usage: fmt.Sprintf("output `FILENAME`, default is %s", defaultFilename),
			},
			&cli.StringFlag{
				Name:  filetypeFlag,
				Usage: fmt.Sprintf("output `FILETYPE`, default is %s", defaultFiletype),
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
	a.args[pathArgs] = defaultPath
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
	a.flags[filenameFlag] = c.String(filenameFlag)
	a.flags[filetypeFlag] = c.String(filetypeFlag)
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
	changelogPath, _, filetype := a.getChangelogPath()

	version, err := a.getVersion()
	if err != nil {
		return err
	}

	switch filetype {
	case markdownFiletype:
		return a.generateMarkdownChangelog(changelogPath, version, commits)
	default:
		return fmt.Errorf("unknown filetype %s", filetype)
	}
}

func (a *action) getChangelogPath() (string, string, string) {
	path := a.args[pathArgs]

	filename := a.flags[filenameFlag]
	if filename == "" {
		filename = defaultFilename
	}

	filetype := a.flags[filetypeFlag]
	if filetype == "" {
		filetype = defaultFiletype
	}

	changelogName := filename + "." + filetype
	changelogPath := filepath.Join(path, changelogName)
	a.log("changelog path %s", changelogPath)

	return changelogPath, filename, filetype
}

func (a *action) getVersion() (string, error) {
	version := a.flags[versionFlag]
	if version == "" {
		version = defaultVersion
	}

	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	if !semver.IsValid(version) {
		return "", fmt.Errorf("invalid semver %s", version)
	}

	return version, nil
}

func (a *action) generateMarkdownChangelog(path, version string, commits []convention.Commit) error {
	var oldData string
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		oldData = string(bytes)
	}

	markdownGenerator := changelog.NewMarkdownGenerator(oldData, version, time.Now())
	newData := markdownGenerator.Generate(commits)

	if err := ioutil.WriteFile(path, []byte(newData), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}

	return nil
}

func (a *action) log(format string, v ...interface{}) {
	if a.verbose {
		log.Printf(format, v...)
	}
}
