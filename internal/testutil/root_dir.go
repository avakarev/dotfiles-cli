package testutil

import (
	"path/filepath"
	"runtime"
)

// RootDir returns path to project's root directory.
func RootDir() string {
	_, file, _, _ := runtime.Caller(0)
	dir, err := filepath.Abs(filepath.Join(filepath.Dir(file), "..", ".."))
	if err != nil {
		panic(err)
	}
	return dir
}
