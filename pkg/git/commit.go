package git

import "github.com/go-git/go-git/v5/plumbing/object"

type Commit struct {
	Hash    string
	Message string
}

func newCommit(commit *object.Commit) Commit {
	return Commit{
		Hash:    commit.Hash.String(),
		Message: commit.Message,
	}
}
