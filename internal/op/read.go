package op

import (
	"errors"

	"github.com/avakarev/go-symlink"
)

// ReadResult represent symlink's read operation result
type ReadResult struct {
	result
}

// TargetState returns symlink's target state
func (res *ReadResult) TargetState() *State {
	if res.symlink.IsLinked() {
		return NewCompleteState("linked")
	}
	if errors.Is(res.error, symlink.ErrTargetNotExist) {
		return NewIncompleteState("not linked")
	}
	if symlink.IsTargetErr(res.error) {
		return NewErrorState(res.error.Error())
	}
	return NewUnknownState("?")
}

// Read runs read op and return result
func Read(s *symlink.Symlink) Result {
	return &ReadResult{result{
		symlink: s,
		error:   s.Read(),
	}}
}
