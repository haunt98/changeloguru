package changelog

import (
	"github.com/haunt98/changeloguru/internal/convention"
)

func filter(commits []convention.Commit, scopes map[string]struct{}) map[string][]convention.Commit {
	if len(commits) == 0 {
		return nil
	}

	filteredCommits := make(map[string][]convention.Commit)

	for _, commitType := range changelogTypes {
		filteredCommits[commitType] = make([]convention.Commit, 0, defaultLen)
	}

	for _, commit := range commits {
		// If scopes is empty or commit scope is empty, pass all commits
		if len(scopes) != 0 && commit.Scope != "" {
			// Skip commit outside scopes
			if _, ok := scopes[commit.Scope]; !ok {
				continue
			}
		}

		t := getType(commit)
		filteredCommits[t] = append(filteredCommits[t], commit)
	}

	return filteredCommits
}
