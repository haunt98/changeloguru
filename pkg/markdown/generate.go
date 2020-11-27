package markdown

import "strings"

func Generate(bases []Base) string {
	lines := make([]string, len(bases))

	for i, base := range bases {
		lines[i] = base.String()
	}

	return strings.Join(lines, string(newlineToken))
}
