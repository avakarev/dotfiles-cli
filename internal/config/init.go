package config

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// HomeDir is abs path to home directory
var HomeDir string

// WorkingDir is abs path to working directory containing config file
var WorkingDir string

// ConfigDir is abs path to directory containing config file
var ConfigDir string

// ConfigFile is absolute path to config file
var ConfigFile string

// Init uses the following precedence order to find config file.
// Each item takes precedence over the item below it:
//   * `--config` / `-c` flag
//   * `DOTFILES_CONFIG` environment variable
//   * `.dotfilesrc` file in working directory
func Init() {
	HomeDir = os.Getenv("HOME")
	WorkingDir = os.Getenv("PWD")

	if flag := viper.GetString("config"); flag != "" {
		initWith(flag)
		return
	}

	if env := os.Getenv("DOTFILES_CONFIG"); env != "" {
		initWith(env)
		return
	}

	ConfigFile = ".dotfilesrc"
	ConfigDir = WorkingDir
}

func initWith(file string) {
	ConfigFile = strings.NewReplacer(
		"$HOME", HomeDir,
		"~", HomeDir,
	).Replace(file)
	ConfigDir = path.Dir(ConfigFile)
}
