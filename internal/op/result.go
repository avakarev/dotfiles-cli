package op

import (
	"errors"

	"github.com/avakarev/go-symlink"
	"github.com/fatih/color"

	"github.com/avakarev/dotfiles-cli/internal/pathutil"
)

// Result represent symlink's operation result
type Result interface {
	Status() string
	TargetPath() string
	TargetState() string
	SourcePath() string
	SourceState() string
}

// TargetStates represents possible target state names after operation
type TargetStates struct {
	complete   string
	incomplete string
	unknown    string
}

// result represent symlink's operation result
type result struct {
	states  TargetStates
	symlink *symlink.Symlink
	err     error
}

// Status returns result's status
func (res *result) Status() string {
	if res.err != nil {
		return " "
	}
	return color.New(color.FgGreen).Sprint("✔")
}

// TargetPath returns symlink's target path
func (res *result) TargetPath() string {
	return pathutil.Prettify(res.symlink.Target.Path())
}

// TargetState returns symlink's target state
func (res *result) TargetState() string {
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

// SourcePath returns symlink's source path
func (res *result) SourcePath() string {
	return pathutil.Prettify(res.symlink.Source.Path())
}

// SourceState returns symlink's source state
func (res *result) SourceState() string {
	if res.symlink.Source.Exists() {
		return color.New(color.FgGreen).Sprint("ok")
	}

	if symlink.IsSourceErr(res.err) {
		return color.New(color.FgRed).Sprintf("err: %s", res.err.Error())
	}

	return "?"
}
