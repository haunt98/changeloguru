package changelog

import (
	"testing"
	"time"

	"github.com/haunt98/changeloguru/pkg/convention"
	"github.com/sebdah/goldie/v2"
)

func TestMarkdownGeneratorGetLines(t *testing.T) {
	tests := []struct {
		name    string
		commits []convention.Commit
	}{
		{
			name: "empty",
		},
		{
			name: "1 commit feat",
			commits: []convention.Commit{
				{
					RawHeader:   "feat: description",
					Type:        convention.FeatType,
					Scope:       "",
					Description: "description",
				},
			},
		},
		{
			name: "1 commit fixed",
			commits: []convention.Commit{
				{
					RawHeader:   "fix: description",
					Type:        convention.FixType,
					Scope:       "",
					Description: "description",
				},
			},
		},
		{
			name: "1 commit other",
			commits: []convention.Commit{
				{
					RawHeader:   "ci: description",
					Type:        convention.CiType,
					Scope:       "",
					Description: "description",
				},
			},
		},
		{
			name: "mixed",
			commits: []convention.Commit{
				{
					RawHeader:   "feat: description feat",
					Type:        convention.FeatType,
					Scope:       "",
					Description: "description feat",
				},
				{
					RawHeader:   "fix: description fix",
					Type:        convention.FixType,
					Scope:       "",
					Description: "description fix",
				},
				{
					RawHeader:   "ci: description ci",
					Type:        convention.CiType,
					Scope:       "",
					Description: "description ci",
				},
				{
					RawHeader:   "build: description build",
					Type:        convention.BuildType,
					Scope:       "",
					Description: "description build",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)

			generator := NewMarkdownGenerator("path", "v1.0.0", time.Time{})
			lines := generator.getLines(tc.commits)
			g.AssertJson(t, t.Name(), lines)
		})
	}
}
