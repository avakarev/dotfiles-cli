package op

import (
	"errors"

	"github.com/avakarev/go-symlink"
)

// UnlinkResult represent symlink's unlink operation result
type UnlinkResult struct {
	result
}

// TargetState returns symlink's target state
func (res *UnlinkResult) TargetState() *State {
	if res.error == nil {
		return NewCompleteState("unlinked")
	}
	if errors.Is(res.error, symlink.ErrTargetNotExist) {
		return NewIncompleteState("skipped")
	}
	if symlink.IsTargetErr(res.error) {
		return NewErrorState(res.error.Error())
	}
	return NewUnknownState("?")
}

// Unlink runs link op and return result
func Unlink(s *symlink.Symlink) Result {
	return &UnlinkResult{result{
		symlink: s,
		error:   s.Unlink(),
	}}
}
