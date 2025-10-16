package changelog

import (
	"github.com/haunt98/changeloguru/internal/convention"
)

func filter(commits []convention.Commit) map[string][]convention.Commit {
	if len(commits) == 0 {
		return nil
	}

	filteredCommits := make(map[string][]convention.Commit, len(changelogTypes))

	for _, commitType := range changelogTypes {
		filteredCommits[commitType] = make([]convention.Commit, 0, defaultLen)
	}

	for _, commit := range commits {
		t := getType(commit)
		filteredCommits[t] = append(filteredCommits[t], commit)
	}

	return filteredCommits
}
