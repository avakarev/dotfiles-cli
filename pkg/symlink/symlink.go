package symlink

import (
	"os"
)

// Symlink represents file of directory symlink
type Symlink struct {
	Source *Source
	Target *Target
	read   bool
}

// IsLinked check whether target linked to the given source
func (l Symlink) IsLinked() bool {
	return l.Target.Exists && l.Target.Link == l.Source.Path
}

func (l *Symlink) Read() error {
	l.read = true
	if err := l.Source.Read(); err != nil {
		return err
	}
	if err := l.Target.Read(); err != nil {
		return err
	}
	return l.Validate()
}

// Validate check whether target linked to the given source
func (l *Symlink) Validate() error {
	if !l.read {
		if err := l.Read(); err != nil {
			return err
		}
	}
	if l.Source.Exists && l.Target.Exists && l.Target.Link != l.Source.Path {
		return ErrTargetMismatch
	}
	return nil
}

// Link creates symlink
func (l *Symlink) Link() error {
	if err := l.Validate(); err != nil {
		return err
	}

	return os.Symlink(l.Source.Path, l.Target.Path)
}

// Unlink deletes symlink (only target, source file/dir stays)
func (l *Symlink) Unlink() error {
	if err := l.Validate(); err != nil {
		return err
	}
	return os.Remove(l.Target.Path)
}

// New returns new Symlink value
func New(s string, t string) Symlink {
	return Symlink{
		Source: NewSource(s),
		Target: NewTarget(t),
	}
}
