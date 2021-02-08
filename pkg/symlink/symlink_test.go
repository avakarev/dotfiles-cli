package symlink_test

import (
	"testing"

	"github.com/avakarev/dotfiles-cli/internal/testutil"
	"github.com/avakarev/dotfiles-cli/pkg/symlink"
)

func TestNew(t *testing.T) {
	source := testutil.FixturePath("dotfiles/rc")
	target := testutil.FixturePath("home/.rc")
	sl := symlink.New(source, target)

	testutil.Diff(false, sl.Source.Exists, t)
	testutil.Diff(source, sl.Source.Path, t)
	testutil.Diff(false, sl.Target.Exists, t)
	testutil.Diff(target, sl.Target.Path, t)
	testutil.Diff(false, sl.IsLinked(), t)
}

func TestReadOnSucess(t *testing.T) {
	source := testutil.FixturePath("dotfiles/rc")
	target := testutil.FixturePath("home/.rc")
	sl := symlink.New(source, target)

	err := sl.Read()

	testutil.NoErr(err, t)
	testutil.Diff(true, sl.Source.Exists, t)
	testutil.Diff(source, sl.Source.Path, t)
	testutil.Diff(true, sl.Target.Exists, t)
	testutil.Diff(target, sl.Target.Path, t)
	testutil.Diff(true, sl.IsLinked(), t)
}

func TestReadWhenSourceNotExist(t *testing.T) {
	source := testutil.FixturePath("dotfiles/not.exist")
	target := testutil.FixturePath("home/.rc")
	sl := symlink.New(source, target)

	err := sl.Read()
	if err == nil {
		t.Error("Got nil, but expected error")
	}

	testutil.Diff(false, symlink.IsTargetErr(err), t)
	testutil.Diff(true, symlink.IsSourceErr(err), t)
	testutil.Diff("source does not exist", err.Error(), t)
	testutil.Diff(false, sl.IsLinked(), t)
}

func TestReadWhenTargetNotExist(t *testing.T) {
	source := testutil.FixturePath("dotfiles/rc")
	target := testutil.FixturePath("home/.not.exist")
	sl := symlink.New(source, target)

	err := sl.Read()
	if err == nil {
		t.Error("Got nil, but expected error")
	}

	testutil.Diff(false, symlink.IsSourceErr(err), t)
	testutil.Diff(true, symlink.IsTargetErr(err), t)
	testutil.Diff("target does not exist", err.Error(), t)
	testutil.Diff(false, sl.IsLinked(), t)
}

func TestReadWhenTargetNotLink(t *testing.T) {
	source := testutil.FixturePath("dotfiles/rc")
	target := testutil.FixturePath("home/.rc.file")
	sl := symlink.New(source, target)

	err := sl.Read()
	if err == nil {
		t.Error("Got nil, but expected error")
	}

	testutil.Diff(false, symlink.IsSourceErr(err), t)
	testutil.Diff(true, symlink.IsTargetErr(err), t)
	testutil.Diff("target is not a link", err.Error(), t)
	testutil.Diff(false, sl.IsLinked(), t)
}

func TestReadWhenTargetSourceMismatch(t *testing.T) {
	source := testutil.FixturePath("dotfiles/rc")
	target := testutil.FixturePath("home/.rc2")
	sl := symlink.New(source, target)

	err := sl.Read()
	if err == nil {
		t.Error("Got nil, but expected error")
	}

	testutil.Diff(false, symlink.IsSourceErr(err), t)
	testutil.Diff(true, symlink.IsTargetErr(err), t)
	testutil.Diff("target mismatch", err.Error(), t)
	testutil.Diff(false, sl.IsLinked(), t)
}
