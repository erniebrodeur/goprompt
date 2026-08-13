package segments

import "strings"

// Mid is for returning a variable size segment
type Mid struct {
	Count int
}

// Len return length of string without invisible characters counted
func (m Mid) Len() int {
	return m.Count
}

// Output returns a mid segment set to the Count length.
func (m Mid) Output() string {
	if m.Count <= 0 {
		return ""
	}

	return strings.Repeat("─", m.Count)
}
