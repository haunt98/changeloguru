package changelog

import (
	"strings"
	"time"

	"github.com/make-go-great/markdown-go"

	"github.com/haunt98/changeloguru/internal/convention"
)

const (
	firstLevel  = 1
	secondLevel = 2
	thirdLevel  = 3
)

func GenerateMarkdown(commits []convention.Commit, scopes map[string]struct{}, version string, when time.Time) []markdown.Node {
	filteredCommits := filter(commits, scopes)
	if filteredCommits == nil {
		return nil
	}

	addedNodes := convertToListMarkdownNodes(filteredCommits[addedType])
	fixedNodes := convertToListMarkdownNodes(filteredCommits[fixedType])
	othersNodes := convertToListMarkdownNodes(filteredCommits[othersType])

	// 4 = 3 type header + 1 version header
	nodes := make([]markdown.Node, 0, len(addedNodes)+len(fixedNodes)+len(othersNodes)+4)

	// Adding each type

	if len(addedNodes) != 0 {
		nodes = append(nodes, markdown.NewHeader(thirdLevel, addedType))
		nodes = append(nodes, addedNodes...)
	}

	if len(fixedNodes) != 0 {
		nodes = append(nodes, markdown.NewHeader(thirdLevel, fixedType))
		nodes = append(nodes, fixedNodes...)
	}

	if len(othersNodes) != 0 {
		nodes = append(nodes, markdown.NewHeader(thirdLevel, othersType))
		nodes = append(nodes, othersNodes...)
	}

	// Adding title
	versionHeader := generateVersionHeaderValue(version, when)
	nodes = append([]markdown.Node{
		markdown.NewHeader(firstLevel, title),
		markdown.NewHeader(secondLevel, versionHeader),
	}, nodes...)

	return nodes
}

func ParseMarkdown(data string) []markdown.Node {
	lines := strings.Split(data, "\n\n")
	nodes := markdown.Parse(lines)

	// Remove title
	if len(nodes) > 0 && markdown.Equal(nodes[0], markdown.NewHeader(firstLevel, title)) {
		nodes = nodes[1:]
	}

	return nodes
}

func convertToListMarkdownNodes(commits []convention.Commit) []markdown.Node {
	result := make([]markdown.Node, 0, len(commits))

	for _, commit := range commits {
		result = append(result, markdown.NewListItem(commit.String()))
	}

	return result
}
