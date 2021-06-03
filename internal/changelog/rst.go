package changelog

import (
	"fmt"
	"time"

	"github.com/haunt98/changeloguru/internal/convention"
	"github.com/haunt98/clock"
	"github.com/haunt98/rst-go"
)

func GenerateRST(commits []convention.Commit, scopes map[string]struct{}, version string, when time.Time) []rst.Node {
	if len(commits) == 0 {
		return nil
	}

	commitBases := make(map[string][]rst.Node)
	commitBases[addedType] = make([]rst.Node, 0, defaultNodesLen)
	commitBases[fixedType] = make([]rst.Node, 0, defaultNodesLen)
	commitBases[othersType] = make([]rst.Node, 0, defaultNodesLen)

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
			commitBases[addedType] = append(commitBases[addedType], rst.NewListItem(commit.String()))
		case fixedType:
			commitBases[fixedType] = append(commitBases[fixedType], rst.NewListItem(commit.String()))
		case othersType:
			commitBases[othersType] = append(commitBases[othersType], rst.NewListItem(commit.String()))
		default:
			continue
		}
	}

	// Adding each type and header to nodes
	nodes := make([]rst.Node, 0, len(commitBases[addedType])+len(commitBases[fixedType])+len(commitBases[othersType]))

	if len(commitBases[addedType]) != 0 {
		nodes = append(nodes, rst.NewSubSection(addedType))
		nodes = append(nodes, commitBases[addedType]...)
	}

	if len(commitBases[fixedType]) != 0 {
		nodes = append(nodes, rst.NewSubSection(fixedType))
		nodes = append(nodes, commitBases[fixedType]...)
	}

	if len(commitBases[othersType]) != 0 {
		nodes = append(nodes, rst.NewSubSection(othersType))
		nodes = append(nodes, commitBases[othersType]...)
	}

	// Adding title, version to nodes
	versionHeader := fmt.Sprintf("%s (%s)", version, clock.FormatDate(when))
	nodes = append([]rst.Node{
		rst.NewTitle(title),
		rst.NewSection(versionHeader),
	}, nodes...)

	return nodes
}
