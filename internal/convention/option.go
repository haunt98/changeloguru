package convention

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/make-go-great/date-go"

	"github.com/haunt98/changeloguru/internal/git"
)

const (
	leftScope  = "("
	rightScope = ")"
)

var (
	ErrEmptyCommit = errors.New("empty commit")

	// Self build
	headerRegex = regexp.MustCompile(`(?P<type>[a-zA-Z0-9]+)(?P<scope>\([a-zA-Z0-9]+\))?(?P<attention>!)?:\s(?P<description>.+)`)
)

type OptionFn func(*Commit) error

func GetRawHeader(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		if gitCommit.Message == "" {
			return ErrEmptyCommit
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
		c.Scope = strings.TrimLeft(c.Scope, leftScope)
		c.Scope = strings.TrimRight(c.Scope, rightScope)

		return nil
	}
}

func AddAuthorDate(gitCommit git.Commit) OptionFn {
	return func(c *Commit) error {
		c.RawHeader = fmt.Sprintf("%s (%s)", c.RawHeader, date.FormatDateByDefault(gitCommit.Author.When, time.Local))

		return nil
	}
}
