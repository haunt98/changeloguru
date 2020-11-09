package convention

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/haunt98/changeloguru/pkg/comparision"
	"github.com/haunt98/changeloguru/pkg/git"
	"github.com/sebdah/goldie/v2"
)

func TestNewCommit(t *testing.T) {
	tests := []struct {
		name    string
		wantErr error
	}{
		{
			name: "Commit message with description and breaking change footer",
		},
		{
			name: "Commit message with not character to draw attention to breaking change",
		},
		{
			name: "Commit message with no body",
		},
		{
			name: "Commit message with scope",
		},
		{
			name: "Commit message with multi-paragraph body and multiple footers",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t)
			gName := g.GoldenFileName(t, t.Name())

			inputName := strings.TrimSuffix(gName, ".golden") + ".txt"
			fmt.Println("inputName", inputName)
			bytes, err := ioutil.ReadFile(inputName)
			if err != nil {
				t.Error(err)
			}

			gotResult, gotErr := NewCommit(git.Commit{
				Message: string(bytes),
			})
			comparision.Diff(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
