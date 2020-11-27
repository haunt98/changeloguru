package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name  string
		bases []Base
		want  string
	}{
		{
			name: "normal",
			bases: []Base{
				header{
					level: 1,
					text:  "header",
				},
				listItem{
					text: "item 1",
				},
				listItem{
					text: "item 2",
				},
			},
			want: "# header\n\n- item 1\n\n- item 2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Generate(tc.bases)
			assert.Equal(t, tc.want, got)
		})
	}
}
