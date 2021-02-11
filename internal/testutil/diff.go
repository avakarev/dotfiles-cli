package testutil

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Diff fail the test if `want` differs from `got`, and prints human-readable error
func Diff(want interface{}, got interface{}, t *testing.T) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%sGot unexpected result (-want +got):\n%s", caller(), diff)
	}
}
