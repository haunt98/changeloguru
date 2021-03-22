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

var headerRegex = regexp.MustCompile(`(?P<type>[a-zA-Z]+)(?P<scope>\([a-zA-Z]+\))?(?P<attention>!)?:\s(?P<description>.+)`)

type Commit struct {
	// Commit as is
	RawHeader string

	Type  string
	Scope string
}

func NewCommit(c git.Commit) (result Commit, err error) {
	message := strings.TrimSpace(c.Message)
	messages := strings.Split(message, "\n")
	if len(messages) == 0 {
		err = errors.New("empty commit")
		return
	}

	result.RawHeader = strings.TrimSpace(messages[0])
	result.UpdateType()

	return
}

func (c *Commit) String() string {
	return c.RawHeader
}

func (c *Commit) UpdateType() {
	if !headerRegex.MatchString(c.RawHeader) {
		c.Type = MiscType
		return
	}

	headerSubmatches := headerRegex.FindStringSubmatch(c.RawHeader)
	c.Type = strings.ToLower(headerSubmatches[1])
	c.Scope = strings.ToLower(headerSubmatches[2])
}
