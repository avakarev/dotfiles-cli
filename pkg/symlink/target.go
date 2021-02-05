package symlink

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	// ErrTargetNotExist represents an error when symlink's target doesn't exist
	ErrTargetNotExist = errors.New("target does not exist")

	// ErrTargetNotLink represents an error when symlink's target is supposed to be a link but it's not
	ErrTargetNotLink = errors.New("target is not a link")

	// ErrTargetMismatch represents an error when symlink's actual target is not what is expected
	ErrTargetMismatch = errors.New("target mismatch")
)

// Target represents Symlink's target
type Target struct {
	Path   string
	Exists bool
	Link   string
}

// Read reads the actual file attributes from the file system
func (t *Target) Read() error {
	t.Exists = false
	t.Link = ""

	source, err := filepath.EvalSymlinks(t.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrTargetNotExist
		}
		return err
	}

	t.Exists = true

	if t.Path == source {
		return ErrTargetNotLink
	}

	t.Link = source
	return nil
}

// NewTarget returns new Target value
func NewTarget(path string) *Target {
	return &Target{
		Path: path,
	}
}

// IsTargetErr checks whether error is target error
func IsTargetErr(err error) bool {
	return errors.Is(err, ErrTargetNotExist) ||
		errors.Is(err, ErrTargetNotLink) ||
		errors.Is(err, ErrTargetMismatch)
}
