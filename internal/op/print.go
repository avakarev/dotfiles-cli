package op

import (
	"bytes"
	"fmt"

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
	buf.WriteString("[" + res.TargetState().String() + "]")

	buf.WriteString("  →  ")

	buf.WriteString(res.SourcePath() + " ")
	buf.WriteString("[" + res.SourceState().String() + "]")

	return buf.String()
}

// Println formats op.Result as string and writes it to standard output
// with newline appended
func Println(res Result) (n int, err error) {
	return fmt.Println(Sprint(res))
}
