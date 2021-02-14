package op

import (
	"errors"

	"github.com/avakarev/go-symlink"
)

// LinkResult represent symlink's link operation result
type LinkResult struct {
	result
}

// TargetState returns symlink's target state
func (res *LinkResult) TargetState() *State {
	if res.error == nil {
		return NewCompleteState("linked")
	}
	if errors.Is(res.error, symlink.ErrTargetExist) {
		return NewIncompleteState("skipped")
	}
	if symlink.IsTargetErr(res.error) {
		return NewErrorState(res.error.Error())
	}
	return NewUnknownState("?")
}

// Link runs link op and return result
func Link(s *symlink.Symlink) Result {
	return &LinkResult{result{
		symlink: s,
		error:   s.Link(),
	}}
}
