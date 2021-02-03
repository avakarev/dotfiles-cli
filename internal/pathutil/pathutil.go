package pathutil

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/avakarev/dotfiles-cli/internal/config"
)

func extractTarget(t string) string {
	t = path.Base(t)
	if !strings.HasPrefix(t, ".") {
		t = "." + t
	}
	return filepath.Join(config.HomeDir, t)
}

func extractSource(s string) string {
	return filepath.Join(config.ConfigDir, s)
}

// Abs replaces "$HOME" and "~" substring with home dir value
func Abs(p string) string {
	abs := strings.NewReplacer(
		"$HOME", config.HomeDir,
		"~", config.HomeDir,
	).Replace(p)
	abs, _ = filepath.Abs(abs)
	return abs
}

// Extract returns `source` and `target` paths extracted from the config entry
func Extract(entry string) (string, string, error) {
	if entry == "" {
		return "", "", fmt.Errorf("entry can't be empty")
	}

	parts := strings.Split(entry, ":")
	if len(parts) == 1 {
		return extractSource(parts[0]), extractTarget(parts[0]), nil
	}
	if len(parts) == 2 {
		return extractSource(parts[0]), Abs(parts[1]), nil
	}
	return "", "", fmt.Errorf(entry + " is expected to have \"source\" or \"source:target\" format")
}

// Prettify replaces working dir substring with "." and home dir part with "~"
func Prettify(s string) string {
	r := strings.NewReplacer(
		config.WorkingDir, ".",
		config.HomeDir, "~",
	)

	return r.Replace(s)
}

// Normalize extracts and absolutize paths
func Normalize(s string) (string, string, error) {
	source, target, err := Extract(s)
	if err != nil {
		return "", "", err
	}

	if source, err = filepath.Abs(source); err != nil {
		return "", "", err
	}

	if target, err = filepath.Abs(target); err != nil {
		return "", "", err
	}

	return source, target, nil
}
