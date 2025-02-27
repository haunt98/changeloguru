package changelog

import (
	"github.com/haunt98/changeloguru/internal/convention"
)

func filter(commits []convention.Commit, scopes map[string]struct{}) map[string][]convention.Commit {
	if len(commits) == 0 {
		return nil
	}

	filteredCommits := make(map[string][]convention.Commit)
	filteredCommits[addedType] = make([]convention.Commit, 0, defaultLen)
	filteredCommits[fixedType] = make([]convention.Commit, 0, defaultLen)
	filteredCommits[othersType] = make([]convention.Commit, 0, defaultLen)
	filteredCommits[buildType] = make([]convention.Commit, 0, defaultLen)

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
			filteredCommits[addedType] = append(filteredCommits[addedType], commit)
		case fixedType:
			filteredCommits[fixedType] = append(filteredCommits[fixedType], commit)
		case othersType:
			filteredCommits[othersType] = append(filteredCommits[othersType], commit)
		case buildType:
			filteredCommits[buildType] = append(filteredCommits[buildType], commit)
		default:
			continue
		}
	}

	return filteredCommits
}
