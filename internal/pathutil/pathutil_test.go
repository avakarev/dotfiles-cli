package pathutil_test

import (
	"os"
	"path"
	"testing"

	"github.com/avakarev/dotfiles-cli/internal/config"
	"github.com/avakarev/dotfiles-cli/internal/pathutil"
	"github.com/avakarev/dotfiles-cli/internal/testutil"
)

var (
	homeDir    string
	workingDir string
	configDir  string
	configFile string
)

func setup() {
	homeDir = config.HomeDir
	workingDir = config.WorkingDir
	configDir = config.ConfigDir
	configFile = config.ConfigFile

	config.HomeDir = "/home"
	config.WorkingDir = "/home/foobar"
	config.ConfigDir = "/dotfiles"
	config.ConfigFile = "/dotfiles/rc"
}

func teardown() {
	config.HomeDir = homeDir
	config.WorkingDir = workingDir
	config.ConfigDir = configDir
	config.ConfigFile = configFile
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestAbsWithHomeEnvVar(t *testing.T) {
	want := path.Join(config.HomeDir, ".myrc")
	got := pathutil.Abs("$HOME/.myrc")
	testutil.Diff(want, got, t)
}

func TestAbsWithHomeExpansionSym(t *testing.T) {
	want := path.Join(config.HomeDir, ".myrc")
	got := pathutil.Abs("~/.myrc")
	testutil.Diff(want, got, t)
}

func TestExtractOfSingleEntryWithNoLeadingDot(t *testing.T) {
	source, target, err := pathutil.Extract("myrc")
	testutil.NoErr(err, t)
	testutil.Diff("/dotfiles/myrc", source, t)
	testutil.Diff(path.Join(config.HomeDir, ".myrc"), target, t)
}

func TestExtractOfSingleEntryWithLeadingDot(t *testing.T) {
	source, target, err := pathutil.Extract(".myrc")
	testutil.NoErr(err, t)
	testutil.Diff("/dotfiles/.myrc", source, t)
	testutil.Diff(path.Join(config.HomeDir, ".myrc"), target, t)
}

func TestExtractOfDoubleEntryWithNoLeadingDot(t *testing.T) {
	source, target, err := pathutil.Extract("myrc:~/.myrc")
	testutil.NoErr(err, t)
	testutil.Diff("/dotfiles/myrc", source, t)
	testutil.Diff(path.Join(config.HomeDir, ".myrc"), target, t)
}

func TestExtractOfDoubleEntryWithLeadingDot(t *testing.T) {
	source, target, err := pathutil.Extract(".myrc:$HOME/.myrc")
	testutil.NoErr(err, t)
	testutil.Diff("/dotfiles/.myrc", source, t)
	testutil.Diff(path.Join(config.HomeDir, ".myrc"), target, t)
}

func TestPrettifyOfHomeDir(t *testing.T) {
	want := "~/.myrc"
	got := pathutil.Prettify(path.Join(config.HomeDir, ".myrc"))
	testutil.Diff(want, got, t)
}

func TestPrettifyOfWorkingDir(t *testing.T) {
	want := "./.myrc"
	got := pathutil.Prettify(path.Join(config.WorkingDir, ".myrc"))
	testutil.Diff(want, got, t)
}
