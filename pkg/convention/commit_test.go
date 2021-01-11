package convention

import (
	"testing"

	"github.com/haunt98/changeloguru/pkg/git"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewCommit(t *testing.T) {
	tests := []struct {
		name    string
		c       git.Commit
		wantErr error
	}{
		{
			name: "Commit message with not character to draw attention to breaking change",
			c: git.Commit{
				Message: "refactor!: drop support for Node 6",
			},
		},
		{
			name: "Commit message with no body",
			c: git.Commit{
				Message: "docs: correct spelling of CHANGELOG",
			},
		},
		{
			name: "Commit message with scope",
			c: git.Commit{
				Message: "feat(lang): add polish language",
			},
		},
		{
			name: "Uppercase",
			c: git.Commit{
				Message: "REFACTOR!: drop support for Node 6",
			},
		},
		{
			name: "Mixedcase",
			c: git.Commit{
				Message: "Docs: correct spelling of CHANGELOG",
			},
		},
		{
			name: "Misc",
			c: git.Commit{
				Message: "random git message",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)

			gotResult, gotErr := NewCommit(tc.c)
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
