package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/haunt98/changeloguru/pkg/changelog"
	"github.com/haunt98/changeloguru/pkg/convention"
	"github.com/haunt98/changeloguru/pkg/git"
	"github.com/haunt98/changeloguru/pkg/markdown"
	"github.com/pkg/diff"
	"github.com/urfave/cli/v2"
	"golang.org/x/mod/semver"
)

func (a *action) RunGenerate(c *cli.Context) error {
	// Show help if there is nothing
	if c.NumFlags() == 0 {
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
			a.log("failed to new conventional commits %+v: %s", commit, err)
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

	a.log("version %s", a.flags.version)

	return a.flags.version, nil
}

func (a *action) generateMarkdownChangelog(output, version string, commits []convention.Commit) error {
	// If CHANGELOG file already exist
	var oldData string
	bytes, err := os.ReadFile(output)
	if err == nil {
		oldData = string(bytes)
	}

	markdownGenerator := changelog.NewMarkdownGenerator(oldData, version, time.Now())
	newData := markdownGenerator.Generate(commits, a.flags.scopes)

	if a.flags.dryRun {
		oldLines := strings.Split(oldData, string(markdown.NewlineToken))
		newLines := strings.Split(newData, string(markdown.NewlineToken))
		if err := diff.Slices("old", "new", oldLines, newLines, os.Stdout); err != nil {
			return fmt.Errorf("failed to diff old and new data: %w", err)
		}

		return nil
	}

	if err := os.WriteFile(output, []byte(newData), 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", output, err)
	}

	return nil
}
