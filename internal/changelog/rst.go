package changelog

import (
	"strings"
	"time"

	"github.com/make-go-great/rst-go"

	"github.com/haunt98/changeloguru/internal/convention"
)

// GenerateRST base on GenerateMarkdown
func GenerateRST(commits []convention.Commit, version string, when time.Time) []rst.Node {
	filteredCommits := filter(commits)
	if filteredCommits == nil {
		return nil
	}

	filteredNodes := make(map[string][]rst.Node, len(filteredCommits))
	countNodes := 0
	for commitType, commits := range filteredCommits {
		filteredNodes[commitType] = convertToListRSTNodes(commits)
		countNodes += len(commits)
	}

	nodes := make([]rst.Node, 0, countNodes+len(filteredCommits)+1)

	for _, commitType := range changelogTypes {
		if len(filteredNodes[commitType]) != 0 {
			nodes = append(nodes, rst.NewSubSection(commitType))
			nodes = append(nodes, filteredNodes[commitType]...)
		}
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
