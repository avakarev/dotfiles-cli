package config_test

import (
	"testing"

	"github.com/avakarev/go-testutil"

	"github.com/avakarev/dotfiles-cli/internal/config"
)

func TestLoadWithNoGroups(t *testing.T) {
	content := testutil.FixtureBytes(t, "config_with_no_groups")
	_, err := config.Load(content)
	testutil.MustNoErr(err, t)

	testutil.Diff(map[string][]string{
		"default": {"vim", "vimrc"},
	}, config.MustLoad(content), t)
}

func TestLoadWithGroups(t *testing.T) {
	content := testutil.FixtureBytes(t, "config_with_groups")
	_, err := config.Load(content)
	testutil.MustNoErr(err, t)

	testutil.Diff(map[string][]string{
		"git": {"gitconfig", "gitignore-global", "gitattributes-global"},
		"zsh": {"zsh", "zshrc"},
	}, config.MustLoad(content), t)
}
