package convention

import (
	"errors"
	"regexp"
	"strings"

	"github.com/haunt98/changeloguru/pkg/git"
)

var (
	ErrEmptyCommit = errors.New("empty commit")

	headerRegex = regexp.MustCompile(`(?P<type>[a-zA-Z]+)(?P<scope>\([a-zA-Z]+\))?(?P<attention>!)?:\s(?P<description>.+)`)
)

type OptionFn func(*Commit) error

func GetRawHeader(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		message := strings.TrimSpace(gitCommit.Message)
		messages := strings.Split(message, "\n")
		if len(messages) == 0 {
			return ErrEmptyCommit
		}

		c.RawHeader = messages[0]

		return nil
	}
}

func GetTypeAndScope(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		if !headerRegex.MatchString(c.RawHeader) {
			c.Type = MiscType

			return nil
		}

		headerSubmatches := headerRegex.FindStringSubmatch(c.RawHeader)
		c.Type = strings.ToLower(headerSubmatches[1])
		c.Scope = strings.ToLower(headerSubmatches[2])
		c.Scope = strings.TrimLeft(c.Scope, "(")
		c.Scope = strings.TrimRight(c.Scope, ")")

		return nil
	}
}
