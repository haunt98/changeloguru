package markdown

import "strings"

//https://guides.github.com/features/mastering-markdown/

const (
	headerToken          = '#'
	defaultListToken     = '-'
	alternativeListToken = '*'
	spaceToken           = ' '
	newlineToken         = '\n'
)

// Base is single markdown syntax representation
// Example: header, list, ...
type Base interface {
	String() string
}

type header struct {
	level int
	text  string
}

func NewHeader(level int, text string) Base {
	return header{
		level: level,
		text:  text,
	}
}

func (h header) String() string {
	var builder strings.Builder

	for i := 0; i < h.level; i++ {
		builder.WriteString(string(headerToken))
	}

	builder.WriteString(string(spaceToken))

	text := strings.TrimSpace(h.text)
	builder.WriteString(text)

	return builder.String()
}

type listItem struct {
	text string
}

func NewListItem(text string) Base {
	return listItem{
		text: text,
	}
}

func (i listItem) String() string {
	text := strings.TrimSpace(i.text)

	return string(defaultListToken) + string(spaceToken) + text
}
