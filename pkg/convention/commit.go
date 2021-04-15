package convention

import (
	"github.com/haunt98/changeloguru/pkg/git"
)

// https://www.conventionalcommits.org/en/v1.0.0/
// <type>[optional scope]: <description>
// [optional body]
// [optional footer(s)]

// Commit represens conventional commit
type Commit struct {
	// Commit as is
	RawHeader string

	Type  string
	Scope string
}

// NewCommit return conventional commit from git commit
func NewCommit(c git.Commit) (Commit, error) {
	return NewCommitWithOptions(
		GetRawHeader(c),
		GetTypeAndScope(c),
		AddAuthorDate(c),
	)
}

// NewCommitWithOptions return conventional commit with custom option
func NewCommitWithOptions(opts ...OptionFn) (result Commit, err error) {
	for _, opt := range opts {
		if err = opt(&result); err != nil {
			return
		}
	}

	return
}

func (c *Commit) String() string {
	return c.RawHeader
}
