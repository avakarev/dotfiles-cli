// Package op implement dotfiles operations and reporting
package op

import (
	"bytes"
	"fmt"
	"log"

	"github.com/fatih/color"
)

// Sprint formats op.Result as string
func Sprint(res Result) string {
	buf := &bytes.Buffer{}

	tstate := res.TargetState()
	sstate := res.SourceState()

	buf.WriteString("  ")
	if tstate.IsComplete() && sstate.IsComplete() {
		buf.WriteString(color.New(color.FgGreen).Sprint("✔"))
	} else if tstate.IsError() || sstate.IsError() {
		buf.WriteString(color.New(color.FgRed).Sprint("✘"))
	} else {
		buf.WriteString(" ")
	}
	buf.WriteString(" ")

	buf.WriteString(res.TargetPath() + " ")
	buf.WriteString("[" + tstate.String() + "]")

	buf.WriteString("  →  ")

	buf.WriteString(res.SourcePath() + " ")
	buf.WriteString("[" + sstate.String() + "]")

	return buf.String()
}

// Println formats op.Result as string and writes it to standard output with newline appended
func Println(res Result) (n int, err error) {
	return fmt.Println(Sprint(res))
}

// MustPrintln is like Println but panics in case of error
func MustPrintln(res Result) {
	if _, err := Println(res); err != nil {
		log.Fatalln(err)
	}
}
