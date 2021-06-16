package changelog

import (
	"testing"

	"github.com/haunt98/changeloguru/internal/convention"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name    string
		commits []convention.Commit
		scopes  map[string]struct{}
		want    map[string][]convention.Commit
	}{
		{
			name: "without scopes",
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
						RawHeader: "build",
						Type:      convention.BuildType,
					},
				},
			},
		},
		{
			name: "with scopes",
			commits: []convention.Commit{
				{
					RawHeader: "feature A",
					Type:      convention.FeatType,
					Scope:     "A",
				},
				{
					RawHeader: "fix B",
					Type:      convention.FixType,
					Scope:     "B",
				},
				{
					RawHeader: "build",
					Type:      convention.BuildType,
				},
			},
			scopes: map[string]struct{}{
				"A": {},
			},
			want: map[string][]convention.Commit{
				addedType: {
					{
						RawHeader: "feature A",
						Type:      convention.FeatType,
						Scope:     "A",
					},
				},
				fixedType: {},
				othersType: {
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
			got := filter(tc.commits, tc.scopes)
			assert.Equal(t, tc.want, got)
		})
	}
}
