package git

import (
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/hashicorp/go-version"
)

// Commit stores all git-commit information
type Commit struct {
	Author  Author
	Message string
}

// Convert from go-git
func newCommit(c *object.Commit) Commit {
	return Commit{
		Message: c.Message,
		Author: Author{
			Name:  c.Author.Name,
			Email: c.Author.Email,
			When:  c.Author.When,
		},
	}
}

type SemVerTag struct {
	Version *version.Version
}
