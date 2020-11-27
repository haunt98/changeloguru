package changelog

import (
	"fmt"
	"strings"
	"time"

	"github.com/haunt98/changeloguru/pkg/convention"
	"github.com/haunt98/changeloguru/pkg/markdown"
)

const (
	title = "CHANGELOG"

	defaultBasesLen = 10
)

type MarkdownGenerator struct {
	oldData string
	version string
	t       time.Time
}

func NewMarkdownGenerator(oldData, version string, t time.Time) *MarkdownGenerator {
	return &MarkdownGenerator{
		oldData: oldData,
		version: version,
		t:       t,
	}
}

func (g *MarkdownGenerator) Generate(commits []convention.Commit) string {
	newBases := g.getNewMarkdownBases(commits)
	if len(newBases) == 0 {
		return ""
	}

	bases := make([]markdown.Base, 0, defaultBasesLen)

	// title
	bases = append(bases, markdown.NewHeader(1, title))

	// version
	year, month, day := g.t.Date()
	versionHeader := fmt.Sprintf("%s (%d-%d-%d)", g.version, year, month, day)
	bases = append(bases, markdown.NewHeader(2, versionHeader))

	// new
	bases = append(bases, newBases...)

	// old
	oldBases := g.getOldBases()
	bases = append(bases, oldBases...)

	return markdown.Generate(bases)
}

func (g *MarkdownGenerator) getNewMarkdownBases(commits []convention.Commit) []markdown.Base {
	if len(commits) == 0 {
		return nil
	}

	result := make([]markdown.Base, 0, defaultBasesLen)

	commitBases := make(map[string][]markdown.Base)
	commitBases[addedType] = make([]markdown.Base, 0, defaultBasesLen)
	commitBases[fixedType] = make([]markdown.Base, 0, defaultBasesLen)
	commitBases[othersType] = make([]markdown.Base, 0, defaultBasesLen)

	for _, commit := range commits {
		t := getType(commit.Type)
		switch t {
		case addedType:
			commitBases[addedType] = append(commitBases[addedType], markdown.NewListItem(commit.RawHeader))
		case fixedType:
			commitBases[fixedType] = append(commitBases[fixedType], markdown.NewListItem(commit.RawHeader))
		case othersType:
			commitBases[othersType] = append(commitBases[othersType], markdown.NewListItem(commit.RawHeader))
		default:
			continue
		}
	}

	if len(commitBases[addedType]) != 0 {
		result = append(result, markdown.NewHeader(3, addedType))
		result = append(result, commitBases[addedType]...)
	}

	if len(commitBases[fixedType]) != 0 {
		result = append(result, markdown.NewHeader(3, fixedType))
		result = append(result, commitBases[addedType]...)
	}

	if len(commitBases[othersType]) != 0 {
		result = append(result, markdown.NewHeader(3, othersType))
		result = append(result, commitBases[othersType]...)
	}

	return result
}

func (g *MarkdownGenerator) getOldBases() []markdown.Base {
	result := make([]markdown.Base, 0, defaultBasesLen)

	lines := strings.Split(g.oldData, string(markdown.NewlineToken))

	result = append(result, markdown.Parse(lines)...)

	if len(result) > 0 && markdown.Equal(result[0], markdown.NewHeader(1, title)) {
		result = result[1:]
	}

	return result
}
