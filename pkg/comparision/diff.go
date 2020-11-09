package comparision

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Diff(t *testing.T, want, got interface{}) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
