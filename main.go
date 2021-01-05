package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/haunt98/changeloguru/pkg/changelog"
	"github.com/haunt98/changeloguru/pkg/convention"
	"github.com/haunt98/changeloguru/pkg/git"
	"github.com/urfave/cli/v2"
	"golang.org/x/mod/semver"
)

const (
	name = "changeloguru"

	currentDir       = "."
	markdownFiletype = "md"

	defaultRepositry = currentDir
	defaultOutput    = currentDir
	defaultFilename  = "CHANGELOG"
	defaultFiletype  = markdownFiletype
	defaultVersion   = "0.1.0"

	fromFlag       = "from"
	toFlag         = "to"
	versionFlag    = "version"
	repositoryFlag = "repository"
	outputFlag     = "output"
	filenameFlag   = "filename"
	filetypeFlag   = "filetype"
	debugFlag      = "debug"
)

func main() {
	a := &action{
		flags: make(map[string]string),
		args:  make(map[string]string),
	}

	app := &cli.App{
		Name:  name,
		Usage: "generate changelog from conventional commits",
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
				Name:        versionFlag,
				Usage:       "`VERSION` to generate, follow Semantic Versioning",
				DefaultText: defaultVersion,
			},
			&cli.StringFlag{
				Name:        repositoryFlag,
				Usage:       "`REPOSITORY` directory path",
				DefaultText: defaultRepositry,
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
			&cli.BoolFlag{
				Name:    debugFlag,
				Aliases: []string{"d"},
				Usage:   "show debugging info",
			},
		},
		Action: a.Run,
	}

	if err := app.Run(os.Args); err != nil {
		// Highlight error
		fmtErr := color.New(color.FgRed)
		fmtErr.Printf("[%s error]: ", name)
		fmt.Printf("%s\n", err.Error())
	}
}

type action struct {
	debug bool
	flags map[string]string
	args  map[string]string
}

func (a *action) Run(c *cli.Context) error {
	// Show help if there is nothing
	if c.NArg() == 0 && c.NumFlags() == 0 {
		return cli.ShowAppHelp(c)
	}

	// Set up
	a.getFlags(c)

	commits, err := a.getCommits()
	if err != nil {
		return err
	}
	a.logDebug("commits %+v", commits)

	conventionalCommits := a.getConventionalCommits(commits)
	a.logDebug("conventional commits %+v", conventionalCommits)

	if err := a.generateChangelog(conventionalCommits); err != nil {
		return err
	}

	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.debug = c.Bool(debugFlag)
	a.flags[fromFlag] = c.String(fromFlag)
	a.flags[toFlag] = c.String(toFlag)
	a.flags[versionFlag] = a.getFlagValue(c, versionFlag, defaultVersion)
	a.flags[repositoryFlag] = a.getFlagValue(c, repositoryFlag, defaultRepositry)
	a.flags[outputFlag] = a.getFlagValue(c, outputFlag, defaultOutput)
	a.flags[filenameFlag] = a.getFlagValue(c, filenameFlag, defaultFilename)
	a.flags[filetypeFlag] = a.getFlagValue(c, filetypeFlag, defaultFiletype)
}

func (a *action) getFlagValue(c *cli.Context, flag, fallback string) string {
	value := c.String(flag)
	if value == "" {
		value = fallback
	}

	return value
}

func (a *action) getCommits() ([]git.Commit, error) {
	repository := a.flags[repositoryFlag]
	a.logDebug("repository %s", repository)

	r, err := git.NewRepository(repository)
	if err != nil {
		return nil, err
	}

	fromRev := a.flags[fromFlag]
	a.logDebug("from revision %s", fromRev)

	toRev := a.flags[toFlag]
	a.logDebug("to revision %s", toRev)

	return r.Log(fromRev, toRev)
}

func (a *action) getConventionalCommits(commits []git.Commit) []convention.Commit {
	conventionalCommits := make([]convention.Commit, 0, len(commits))
	for _, commit := range commits {
		conventionalCommit, err := convention.NewCommit(commit)
		if err != nil {
			a.logDebug("failed to new conventional commits %+v: %s", commit, err)
			continue
		}

		conventionalCommits = append(conventionalCommits, conventionalCommit)
	}

	return conventionalCommits
}

func (a *action) generateChangelog(commits []convention.Commit) error {
	realOutput, filetype := a.getRealOutput()

	version, err := a.getVersion()
	if err != nil {
		return err
	}

	switch filetype {
	case markdownFiletype:
		return a.generateMarkdownChangelog(realOutput, version, commits)
	default:
		return fmt.Errorf("unknown filetype %s", filetype)
	}
}

func (a *action) getRealOutput() (string, string) {
	output := a.flags[outputFlag]
	filename := a.flags[filenameFlag]
	filetype := a.flags[filetypeFlag]

	nameWithExt := filename + "." + filetype
	realOutput := filepath.Join(output, nameWithExt)
	a.logDebug("output path %s", realOutput)

	return realOutput, filetype
}

func (a *action) getVersion() (string, error) {
	version := a.flags[versionFlag]
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	if !semver.IsValid(version) {
		return "", fmt.Errorf("invalid semver %s", version)
	}

	a.logDebug("version %s", version)

	return version, nil
}

func (a *action) generateMarkdownChangelog(output, version string, commits []convention.Commit) error {
	// If CHANGELOG file already exist
	var oldData string
	bytes, err := ioutil.ReadFile(output)
	if err == nil {
		oldData = string(bytes)
	}

	markdownGenerator := changelog.NewMarkdownGenerator(oldData, version, time.Now())
	newData := markdownGenerator.Generate(commits)

	if err := ioutil.WriteFile(output, []byte(newData), 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", output, err)
	}

	return nil
}

func (a *action) logDebug(format string, v ...interface{}) {
	if a.debug {
		log.Printf(format, v...)
	}
}
