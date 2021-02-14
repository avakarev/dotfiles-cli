package op

import (
	"errors"

	"github.com/avakarev/go-symlink"
	"github.com/fatih/color"
)

// UnlinkResult represent symlink's unlink operation result
type UnlinkResult struct {
	result
}

// Status returns result's status
func (res *UnlinkResult) Status() string {
	if res.err == nil {
		return color.New(color.FgGreen).Sprint("✔")
	} else if !errors.Is(res.err, symlink.ErrTargetNotExist) {
		return color.New(color.FgRed).Sprint("✘")
	}
	return " "
}

// TargetState returns symlink's target state
func (res *UnlinkResult) TargetState() string {
	if res.err == nil {
		return color.New(color.FgGreen).Sprint(res.states.complete)
	}
	if errors.Is(res.err, symlink.ErrTargetNotExist) {
		return res.states.incomplete
	}
	if symlink.IsTargetErr(res.err) {
		return color.New(color.FgRed).Sprintf("err: %s", res.err.Error())
	}
	return res.states.unknown
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
