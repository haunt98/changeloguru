package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []Node
	}{
		{
			name: "level 1",
			lines: []string{
				"# abc",
				"- xyz",
			},
			want: []Node{
				header{
					level: 1,
					text:  "abc",
				},
				listItem{
					text: "xyz",
				},
			},
		},
		{
			name: "level 3 with alternative",
			lines: []string{
				"### xyz",
				"* abc",
			},
			want: []Node{
				header{
					level: 3,
					text:  "xyz",
				},
				listItem{
					text: "abc",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Parse(tc.lines)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseHeader(t *testing.T) {
	tests := []struct {
		name string
		line string
		want header
	}{
		{
			name: "level 1",
			line: "# abc",
			want: header{
				level: 1,
				text:  "abc",
			},
		},
		{
			name: "level 3",
			line: "### xyz",
			want: header{
				level: 3,
				text:  "xyz",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseHeader(tc.line)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseListItem(t *testing.T) {
	tests := []struct {
		name string
		line string
		want listItem
	}{
		{
			name: "normal",
			line: "- abc",
			want: listItem{
				text: "abc",
			},
		},
		{
			name: "alternative",
			line: "* xyz",
			want: listItem{
				text: "xyz",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseListItem(tc.line)
			assert.Equal(t, tc.want, got)
		})
	}
}
