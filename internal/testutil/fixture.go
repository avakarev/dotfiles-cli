package testutil

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

// FixturePath returns absolute path to the given fixture
func FixturePath(name string, args ...string) string {
	ext := ""
	if len(args) > 0 {
		ext = args[0]
	}
	return filepath.Join(RootDir(), "test", "fixtures", name+ext)
}

// Fixture returns content of the given fixture
func Fixture(t *testing.T, name string, args ...string) []byte {
	path := FixturePath(name, args...)
	bytes, err := ioutil.ReadFile(path) // #nosec
	if err != err {
		t.Errorf("Failed to read %q fixture: %s", name, err.Error())
	}
	return bytes
}
