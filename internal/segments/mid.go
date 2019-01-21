package segments

import (
	"strings"

	"github.com/mgutz/ansi"
)

type Mid struct {
	Count int
}

func (m Mid) ColoredOutput() string {
	return ansi.ColorFunc("blue+h:black")(m.Output())
}

// Len return length of string without invisible characters counted
func (m Mid) Len() int {
	return m.Count
}

func (m Mid) Output() string {
	if m.Count <= 0 {
		return ""
	}

	return strings.Repeat("─", m.Count)
}
