package printer

import (
	"fmt"
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

// Flush writes buffered data to output
func (p *Printer) Flush() error {
	return p.w.Flush()
}

// Default returns new Printer of 2-column format
func Default() *Printer {
	return &Printer{
		w: tabwriter.NewWriter(log.Writer(), 0, 0, 2, ' ', 0),
		f: "%s\t%s\t\n",
	}
}
