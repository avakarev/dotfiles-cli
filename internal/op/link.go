package op

import (
	"github.com/avakarev/go-symlink"
)

// LinkResult represent symlink's link operation result
type LinkResult struct {
	result
}

// Link runs link op and return result
func Link(s *symlink.Symlink) Result {
	return &LinkResult{result{
		states: TargetStates{
			complete:   "linked",
			incomplete: "skipped",
			unknown:    "?",
		},
		symlink: s,
		err:     s.Link(),
	}}
}
