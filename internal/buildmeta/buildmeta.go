package buildmeta

import "runtime"

// Build meta information, populated at build-time
var (
	Version  string
	Date     string
	Commit   string
	Compiler = runtime.Version()
)
