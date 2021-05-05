package changelog

import (
	"fmt"
	"strings"
	"time"

	"github.com/haunt98/changeloguru/internal/convention"
	"github.com/haunt98/changeloguru/internal/markdown"
	"github.com/haunt98/clock"
)

const (
	title = "CHANGELOG"

	defaultNodesLen = 10

	firstLevel  = 1
	secondLevel = 2
	thirdLevel  = 3
)

func GenerateMarkdown(commits []convention.Commit, scopes map[string]struct{}, version string, when time.Time) []markdown.Node {
	if len(commits) == 0 {
		return nil
	}

	commitBases := make(map[string][]markdown.Node)
	commitBases[addedType] = make([]markdown.Node, 0, defaultNodesLen)
	commitBases[fixedType] = make([]markdown.Node, 0, defaultNodesLen)
	commitBases[othersType] = make([]markdown.Node, 0, defaultNodesLen)

	for _, commit := range commits {
		// If scopes is empty or commit scope is empty, pass all commits
		if len(scopes) != 0 && commit.Scope != "" {
			// Skip commit outside scopes
			if _, ok := scopes[commit.Scope]; !ok {
				continue
			}
		}

		t := getType(commit.Type)
		switch t {
		case addedType:
			commitBases[addedType] = append(commitBases[addedType], markdown.NewListItem(commit.String()))
		case fixedType:
			commitBases[fixedType] = append(commitBases[fixedType], markdown.NewListItem(commit.String()))
		case othersType:
			commitBases[othersType] = append(commitBases[othersType], markdown.NewListItem(commit.String()))
		default:
			continue
		}
	}

	// Adding each type and header to nodes
	nodes := make([]markdown.Node, 0, len(commitBases[addedType])+len(commitBases[fixedType])+len(commitBases[othersType]))

	if len(commitBases[addedType]) != 0 {
		nodes = append(nodes, markdown.NewHeader(thirdLevel, addedType))
		nodes = append(nodes, commitBases[addedType]...)
	}

	if len(commitBases[fixedType]) != 0 {
		nodes = append(nodes, markdown.NewHeader(thirdLevel, fixedType))
		nodes = append(nodes, commitBases[fixedType]...)
	}

	if len(commitBases[othersType]) != 0 {
		nodes = append(nodes, markdown.NewHeader(thirdLevel, othersType))
		nodes = append(nodes, commitBases[othersType]...)
	}

	// Adding title, version to nodes
	versionHeader := fmt.Sprintf("%s (%s)", version, clock.FormatDate(when))
	nodes = append([]markdown.Node{
		markdown.NewHeader(firstLevel, title),
		markdown.NewHeader(secondLevel, versionHeader),
	}, nodes...)

	return nodes
}

func ParseMarkdown(data string) []markdown.Node {
	lines := strings.Split(data, string(markdown.NewlineToken))
	nodes := markdown.Parse(lines)

	// Remove title
	if len(nodes) > 0 && markdown.Equal(nodes[0], markdown.NewHeader(firstLevel, title)) {
		nodes = nodes[1:]
	}

	return nodes
}
