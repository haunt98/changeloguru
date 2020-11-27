package markdown

import "strings"

const (
	defaultBaseLen = 10
)

func Parse(lines []string) []Node {
	bases := make([]Node, 0, defaultBaseLen)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, string(headerToken)) {
			bases = append(bases, parseHeader(line))
			continue
		}

		if strings.HasPrefix(line, string(defaultListToken)) ||
			strings.HasPrefix(line, string(alternativeListToken)) {
			bases = append(bases, parseListItem(line))
			continue
		}
	}

	return bases
}

func parseHeader(line string) header {
	level := 0

	for _, c := range line {
		if c != headerToken {
			break
		}

		level++
	}

	line = strings.TrimLeft(line, string(headerToken))
	line = strings.TrimSpace(line)

	return header{
		level: level,
		text:  line,
	}
}

func parseListItem(line string) listItem {
	line = strings.TrimLeft(line, string(defaultListToken))
	line = strings.TrimLeft(line, string(alternativeListToken))
	line = strings.TrimSpace(line)

	return listItem{
		text: line,
	}
}
