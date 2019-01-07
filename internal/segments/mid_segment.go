package segments

import (
	"strings"
)

type Mid struct {
	Count int
}

func (m Mid) ColoredOutput() string {
	return m.Output()
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
