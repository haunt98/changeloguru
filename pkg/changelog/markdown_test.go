package changelog

import (
	"testing"
	"time"

	"github.com/haunt98/changeloguru/pkg/convention"
	"github.com/sebdah/goldie/v2"
)

func TestMarkdownGeneratorGenerate(t *testing.T) {
	tests := []struct {
		name    string
		oldData string
		version string
		t       time.Time
		commits []convention.Commit
		scopes  map[string]struct{}
	}{
		{
			name:    "empty old data",
			version: "v1.0.0",
			t:       time.Date(2020, 1, 18, 0, 0, 0, 0, time.Local),
			commits: []convention.Commit{
				{
					RawHeader: "feat: new feature",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "fix: new fix",
					Type:      convention.FixType,
				},
				{
					RawHeader: "chore: new build",
					Type:      convention.ChoreType,
				},
			},
		},
		{
			name:    "many commits",
			version: "v1.0.0",
			t:       time.Date(2020, 1, 18, 0, 0, 0, 0, time.Local),
			commits: []convention.Commit{
				{
					RawHeader: "feat: new feature",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "feat: support new client",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "fix: new fix",
					Type:      convention.FixType,
				},
				{
					RawHeader: "fix: wrong color",
					Type:      convention.FixType,
				},
				{
					RawHeader: "chore: new build",
					Type:      convention.ChoreType,
				},
				{
					RawHeader: "chore(github): release on github",
					Type:      convention.ChoreType,
				},
				{
					RawHeader: "chore(gitlab): release on gitlab",
					Type:      convention.ChoreType,
				},
				{
					RawHeader: "unleash the dragon",
					Type:      convention.MiscType,
				},
			},
		},
		{
			name:    "ignore commits outside scope",
			version: "v1.0.0",
			t:       time.Date(2020, 3, 22, 0, 0, 0, 0, time.Local),
			commits: []convention.Commit{
				{
					RawHeader: "feat: new feature",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "feat(a): support new client",
					Type:      convention.FeatType,
					Scope:     "a",
				},
				{
					RawHeader: "fix: new fix",
					Type:      convention.FixType,
				},
				{
					RawHeader: "fix(b): wrong color",
					Type:      convention.FixType,
					Scope:     "b",
				},
				{
					RawHeader: "chore(a): new build",
					Type:      convention.ChoreType,
					Scope:     "a",
				},
				{
					RawHeader: "chore(b): new build",
					Type:      convention.ChoreType,
					Scope:     "b",
				},
				{
					RawHeader: "unleash the dragon",
					Type:      convention.MiscType,
				},
			},
			scopes: map[string]struct{}{
				"a": struct{}{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)
			markdownGenerator := NewMarkdownGenerator(tc.oldData, tc.version, tc.t)
			got := markdownGenerator.Generate(tc.commits, tc.scopes)
			g.Assert(t, t.Name(), []byte(got))
		})
	}
}
