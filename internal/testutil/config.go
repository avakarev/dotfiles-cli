package testutil

import (
	"github.com/avakarev/dotfiles-cli/internal/config"
)

var origHomeDir string
var origWorkingDir string

// MockConfig sets config fields to default mocked values
func MockConfig() {
	config.HomeDir = FixturePath("home")
	config.WorkingDir = FixturePath("home/dotfiles")
}

// ResetConfig sets configuration to the original values
func ResetConfig() {
	config.HomeDir = origHomeDir
	config.WorkingDir = origWorkingDir
}

func init() {
	origHomeDir = config.HomeDir
	origWorkingDir = config.WorkingDir
}
