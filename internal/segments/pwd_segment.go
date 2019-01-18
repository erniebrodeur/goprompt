package segments

import (
	"fmt"
	"strings"
)

type Pwd struct {
	TerminalWidth int
	Path          string
}

func (p Pwd) ColoredOutput() string {
	return p.Output()
}

// Len return length of string without invisible characters counted
func (p Pwd) Len() int {
	return len(p.Output())
}

func (p Pwd) Output() string {
	return fmt.Sprintf(".../%v", p.t())
}

func (p Pwd) t() string {
	parts := strings.Split(p.Path, "/")
	outputLen := 0
	end := 0

	for i := len(parts) - 1; i >= 0; i-- {
		outputLen += len(parts[i]) + 1 // for the / char

		if outputLen > p.TerminalWidth/4 {
			end = i
			break
		}
	}

	return strings.Join(parts[end:len(parts)], "/")
}
