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
	appName = "changeloguru"

	currentDir       = "."
	markdownFiletype = "md"

	defaultRepository = currentDir
	defaultOutput     = currentDir
	defaultFilename   = "CHANGELOG"
	defaultFiletype   = markdownFiletype

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
	a := &action{}

	app := &cli.App{
		Name:  appName,
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
				Name:  versionFlag,
				Usage: "`VERSION` to generate, follow Semantic Versioning",
			},
			&cli.StringFlag{
				Name:        repositoryFlag,
				Usage:       "`REPOSITORY` directory path",
				DefaultText: defaultRepository,
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
		fmtErr.Printf("[%s error]: ", appName)
		fmt.Printf("%s\n", err.Error())
	}
}

type action struct {
	flags struct {
		debug      bool
		from       string
		to         string
		version    string
		repository string
		output     string
		filename   string
		filetype   string
	}
}

func (a *action) Run(c *cli.Context) error {
	// Show help if there is nothing
	if c.NArg() == 0 && c.NumFlags() == 0 {
		return cli.ShowAppHelp(c)
	}

	a.getFlags(c)

	commits, err := a.getCommits()
	if err != nil {
		return err
	}

	conventionalCommits := a.getConventionalCommits(commits)

	if err := a.generateChangelog(conventionalCommits); err != nil {
		return err
	}

	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.debug = c.Bool(debugFlag)
	a.flags.from = c.String(fromFlag)
	a.flags.to = c.String(toFlag)
	a.flags.version = c.String(versionFlag)
	a.flags.repository = a.getFlagValue(c, repositoryFlag, defaultRepository)
	a.flags.output = a.getFlagValue(c, outputFlag, defaultOutput)
	a.flags.filename = a.getFlagValue(c, filenameFlag, defaultFilename)
	a.flags.filetype = a.getFlagValue(c, filetypeFlag, defaultFiletype)
}

func (a *action) getFlagValue(c *cli.Context, flag, fallback string) string {
	value := c.String(flag)
	if value == "" {
		value = fallback
	}

	return value
}

func (a *action) getCommits() ([]git.Commit, error) {
	r, err := git.NewRepository(a.flags.repository)
	if err != nil {
		return nil, err
	}

	return r.Log(a.flags.from, a.flags.to)
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
	realOutput := a.getRealOutput()

	version, err := a.getVersion()
	if err != nil {
		return err
	}

	switch a.flags.filetype {
	case markdownFiletype:
		return a.generateMarkdownChangelog(realOutput, version, commits)
	default:
		return fmt.Errorf("unknown filetype %s", a.flags.filetype)
	}
}

func (a *action) getRealOutput() string {
	nameWithExt := a.flags.filename + "." + a.flags.filetype
	realOutput := filepath.Join(a.flags.output, nameWithExt)

	return realOutput
}

func (a *action) getVersion() (string, error) {
	if a.flags.version == "" {
		return "", fmt.Errorf("empty version")
	}

	if !strings.HasPrefix(a.flags.version, "v") {
		a.flags.version = "v" + a.flags.version
	}

	if !semver.IsValid(a.flags.version) {
		return "", fmt.Errorf("invalid semver %s", a.flags.version)
	}

	a.logDebug("version %s", a.flags.version)

	return a.flags.version, nil
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
	if a.flags.debug {
		log.Printf(format, v...)
	}
}
