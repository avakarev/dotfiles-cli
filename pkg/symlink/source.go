package symlink

import (
	"errors"
	"os"
)

var (
	// ErrSourceNotExist represents an error when symlink's source doesn't exist
	ErrSourceNotExist = errors.New("source does not exist")
)

// Source represents Symlink's source
type Source struct {
	Path   string
	Exists bool
}

// Read reads the actual file attributes from the file system
func (s *Source) Read() error {
	s.Exists = false

	if _, err := os.Lstat(s.Path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ErrSourceNotExist
		}
		return err
	}

	s.Exists = true
	return nil
}

// NewSource returns new Source value
func NewSource(path string) *Source {
	return &Source{
		Path: path,
	}
}

// IsSourceErr checks whether error is target error
func IsSourceErr(err error) bool {
	return errors.Is(err, ErrSourceNotExist)
}
