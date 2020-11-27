package markdown

import (
	"testing"

	"github.com/haunt98/changeloguru/pkg/comparision"
)

func TestHeaderToString(t *testing.T) {
	tests := []struct {
		name   string
		header header
		want   string
	}{
		{
			name: "level 1",
			header: header{
				level: 1,
				text:  "abc",
			},
			want: "# abc",
		},
		{
			name: "level 3",
			header: header{
				level: 3,
				text:  "xyz",
			},
			want: "### xyz",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.header.ToString()
			comparision.Diff(t, tc.want, got)
		})
	}
}

func TestListItemToString(t *testing.T) {
	tests := []struct {
		name     string
		listItem listItem
		want     string
	}{
		{
			name: "normal",
			listItem: listItem{
				text: "abc",
			},
			want: "- abc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.listItem.ToString()
			comparision.Diff(t, tc.want, got)
		})
	}
}
