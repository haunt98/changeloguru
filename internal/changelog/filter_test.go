package changelog

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/haunt98/changeloguru/internal/convention"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name    string
		commits []convention.Commit
		want    map[string][]convention.Commit
	}{
		{
			name: "all types",
			commits: []convention.Commit{
				{
					RawHeader: "feature A",
					Type:      convention.FeatType,
				},
				{
					RawHeader: "fix B",
					Type:      convention.FixType,
				},
				{
					RawHeader: "build",
					Type:      convention.BuildType,
				},
				{
					RawHeader: "do something other",
					Type:      convention.ChoreType,
				},
			},
			want: map[string][]convention.Commit{
				addedType: {
					{
						RawHeader: "feature A",
						Type:      convention.FeatType,
					},
				},
				fixedType: {
					{
						RawHeader: "fix B",
						Type:      convention.FixType,
					},
				},
				othersType: {
					{
						RawHeader: "do something other",
						Type:      convention.ChoreType,
					},
				},
				buildType: {
					{
						RawHeader: "build",
						Type:      convention.BuildType,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := filter(tc.commits)
			assert.Equal(t, tc.want, got)
		})
	}
}
