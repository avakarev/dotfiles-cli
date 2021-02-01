package printer_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/avakarev/dotfiles-cli/internal/printer"
	"github.com/avakarev/dotfiles-cli/internal/testutil"
)

func TestDefault(t *testing.T) {
	buf := &bytes.Buffer{}
	p := printer.Default(buf)

	p.Addln("hello", "world")
	p.Addln("foo", "bar")
	p.Addf("%d\n", 42)
	p.Flush()

	testutil.Diff(strings.Join([]string{
		"hello  world  \n",
		"foo    bar    \n",
		"42\n",
	}, ""), buf.String(), t)
}
