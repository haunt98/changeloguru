package convention

import (
	"testing"

	"github.com/haunt98/changeloguru/pkg/comparision"
	"github.com/haunt98/changeloguru/pkg/git"
	"github.com/sebdah/goldie/v2"
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)

			gotResult, gotErr := NewCommit(tc.c)
			comparision.Diff(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
