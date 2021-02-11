package testutil

import (
	"testing"
)

// NoErr fail the test if `err` is not nil
func NoErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("%sGot unexpected error: %v", caller(), err)
	}
}
