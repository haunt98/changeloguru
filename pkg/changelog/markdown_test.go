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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)
			markdownGenerator := NewMarkdownGenerator(tc.oldData, tc.version, tc.t)
			got := markdownGenerator.Generate(tc.commits)
			g.Assert(t, t.Name(), []byte(got))
		})
	}
}
