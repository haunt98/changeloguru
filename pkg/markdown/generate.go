package markdown

import "strings"

// GenerateText return single string which represents all markdown nodes
func GenerateText(bases []Node) string {
	lines := make([]string, len(bases))
	for i, base := range bases {
		lines[i] = base.String()
	}

	result := strings.Join(lines, string(NewlineToken)+string(NewlineToken))
	// Fix no newline at end of file
	result += string(NewlineToken)
	return result
}
