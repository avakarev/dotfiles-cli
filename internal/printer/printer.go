package printer

import (
	"fmt"
	"io"
	"log"
	"text/tabwriter"
)

// Printer implements column aligned text writer
type Printer struct {
	w *tabwriter.Writer
	f string
}

// Addln adds new line with a and b as values of first and second column
func (p *Printer) Addln(a ...interface{}) (int, error) {
	return fmt.Fprintf(p.w, p.f, a...)
}

// Addf formats according to a format specifier and adds
func (p *Printer) Addf(f string, a ...interface{}) (int, error) {
	return fmt.Fprintf(p.w, f, a...)
}

// Flush writes buffered data to output
func (p *Printer) Flush() error {
	return p.w.Flush()
}

// Default returns new Printer of 2-column format
func Default(w io.Writer) *Printer {
	if w == nil {
		w = log.Writer()
	}
	return &Printer{
		w: tabwriter.NewWriter(w, 0, 0, 2, ' ', 0),
		f: "%s\t%s\t\n",
	}
}
