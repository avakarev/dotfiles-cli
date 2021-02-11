package op

import (
	"bytes"
	"fmt"
)

// Sprint formats op.Result as string
func Sprint(res Result) string {
	buf := &bytes.Buffer{}

	buf.WriteString("  " + res.Status() + " ")

	buf.WriteString(res.TargetPath() + " ")
	buf.WriteString("[" + res.TargetState() + "]")

	buf.WriteString("  →  ")

	buf.WriteString(res.SourcePath() + " ")
	buf.WriteString("[" + res.SourceState() + "]")

	return buf.String()
}

// Println formats op.Result as string and writes it to standard output
// with newline appended
func Println(res Result) (n int, err error) {
	return fmt.Println(Sprint(res))
}
