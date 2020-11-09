package git

import "github.com/go-git/go-git/v5/plumbing/object"

type Commit struct {
	Message string
}

func newCommit(commit *object.Commit) Commit {
	return Commit{
		Message: commit.Message,
	}
}
