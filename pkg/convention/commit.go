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
	GitCommit git.Commit
	RawHeader string
	Type      string
}

func NewCommit(c git.Commit) (result Commit, err error) {
	// For directly use git commit
	result.GitCommit = c

	message := strings.TrimSpace(c.Message)
	messages := strings.Split(message, "\n")
	if len(messages) == 0 {
		err = errors.New("empty commit")
		return
	}

	header := strings.TrimSpace(messages[0])
	if err = parseHeader(header, &result); err != nil {
		return
	}

	return
}

func parseHeader(header string, commit *Commit) error {
	if !headerRegex.MatchString(header) {
		return errors.New("wrong header format")
	}

	commit.RawHeader = header

	headerSubmatches := headerRegex.FindStringSubmatch(header)

	commit.Type = strings.ToLower(headerSubmatches[1])

	return nil
}
