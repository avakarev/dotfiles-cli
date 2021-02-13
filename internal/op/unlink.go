package op

import (
	"github.com/avakarev/go-symlink"
)

// UnlinkResult represent symlink's unlink operation result
type UnlinkResult struct {
	result
}

// Unlink runs link op and return result
func Unlink(s *symlink.Symlink) Result {
	return &UnlinkResult{result{
		states: TargetStates{
			complete:   "unlinked",
			incomplete: "skipped",
			unknown:    "?",
		},
		symlink: s,
		err:     s.Unlink(),
	}}
}
