package changelog

import (
	"testing"
	"time"

	"github.com/sebdah/goldie/v2"

	"github.com/make-go-great/markdown-go"

	"github.com/haunt98/changeloguru/internal/convention"
)

func TestGenerateMarkdown(t *testing.T) {
	tests := []struct {
		name    string
		commits []convention.Commit
		scopes  map[string]struct{}
		version string
		when    time.Time
	}{
		{
			name: "empty old data",
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
			version: "v1.0.0",
			when:    time.Date(2020, 1, 18, 0, 0, 0, 0, time.Local),
		},
		{
			name: "many commits",
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
			version: "v1.0.0",
			when:    time.Date(2020, 1, 18, 0, 0, 0, 0, time.Local),
		},
		{
			name: "ignore commits outside scope",
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
				"a": {},
			},
			version: "v1.0.0",
			when:    time.Date(2020, 3, 22, 0, 0, 0, 0, time.Local),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)
			got := GenerateMarkdown(tc.commits, tc.scopes, tc.version, tc.when)
			g.Assert(t, t.Name(), []byte(markdown.GenerateText(got)))
		})
	}
}
