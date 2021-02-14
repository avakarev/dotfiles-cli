package op

import (
	"errors"

	"github.com/avakarev/go-symlink"
	"github.com/fatih/color"
)

// LinkResult represent symlink's link operation result
type LinkResult struct {
	result
}

// Status returns result's status
func (res *LinkResult) Status() string {
	if res.err == nil {
		return color.New(color.FgGreen).Sprint("✔")
	} else if !errors.Is(res.err, symlink.ErrTargetExist) {
		return color.New(color.FgRed).Sprint("✘")
	}
	return " "
}

// TargetState returns symlink's target state
func (res *LinkResult) TargetState() string {
	if res.err == nil {
		return color.New(color.FgGreen).Sprint(res.states.complete)
	}
	if errors.Is(res.err, symlink.ErrTargetExist) {
		return res.states.incomplete
	}
	if symlink.IsTargetErr(res.err) {
		return color.New(color.FgRed).Sprintf("err: %s", res.err.Error())
	}
	return res.states.unknown
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
