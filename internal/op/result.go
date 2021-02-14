package op

import (
	"github.com/avakarev/go-symlink"

	"github.com/avakarev/dotfiles-cli/internal/pathutil"
)

// Result represent symlink's operation result
type Result interface {
	TargetPath() string
	TargetState() *State
	SourcePath() string
	SourceState() *State
}

// result represent symlink's operation result
type result struct {
	symlink *symlink.Symlink
	error   error
}

// TargetPath returns symlink's target path
func (res *result) TargetPath() string {
	return pathutil.Prettify(res.symlink.Target.Path())
}

// TargetState returns symlink's target state
func (res *result) TargetState() string {
	panic("not implemented")
}

// SourcePath returns symlink's source path
func (res *result) SourcePath() string {
	return pathutil.Prettify(res.symlink.Source.Path())
}

// SourceState returns symlink's source state
func (res *result) SourceState() *State {
	if res.symlink.Source.Exists() {
		return NewCompleteState("ok")
	}

	if symlink.IsSourceErr(res.error) {
		return NewErrorState(res.error.Error())
	}

	return NewUnknownState("?")
}
