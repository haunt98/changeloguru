package changelog

import (
	"fmt"
	"time"

	"github.com/haunt98/changeloguru/pkg/convention"
)

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

	lines = append(lines, g.composeTypeHeader(addedType))
	lines = append(lines, addedLines...)

	lines = append(lines, g.composeTypeHeader(fixedType))
	lines = append(lines, fixedLines...)

	lines = append(lines, g.composeTypeHeader(othersType))
	lines = append(lines, othersLines...)

	return lines
}

func (g *MarkdownGenerator) composeVersionHeader() string {
	year, month, day := g.t.Date()
	return fmt.Sprintf("## %s (%d-%d-%d)", g.version, year, month, day)
}

func (g *MarkdownGenerator) composeTypeHeader(t string) string {
	return fmt.Sprintf("### %s", t)
}

func (g *MarkdownGenerator) composeListItem(text string) string {
	return fmt.Sprintf("* %s", text)
}
