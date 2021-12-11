package changelog

import (
	"strings"
	"time"

	"github.com/haunt98/changeloguru/internal/convention"
	"github.com/make-go-great/rst-go"
)

// GenerateRST base on GenerateMarkdown
func GenerateRST(commits []convention.Commit, scopes map[string]struct{}, version string, when time.Time) []rst.Node {
	filteredCommits := filter(commits, scopes)
	if filteredCommits == nil {
		return nil
	}

	addedNodes := convertToListRSTNodes(filteredCommits[addedType])
	fixedNodes := convertToListRSTNodes(filteredCommits[fixedType])
	othersNodes := convertToListRSTNodes(filteredCommits[othersType])

	nodes := make([]rst.Node, 0, len(addedNodes)+len(fixedNodes)+len(othersNodes)+4)

	if len(addedNodes) != 0 {
		nodes = append(nodes, rst.NewSubSection(addedType))
		nodes = append(nodes, addedNodes...)
	}

	if len(fixedNodes) != 0 {
		nodes = append(nodes, rst.NewSubSection(fixedType))
		nodes = append(nodes, fixedNodes...)
	}

	if len(othersNodes) != 0 {
		nodes = append(nodes, rst.NewSubSection(othersType))
		nodes = append(nodes, othersNodes...)
	}

	versionHeader := generateVersionHeaderValue(version, when)
	nodes = append([]rst.Node{
		rst.NewTitle(title),
		rst.NewSection(versionHeader),
	}, nodes...)

	return nodes
}

func ParseRST(data string) []rst.Node {
	lines := strings.Split(data, "\n\n")
	nodes := rst.Parse(lines)

	// Remove title
	if len(nodes) > 0 && rst.Equal(nodes[0], rst.NewTitle(title)) {
		nodes = nodes[1:]
	}

	return nodes
}

func convertToListRSTNodes(commits []convention.Commit) []rst.Node {
	result := make([]rst.Node, 0, len(commits))

	for _, commit := range commits {
		result = append(result, rst.NewListItem(commit.String()))
	}

	return result
}
