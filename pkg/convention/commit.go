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
)

type Commit struct {
	Type           string
	Scope          string
	Description    string
	BodyAndFooters string
}

func NewCommit(c git.Commit) (result Commit, err error) {
	message := strings.TrimSpace(c.Message)
	messages := strings.Split(message, "\n")
	if len(messages) == 0 {
		err = errors.New("empty commit")
		return
	}

	header := messages[0]
	if err = parseHeader(header, &result); err != nil {
		return
	}

	bodyAndFooters := message[len(header):]
	parseBodyAndFooters(bodyAndFooters, &result)

	return
}

func parseHeader(header string, commit *Commit) error {
	if !headerRegex.MatchString(header) {
		return errors.New("wrong header format")
	}

	headerSubmatches := headerRegex.FindStringSubmatch(header)

	commit.Type = headerSubmatches[1]

	commit.Scope = headerSubmatches[2]
	commit.Scope = strings.TrimPrefix(commit.Scope, "(")
	commit.Scope = strings.TrimSuffix(commit.Scope, ")")

	commit.Description = headerSubmatches[4]

	return nil
}

func parseBodyAndFooters(bodyAndFooters string, commit *Commit) {
	bodyAndFooters = strings.TrimSpace(bodyAndFooters)

	commit.BodyAndFooters = bodyAndFooters
}
