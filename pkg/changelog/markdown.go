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

	defaultNodesLen = 10

	firstLevel  = 1
	secondLevel = 2
	thirdLevel  = 3
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
	newBases := g.getNewNodes(commits)
	if len(newBases) == 0 {
		return ""
	}

	nodes := make([]markdown.Node, 0, defaultNodesLen)

	// title
	nodes = append(nodes, markdown.NewHeader(firstLevel, title))

	// version
	nodes = append(nodes, markdown.NewHeader(secondLevel, g.getVersionHeader()))

	// new
	nodes = append(nodes, newBases...)

	// old
	oldNodes := g.getOldNodes()
	nodes = append(nodes, oldNodes...)

	return markdown.Generate(nodes)
}

func (g *MarkdownGenerator) getNewNodes(commits []convention.Commit) []markdown.Node {
	if len(commits) == 0 {
		return nil
	}

	result := make([]markdown.Node, 0, defaultNodesLen)

	commitBases := make(map[string][]markdown.Node)
	commitBases[addedType] = make([]markdown.Node, 0, defaultNodesLen)
	commitBases[fixedType] = make([]markdown.Node, 0, defaultNodesLen)
	commitBases[othersType] = make([]markdown.Node, 0, defaultNodesLen)

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
		result = append(result, markdown.NewHeader(thirdLevel, addedType))
		result = append(result, commitBases[addedType]...)
	}

	if len(commitBases[fixedType]) != 0 {
		result = append(result, markdown.NewHeader(thirdLevel, fixedType))
		result = append(result, commitBases[addedType]...)
	}

	if len(commitBases[othersType]) != 0 {
		result = append(result, markdown.NewHeader(thirdLevel, othersType))
		result = append(result, commitBases[othersType]...)
	}

	return result
}

func (g *MarkdownGenerator) getOldNodes() []markdown.Node {
	if g.oldData == "" {
		return nil
	}

	result := make([]markdown.Node, 0, defaultNodesLen)

	lines := strings.Split(g.oldData, string(markdown.NewlineToken))

	result = append(result, markdown.Parse(lines)...)

	// remove title
	if len(result) > 0 && markdown.Equal(result[0], markdown.NewHeader(firstLevel, title)) {
		result = result[1:]
	}

	return result
}

func (g *MarkdownGenerator) getVersionHeader() string {
	year, month, day := g.t.Date()
	return fmt.Sprintf("%s (%d-%d-%d)", g.version, year, month, day)
}
