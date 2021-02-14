package op

import (
	"errors"

	"github.com/avakarev/go-symlink"
	"github.com/fatih/color"
)

// ReadResult represent symlink's read operation result
type ReadResult struct {
	result
}

// Status returns result's status
func (res *ReadResult) Status() string {
	if res.symlink.IsLinked() {
		return color.New(color.FgGreen).Sprint("✔")
	} else if !errors.Is(res.err, symlink.ErrTargetNotExist) {
		return color.New(color.FgRed).Sprint("✘")
	}
	return " "
}

// TargetState returns symlink's target state
func (res *ReadResult) TargetState() string {
	if res.symlink.IsLinked() {
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

// Read runs read op and return result
func Read(s *symlink.Symlink) Result {
	return &ReadResult{result{
		states: TargetStates{
			complete:   "linked",
			incomplete: "not linked",
			unknown:    "?",
		},
		symlink: s,
		err:     s.Read(),
	}}
}
