package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/haunt98/changeloguru/internal/changelog"
	"github.com/haunt98/changeloguru/internal/convention"
	"github.com/haunt98/changeloguru/internal/git"
	"github.com/make-go-great/ioe-go"
	"github.com/make-go-great/markdown-go"
	"github.com/make-go-great/rst-go"
	"github.com/pkg/diff"
	"github.com/pkg/diff/write"
	"github.com/urfave/cli/v2"
	"golang.org/x/mod/semver"
)

func (a *action) RunGenerate(c *cli.Context) error {
	a.getFlags(c)

	if !a.flags.interactive {
		// Show help if there is nothing and not in interactive mode
		if c.NumFlags() == 0 {
			return cli.ShowAppHelp(c)
		}
	} else {
		fmt.Printf("Input version (%s):\n", usageFlagVersion)
		a.flags.version = ioe.ReadInput()

		fmt.Printf("Input from (%s):\n", usageFrom)
		a.flags.from = ioe.ReadInputEmpty()

		fmt.Printf("Input to (%s):\n", usageTo)
		a.flags.to = ioe.ReadInputEmpty()
	}

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
			// Skip bad commit and move on
			continue
		}

		conventionalCommits = append(conventionalCommits, conventionalCommit)
	}

	return conventionalCommits
}

func (a *action) generateChangelog(commits []convention.Commit) error {
	finalOutput := a.getFinalOutput()

	version, err := a.getVersion()
	if err != nil {
		return err
	}

	switch a.flags.filetype {
	case markdownFiletype:
		return a.generateMarkdownChangelog(finalOutput, version, commits)
	case rstFiletype:
		return a.generateRSTChangelog(finalOutput, version, commits)
	default:
		return fmt.Errorf("unknown filetype %s", a.flags.filetype)
	}
}

func (a *action) getFinalOutput() string {
	nameWithExt := a.flags.filename + "." + a.flags.filetype
	finalOutput := filepath.Join(a.flags.output, nameWithExt)

	a.log("final output %s", finalOutput)

	return finalOutput
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
	// If changelog file already exist, parse markdown from exist file
	var oldNodes []markdown.Node
	bytes, err := os.ReadFile(output)
	if err == nil {
		oldNodes = changelog.ParseMarkdown(string(bytes))
	}

	// Generate markdown from commits
	newNodes := changelog.GenerateMarkdown(commits, a.flags.scopes, version, time.Now())

	// Final changelog with new commits above old commits
	nodes := append(newNodes, oldNodes...)
	changelogText := markdown.GenerateText(nodes)

	// Demo run
	if a.flags.dryRun {
		if err := diff.Text("old", "new", string(bytes), changelogText, os.Stdout, write.TerminalColor()); err != nil {
			return fmt.Errorf("failed to diff old and new changelog: %w", err)
		}

		return nil
	}

	// Actually writing to changelog file
	if err := os.WriteFile(output, []byte(changelogText), 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", output, err)
	}

	return nil
}

func (a *action) generateRSTChangelog(output, version string, commits []convention.Commit) error {
	// If changelog file already exist, parse markdown from exist file
	var oldNodes []rst.Node
	bytes, err := os.ReadFile(output)
	if err == nil {
		oldNodes = changelog.ParseRST(string(bytes))
	}

	// Generate markdown from commits
	newNodes := changelog.GenerateRST(commits, a.flags.scopes, version, time.Now())

	// Final changelog with new commits above old commits
	nodes := append(newNodes, oldNodes...)
	changelogText := rst.GenerateText(nodes)

	// Demo run
	if a.flags.dryRun {
		if err := diff.Text("old", "new", string(bytes), changelogText, os.Stdout, write.TerminalColor()); err != nil {
			return fmt.Errorf("failed to diff old and new changelog: %w", err)
		}

		return nil
	}

	// Actually writing to changelog file
	if err := os.WriteFile(output, []byte(changelogText), 0o644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", output, err)
	}

	return nil
}
