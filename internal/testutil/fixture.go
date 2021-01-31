package testutil

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

// FixturePath returns absolute path to the given fixture
func FixturePath(name string, args ...string) string {
	fixturesDir := filepath.Join(RootDir(), "test", "fixtures")
	ext := ""
	if len(args) > 0 {
		ext = args[0]
	}
	return filepath.Join(fixturesDir, strings.ReplaceAll(name, "/", "_")+ext)
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
