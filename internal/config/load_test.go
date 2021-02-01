package config_test

import (
	"testing"

	"github.com/avakarev/dotfiles-cli/internal/config"
	"github.com/avakarev/dotfiles-cli/internal/testutil"
)

func TestLoadWithNoGroups(t *testing.T) {
	content := testutil.Fixture(t, "config_with_no_groups")
	config.Load(content)

	testutil.Diff(map[string][]string{
		"default": {"vim", "vimrc"},
	}, config.MustLoad(content), t)
}

func TestLoadWithGroups(t *testing.T) {
	content := testutil.Fixture(t, "config_with_groups")
	config.Load(content)

	testutil.Diff(map[string][]string{
		"git": {"gitconfig", "gitignore-global", "gitattributes-global"},
		"zsh": {"zsh", "zshrc"},
	}, config.MustLoad(content), t)
}
