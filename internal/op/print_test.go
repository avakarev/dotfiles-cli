package op_test

import (
	"os"
	"testing"

	"github.com/avakarev/go-symlink"
	"github.com/avakarev/go-testutil"

	"github.com/avakarev/dotfiles-cli/internal/config"
	"github.com/avakarev/dotfiles-cli/internal/op"
)

func mockConfig() (reset func()) {
	origHomeDir := config.HomeDir
	origWorkingDir := config.WorkingDir

	config.HomeDir = testutil.FixturePath("home")
	config.WorkingDir = testutil.FixturePath("home/dotfiles")

	return func() {
		config.HomeDir = origHomeDir
		config.WorkingDir = origWorkingDir
	}
}

func TestSprintOnReadResultWithNoErrors(t *testing.T) {
	sym := symlink.New(
		testutil.FixturePath("home/dotfiles/rc"),
		testutil.FixturePath("home/.rc"),
	)
	testutil.Diff(
		"  ✔ ~/.rc [linked]  →  ./rc [ok]",
		op.Sprint(op.Read(&sym)),
		t,
	)
}

func TestSprintOnReadResultWithSourceError(t *testing.T) {
	sym := symlink.New(
		testutil.FixturePath("home/dotfiles/rc.not.exist"),
		testutil.FixturePath("home/.rc"),
	)
	testutil.Diff(
		"  ✘ ~/.rc [?]  →  ./rc.not.exist [err: source does not exist]",
		op.Sprint(op.Read(&sym)),
		t,
	)
}

func TestSprintOnReadResultWithTargetError(t *testing.T) {
	sym := symlink.New(
		testutil.FixturePath("home/dotfiles/rc"),
		testutil.FixturePath("home/.rc.file"),
	)
	testutil.Diff(
		"  ✘ ~/.rc.file [err: target is not a link]  →  ./rc [ok]",
		op.Sprint(op.Read(&sym)),
		t,
	)
}

func TestMain(m *testing.M) {
	resetConfig := mockConfig()
	defer resetConfig()
	os.Exit(m.Run())
}
