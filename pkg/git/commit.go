package git

import "github.com/go-git/go-git/v5/plumbing/object"

// Commit stores all git-commit information
type Commit struct {
	Message string
}

// Convert from git-commit
func newCommit(commit *object.Commit) Commit {
	return Commit{
		Message: commit.Message,
	}
}
