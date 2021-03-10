package buildmeta

import "runtime"

// Build meta information, populated at build-time
var (
	GitCommit string
	BuildDate string
	Version   string
	Compiler  = runtime.Version()
)
