package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
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

	// If user does not specific `flag to`, we automatically choose latest tag
	fallbackLatestTag := false

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
		} else {
			fallbackLatestTag = true
		}
	}

	repo, err := git.NewRepository(a.flags.repository)
	if err != nil {
		return err
	}

	tags, err := repo.SemVerTags()
	if err != nil {
		return err
	}

	ver, err := a.getVersion()
	if err != nil {
		return err
	}

	if len(tags) != 0 {
		if ver.LessThanOrEqual(tags[len(tags)-1].Version) {
			return fmt.Errorf("not latest version, expect > %s: %w", tags[len(tags)-1].Version.String(), ErrInvalidVersion)
		}

		if fallbackLatestTag {
			a.flags.to = tags[len(tags)-1].Version.Original()
		}
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

	commits, err := repo.Log(a.flags.from, a.flags.to)
	if err != nil {
		return err
	}

	conventionalCommits := a.getConventionalCommits(commits)

	finalOutput := a.getFinalOutput()

	if err := a.generateChangelog(conventionalCommits, finalOutput, ver.Original()); err != nil {
		return err
	}

	if err := a.doGit(finalOutput, ver.Original()); err != nil {
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

	// I prefer having prefix `v` in version
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

	if err := a.execCommand([]string{"git", "add", finalOutput}); err != nil {
		return err
	}

	commitMsg := fmt.Sprintf(autoCommitMessageTemplate, ver)
	if err := a.execCommand([]string{"git", "commit", "-m", commitMsg}); err != nil {
		return err
	}

	if a.flags.autoGitTag {
		if err := a.execCommand([]string{"git", "tag", ver, "-m", commitMsg}); err != nil {
			return err
		}
	}

	if a.flags.autoGitPush {
		if err := a.execCommand([]string{"git", "push", "origin"}); err != nil {
			return err
		}

		if a.flags.autoGitTag {
			if err := a.execCommand([]string{"git", "push", "origin", ver}); err != nil {
				return err
			}
		}
	}

	return nil
}

// Wrap with dry run
func (a *action) execCommand(cmds []string) error {
	if a.flags.dryRun {
		log.Printf("%s\n", strings.Join(cmds, " "))
		return nil
	}

	// Safety check
	if len(cmds) < 2 {
		return nil
	}

	cmdOutput, err := exec.Command(cmds[0], cmds[1:]...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("exec: failed to command: %w", err)
	}
	a.log("%s\n%s", strings.Join(cmds, " "), cmdOutput)

	return nil
}
