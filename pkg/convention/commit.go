package convention

import (
	"errors"
	"regexp"
	"strings"

	"github.com/haunt98/changeloguru/pkg/git"
)

// https://www.conventionalcommits.org/en/v1.0.0/
// <type>[optional scope]: <description>
// [optional body]
// [optional footer(s)]

var (
	headerRegex = regexp.MustCompile(`(?P<type>[a-zA-Z]+)(?P<scope>\([a-zA-Z]+\))?(?P<attention>!)?:\s(?P<description>.+)`)

	ErrEmptyCommit = errors.New("empty commit")
)

// Commit represens conventional commit
type Commit struct {
	// Commit as is
	RawHeader string

	Type  string
	Scope string
}

// NewCommit return conventional commit from git commit
func NewCommit(c git.Commit) (result Commit, err error) {
	// Preprocess
	message := strings.TrimSpace(c.Message)
	messages := strings.Split(message, "\n")
	if len(messages) == 0 {
		err = ErrEmptyCommit
		return
	}

	// Use regex to detect
	result.RawHeader = strings.TrimSpace(messages[0])
	if !headerRegex.MatchString(result.RawHeader) {
		result.Type = MiscType
		return
	}

	headerSubmatches := headerRegex.FindStringSubmatch(result.RawHeader)
	result.Type = strings.ToLower(headerSubmatches[1])
	result.Scope = strings.ToLower(headerSubmatches[2])
	result.Scope = strings.TrimLeft(result.Scope, "(")
	result.Scope = strings.TrimRight(result.Scope, ")")

	return
}

func (c *Commit) String() string {
	return c.RawHeader
}
