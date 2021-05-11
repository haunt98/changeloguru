package convention

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/haunt98/changeloguru/internal/git"
	"github.com/haunt98/clock"
)

var headerRegex = regexp.MustCompile(`(?P<type>[a-zA-Z]+)(?P<scope>\([a-zA-Z]+\))?(?P<attention>!)?:\s(?P<description>.+)`)

type OptionFn func(*Commit) error

func GetRawHeader(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		// Skip empty commit
		if gitCommit.Message == "" {
			return nil
		}

		message := strings.TrimSpace(gitCommit.Message)
		messages := strings.Split(message, "\n")

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

func AddAuthorDate(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		c.RawHeader = fmt.Sprintf("%s (%s)", c.RawHeader, clock.FormatDate(gitCommit.Author.When))

		return nil
	}
}
