package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/pkg/diff"
	"github.com/pkg/diff/write"
	"github.com/urfave/cli/v3"

	"github.com/make-go-great/color-go"
	"github.com/make-go-great/ioe-go"
	"github.com/make-go-great/markdown-go"
	"github.com/make-go-great/rst-go"

	"github.com/haunt98/changeloguru/internal/changelog"
	"github.com/haunt98/changeloguru/internal/convention"
	"github.com/haunt98/changeloguru/internal/git"
)

const autoCommitMessageTemplate = "chore(changelog): generate %s"

var (
	ErrUnknownFiletype = errors.New("unknown filetype")
	ErrInvalidVersion  = errors.New("invalid version")
)

func (a *action) RunGenerate(ctx context.Context, c *cli.Command) error {
	a.getFlags(c)

	// Show help if there is no flags
	if c.NumFlags() == 0 {
		return cli.ShowSubcommandHelp(c)
	}

	if a.flags.interactive {
		fmt.Printf("Input version (%s):\n", flagVersionUsage)
		a.flags.version = ioe.ReadInput()

		if a.flags.interactiveFrom {
			fmt.Printf("Input from (%s):\n", flagFromUsage)
			a.flags.from = ioe.ReadInputEmpty()
		}

		if a.flags.interactiveTo {
			fmt.Printf("Input to (%s):\n", flagToUsage)
			a.flags.to = ioe.ReadInputEmpty()
		}
	}

	repo, err := git.NewRepository(a.flags.repository)
	if err != nil {
		return err
	}

	aliasFrom := a.flags.from
	if aliasFrom == "" {
		aliasFrom = "latest"
	}

	aliasTo := a.flags.to
	if aliasTo == "" {
		aliasTo = "earliest"
	}

	color.PrintAppOK(name, fmt.Sprintf("Generate changelog from [%s] to [%s]", aliasFrom, aliasTo))

	ver, err := a.getVersion()
	if err != nil {
		return err
	}

	semVerTags, err := repo.SemVerTags()
	if err != nil {
		return err
	}

	if len(semVerTags) != 0 &&
		ver.LessThanOrEqual(semVerTags[len(semVerTags)-1].Version) {
		return fmt.Errorf("not latest version, expect > %s: %w", semVerTags[len(semVerTags)-1].Version.String(), ErrInvalidVersion)
	}

	commits, err := repo.Log(a.flags.from, a.flags.to)
	if err != nil {
		return err
	}

	conventionalCommits := a.getConventionalCommits(commits)

	finalOutput := a.getFinalOutput()

	if err := a.generateChangelog(conventionalCommits, finalOutput, ver.String()); err != nil {
		return err
	}

	if err := a.doGit(finalOutput, ver.String()); err != nil {
		return err
	}

	return nil
}

func (a *action) getConventionalCommits(commits []git.Commit) []convention.Commit {
	conventionalCommits := make([]convention.Commit, 0, len(commits))
	for _, commit := range commits {
		conventionalCommit, err := convention.NewCommit(commit)
		if err != nil {
			a.log("Failed to new conventional commits %+v: %s", commit, err)
			// Skip bad commit and move on
			continue
		}

		conventionalCommits = append(conventionalCommits, conventionalCommit)
	}

	return conventionalCommits
}

func (a *action) getFinalOutput() string {
	nameWithExt := a.flags.filename + "." + a.flags.filetype
	finalOutput := filepath.Join(a.flags.repository, a.flags.output, nameWithExt)

	a.log("Final output %s", finalOutput)

	return finalOutput
}

func (a *action) getVersion() (*version.Version, error) {
	if a.flags.version == "" {
		return nil, fmt.Errorf("empty version: %w", ErrInvalidVersion)
	}

	if !strings.HasPrefix(a.flags.version, "v") {
		a.flags.version = "v" + a.flags.version
	}

	v, err := version.NewVersion(a.flags.version)
	if err != nil {
		return nil, fmt.Errorf("version: failed to new version: %w", err)
	}

	a.log("Version %s", a.flags.version)

	return v, nil
}

func (a *action) generateChangelog(commits []convention.Commit, finalOutput, ver string) error {
	switch a.flags.filetype {
	case markdownFiletype:
		return a.generateMarkdownChangelog(finalOutput, ver, commits)
	case rstFiletype:
		return a.generateRSTChangelog(finalOutput, ver, commits)
	default:
		return fmt.Errorf("unknown filetype %s: %w", a.flags.filetype, ErrUnknownFiletype)
	}
}

func (a *action) generateMarkdownChangelog(output, ver string, commits []convention.Commit) error {
	// If changelog file already exist, parse markdown from exist file
	var oldNodes []markdown.Node
	bytes, err := os.ReadFile(filepath.Clean(output))
	if err == nil {
		oldNodes = changelog.ParseMarkdown(string(bytes))
	}

	// Generate markdown from commits
	nodes := changelog.GenerateMarkdown(commits, ver, time.Now())

	// Final changelog with new commits above old commits
	nodes = append(nodes, oldNodes...)
	changelogText := markdown.GenerateText(nodes)

	// Demo run
	if a.flags.dryRun {
		if err := diff.Text("old", "new", string(bytes), changelogText, os.Stdout, write.TerminalColor()); err != nil {
			return fmt.Errorf("failed to diff old and new changelog: %w", err)
		}

		return nil
	}

	// Actually writing to changelog file
	if err := os.WriteFile(output, []byte(changelogText), 0o600); err != nil {
		return fmt.Errorf("failed to write file %s: %w", output, err)
	}

	return nil
}

func (a *action) generateRSTChangelog(output, ver string, commits []convention.Commit) error {
	// If changelog file already exist, parse markdown from exist file
	var oldNodes []rst.Node
	bytes, err := os.ReadFile(filepath.Clean(output))
	if err == nil {
		oldNodes = changelog.ParseRST(string(bytes))
	}

	// Generate markdown from commits
	nodes := changelog.GenerateRST(commits, ver, time.Now())

	// Final changelog with new commits above old commits
	nodes = append(nodes, oldNodes...)
	changelogText := rst.GenerateText(nodes)

	// Demo run
	if a.flags.dryRun {
		if err := diff.Text("old", "new", string(bytes), changelogText, os.Stdout, write.TerminalColor()); err != nil {
			return fmt.Errorf("failed to diff old and new changelog: %w", err)
		}

		return nil
	}

	// Actually writing to changelog file
	if err := os.WriteFile(output, []byte(changelogText), 0o600); err != nil {
		return fmt.Errorf("failed to write file %s: %w", output, err)
	}

	return nil
}

func (a *action) doGit(finalOutput, ver string) error {
	if !a.flags.autoGitCommit {
		return nil
	}

	cmdOutput, err := exec.Command("git", "add", finalOutput).CombinedOutput()
	if err != nil {
		return err
	}
	a.log("Git add output:\n%s", cmdOutput)

	commitMsg := fmt.Sprintf(autoCommitMessageTemplate, ver)

	cmdOutput, err = exec.Command("git", "commit", "-m", commitMsg).CombinedOutput()
	if err != nil {
		return err
	}
	a.log("Git commit output:\n%s", cmdOutput)

	if a.flags.autoGitTag {
		cmdOutput, err = exec.Command("git", "tag", ver, "-m", commitMsg).CombinedOutput()
		if err != nil {
			return err
		}
		a.log("Git tag output:\n%s", cmdOutput)
	}

	if a.flags.autoGitPush {
		cmdOutput, err = exec.Command("git", "push").CombinedOutput()
		if err != nil {
			return err
		}
		a.log("Git push output:\n%s", cmdOutput)

		if a.flags.autoGitTag {
			cmdOutput, err = exec.Command("git", "push", "origin", ver).CombinedOutput()
			if err != nil {
				return err
			}
			a.log("Git push tag output:\n%s", cmdOutput)
		}
	}

	return nil
}
