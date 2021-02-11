package testutil

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func caller() string {
	_, abs, no, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	rel, _ := filepath.Rel(RootDir(), abs)
	return fmt.Sprintf("\nFailed at %s:%d\n", rel, no)
}
