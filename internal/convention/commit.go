package convention

import (
	"github.com/haunt98/changeloguru/internal/git"
)

// https://www.conventionalcommits.org/en/v1.0.0/
// <type>[optional scope]: <description>
// [optional body]
// [optional footer(s)]

// Commit represents conventional commit
type Commit struct {
	RawHeader string // Commit as is
	Type      string
	Scope     string
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
func NewCommitWithOptions(opts ...OptionFn) (Commit, error) {
	result := Commit{}

	for _, opt := range opts {
		if err := opt(&result); err != nil {
			return Commit{}, err
		}
	}

	return result, nil
}

func (c *Commit) String() string {
	return c.RawHeader
}
