package op_test

import (
	"os"
	"testing"

	"github.com/avakarev/go-symlink"

	"github.com/avakarev/dotfiles-cli/internal/op"
	"github.com/avakarev/dotfiles-cli/internal/testutil"
)

func setup() {
	testutil.MockConfig()
}

func teardown() {
	testutil.ResetConfig()
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
		"    ~/.rc [?]  →  ./rc.not.exist [err: source does not exist]",
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
		"    ~/.rc.file [err: target is not a link]  →  ./rc [ok]",
		op.Sprint(op.Read(&sym)),
		t,
	)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
