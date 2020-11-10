package changelog

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/haunt98/changeloguru/pkg/convention"
)

// https://guides.github.com/features/mastering-markdown/

const (
	markdownTitle = "# CHANGELOG"

	defaultLinesLen = 10
)

type MarkdownGenerator struct {
	path    string
	version string
	t       time.Time
}

func NewMarkdownGenerator(path string, version string, t time.Time) *MarkdownGenerator {
	return &MarkdownGenerator{
		path:    path,
		version: version,
		t:       t,
	}
}

func (g *MarkdownGenerator) Generate(commits []convention.Commit) error {
	lines := g.getLines(commits)
	if len(lines) == 0 {
		return nil
	}

	previousLines := g.getPreviousLines()

	lines = append(lines, previousLines...)

	if err := g.writeLines(lines); err != nil {
		return err
	}

	return nil
}

func (g *MarkdownGenerator) getLines(commits []convention.Commit) []string {
	if len(commits) == 0 {
		return nil
	}

	lines := make([]string, 0, defaultLinesLen)
	lines = append(lines, markdownTitle)
	lines = append(lines, g.composeVersionHeader())

	addedLines := make([]string, 0, defaultLinesLen)
	fixedLines := make([]string, 0, defaultLinesLen)
	othersLines := make([]string, 0, defaultLinesLen)

	for _, commit := range commits {
		t := getType(commit.Type)
		switch t {
		case addedType:
			addedLines = append(addedLines, g.composeListItem(commit.RawHeader))
		case fixedType:
			fixedLines = append(fixedLines, g.composeListItem(commit.RawHeader))
		case othersType:
			othersLines = append(othersLines, g.composeListItem(commit.RawHeader))
		default:
			continue
		}
	}

	if len(addedLines) != 0 {
		lines = append(lines, g.composeTypeHeader(addedType))
		lines = append(lines, addedLines...)
	}

	if len(fixedLines) != 0 {
		lines = append(lines, g.composeTypeHeader(fixedType))
		lines = append(lines, fixedLines...)
	}

	if len(othersLines) != 0 {
		lines = append(lines, g.composeTypeHeader(othersType))
		lines = append(lines, othersLines...)
	}

	return lines
}

func (g *MarkdownGenerator) getPreviousLines() []string {
	prevData, err := ioutil.ReadFile(g.path)
	if err != nil {
		return nil
	}

	prevLines := strings.Split(string(prevData), "\n")
	finalPrevLines := make([]string, 0, len(prevLines))
	for _, prevLine := range prevLines {
		prevLine = strings.TrimSpace(prevLine)
		if prevLine == "" || prevLine == markdownTitle {
			continue
		}

		finalPrevLines = append(finalPrevLines, prevLine)
	}

	return finalPrevLines
}

func (g *MarkdownGenerator) writeLines(lines []string) error {
	data := strings.Join(lines, "\n\n")
	if err := ioutil.WriteFile(g.path, []byte(data), 0644); err != nil {
		return err
	}

	return nil
}

func (g *MarkdownGenerator) composeVersionHeader() string {
	year, month, day := g.t.Date()
	return fmt.Sprintf("## %s (%d-%d-%d)", g.version, year, month, day)
}

func (g *MarkdownGenerator) composeTypeHeader(t string) string {
	return fmt.Sprintf("### %s", t)
}

func (g *MarkdownGenerator) composeListItem(text string) string {
	return fmt.Sprintf("- %s", text)
}
