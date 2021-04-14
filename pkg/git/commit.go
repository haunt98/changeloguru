package git

import (
	"time"

	"github.com/go-git/go-git/v5/plumbing/object"
)

// Commit stores all git-commit information
type Commit struct {
	Message string
	Time    time.Time
}

// Convert from git-commit
func newCommit(commit *object.Commit) Commit {
	return Commit{
		Message: commit.Message,
		Time:    commit.Author.When,
	}
}
