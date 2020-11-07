package git

import "github.com/go-git/go-git/v5/plumbing/object"

type Commit struct {
	Hash    string
	Author  Author
	Message string
}

type Author struct {
	Name  string
	Email string
}

func newCommit(commit *object.Commit) Commit {
	return Commit{
		Hash: commit.Hash.String(),
		Author: Author{
			Name:  commit.Author.Name,
			Email: commit.Author.Email,
		},
		Message: commit.Message,
	}
}
