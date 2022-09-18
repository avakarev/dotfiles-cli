// Package dotfiles provides access dotfiles groups
package dotfiles

import (
	"fmt"
	"sort"
	"strings"

	"github.com/avakarev/go-symlink"

	"github.com/avakarev/dotfiles-cli/internal/config"
	"github.com/avakarev/dotfiles-cli/internal/pathutil"
)

// Dotfiles represents parsed and loaded dotfiles config
type Dotfiles struct {
	Groups []Group
}

func (d *Dotfiles) validate(names []string) error {
	hash := make(map[string]bool, len(d.Groups))
	groups := make([]string, len(d.Groups))
	for i, g := range d.Groups {
		hash[g.Name] = true
		groups[i] = g.Name
	}
	for _, n := range names {
		if !hash[n] {
			return fmt.Errorf("group %q is not one of %s", n, strings.Join(groups, ", "))
		}
	}
	return nil
}

// Filter returns only groups that match given names
func (d *Dotfiles) Filter(names []string) []Group {
	if err := d.validate(names); err != nil {
		panic(err)
	}

	if len(names) == 0 {
		return d.Groups
	}

	filtered := make([]Group, 0, len(names))
	for _, g := range d.Groups {
		for _, n := range names {
			if g.Name == n {
				filtered = append(filtered, g)
			}
		}
	}
	return filtered
}

// Sort sorts the groups by name
func (d *Dotfiles) Sort() *Dotfiles {
	sort.Slice(d.Groups, func(i, j int) bool {
		return d.Groups[i].Name < d.Groups[j].Name
	})
	return d
}

// Group represents dotfiles group
type Group struct {
	Name     string
	Symlinks []symlink.Symlink
}

// NewGroup returns new Group value
func NewGroup(name string, links []string) (*Group, error) {
	symlinks := make([]symlink.Symlink, 0, len(links))
	for _, link := range links {
		s, t, err := pathutil.Normalize(link)
		if err != nil {
			return nil, err
		}
		symlinks = append(symlinks, symlink.New(s, t))
	}

	return &Group{
		Name:     name,
		Symlinks: symlinks,
	}, nil
}

// New returns new Dotfiles value
func New() (*Dotfiles, error) {
	data, err := config.LoadDefault()
	if err != nil {
		return nil, err
	}

	groups := make([]Group, 0, len(data))
	for name, links := range data {
		g, err := NewGroup(name, links)
		if err != nil {
			return nil, err
		}
		groups = append(groups, *g)
	}

	return &Dotfiles{Groups: groups}, nil
}
