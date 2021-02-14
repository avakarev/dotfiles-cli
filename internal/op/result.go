package op

import (
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
func (res *result) SourceState() string {
	if res.symlink.Source.Exists() {
		return color.New(color.FgGreen).Sprint("ok")
	}

	if symlink.IsSourceErr(res.err) {
		return color.New(color.FgRed).Sprintf("err: %s", res.err.Error())
	}

	return "?"
}
